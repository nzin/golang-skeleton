package config

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type OpenidConfiguration struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	EndSessionEndpoint    string `json:"end_session_endpoint"`
	JwksUri               string `json:"jwks_uri"`
}

type keycloakOpenidMiddleware struct {
	config        *OpenidConfiguration
	jwtPublicKeys map[string]*rsa.PublicKey
}

func setupKeycloakOpenidMiddleware() (*keycloakOpenidMiddleware, error) {

	var conf OpenidConfiguration

	resp, err := http.Get(Config.OpenIDIdp + "/auth/realms/master/.well-known/openid-configuration")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &conf)
	if err != nil {
		return nil, err
	}

	return &keycloakOpenidMiddleware{
		config:        &conf,
		jwtPublicKeys: make(map[string]*rsa.PublicKey),
	}, nil
}

const (
	keyTokenExpiry = "oauth2_expiry"
	keyTokenSub    = "jwt_sub"
	keyNextPage    = "next"
	keyState       = "state"
)

func (m *keycloakOpenidMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	for _, path := range Config.OpenIDAuthPrefixWhitelistPaths {
		if strings.HasPrefix(r.URL.Path, path) {
			next(w, r)
			return
		}
	}
	s := sessions.GetSession(r)

	//	logrus.Warnf("path: %s", r.URL.Path)
	if r.URL.Path == "/login" {
		m.login(s, w, r)
		return
	}

	if r.URL.Path == "/logout" {
		m.logout(s, w, r)
		return
	}

	if r.URL.Path == "/oauth2callback" {
		m.handleOAuth2Callback(s, w, r)
		return
	}

	sub := s.Get(keyTokenSub)
	timestamp := s.Get(keyTokenExpiry)
	var expiry time.Time
	if timestamp != nil {
		expiry = time.Unix(timestamp.(int64), 0)
	}
	if sub == nil || expiry.Before(time.Now()) {
		// Set token to null to avoid redirection loop
		s.Set(keyTokenSub, nil)
		s.Set(keyTokenExpiry, nil)
		next := url.QueryEscape(r.URL.RequestURI())
		http.Redirect(w, r, "/login?"+keyNextPage+"="+next, http.StatusFound)
	} else {
		next(w, r)
	}
}

func (m *keycloakOpenidMiddleware) login(s sessions.Session, w http.ResponseWriter, r *http.Request) {
	next := extractPath(r.URL.Query().Get(keyNextPage))

	if s.Get(keyTokenSub) == nil {
		// User is not logged in.
		if next == "" {
			next = "/"
		}

		state := newState()
		// store the next url and state token in the session
		s.Set(keyState, state)
		s.Set(keyNextPage, next)

		config := oauth2.Config{
			ClientID:     Config.OpenIDClientID,
			ClientSecret: Config.OpenIDClientSecret,
			Scopes:       []string{"openid"},
			RedirectURL:  fmt.Sprintf("%s/oauth2callback", Config.OpenIDBaseURL),
			Endpoint: oauth2.Endpoint{
				AuthURL:  m.config.AuthorizationEndpoint, // Config.ExternalFusionAuthUrl + "/oauth2/authorize",
				TokenURL: m.config.TokenEndpoint,         //Config.FusionAuthUrl + "/oauth2/token",
			},
		}

		if Config.OpenIDExternalIdp != "" {
			http.Redirect(w, r, strings.Replace(config.AuthCodeURL(state, oauth2.AccessTypeOffline), Config.OpenIDIdp, Config.OpenIDExternalIdp, 1), http.StatusFound)
		} else {
			http.Redirect(w, r, config.AuthCodeURL(state, oauth2.AccessTypeOffline), http.StatusFound)
		}
		return
	}
	// No need to login, redirect to the next page.
	http.Redirect(w, r, next, http.StatusFound)
}

func (m *keycloakOpenidMiddleware) logout(s sessions.Session, w http.ResponseWriter, r *http.Request) {
	s.Delete(keyTokenSub)

	if Config.OpenIDExternalIdp != "" {
		http.Redirect(w, r, strings.Replace(m.config.EndSessionEndpoint, Config.OpenIDIdp, Config.OpenIDExternalIdp, 1), http.StatusFound) // +"?redirect_uri"+url.QueryEscape("http://"+host+":8000"), http.StatusFound)
	} else {
		http.Redirect(w, r, m.config.EndSessionEndpoint, http.StatusFound) // +"?redirect_uri"+url.QueryEscape("http://"+host+":8000"), http.StatusFound)

	}
}

func (m *keycloakOpenidMiddleware) handleOAuth2Callback(s sessions.Session, w http.ResponseWriter, r *http.Request) {
	providedState := extractPath(r.URL.Query().Get("state"))

	//verify that the provided state is the state we generated
	//if it is not, then redirect to the error page
	originalState := s.Get(keyState)
	if providedState != originalState {
		logrus.Infof("providedState != originalState, original state: %v", originalState)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	next := s.Get(keyNextPage).(string)
	code := r.URL.Query().Get("code")
	config := oauth2.Config{
		ClientID:     Config.OpenIDClientID,
		ClientSecret: Config.OpenIDClientSecret,
		Scopes:       []string{"openid"},
		RedirectURL:  fmt.Sprintf("%s/oauth2callback", Config.OpenIDBaseURL),
		Endpoint: oauth2.Endpoint{
			AuthURL:  m.config.AuthorizationEndpoint, // Config.ExternalFusionAuthUrl + "/oauth2/authorize",
			TokenURL: m.config.TokenEndpoint,         //Config.FusionAuthUrl + "/oauth2/token",
		},
	}
	t, err := config.Exchange(context.Background(), code)
	if err != nil {
		// Pass the error message, or allow dev to provide its own
		// error handler.
		logrus.Infof("error exchanging (%v): %v", config, err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	// we dont want to store the whole JWT token
	s.Set(keyTokenExpiry, t.Expiry.Unix())

	// now we check the validity of the JWT token (issued by the correct provider)
	token, err := jwt.Parse(t.AccessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		keyid := token.Header["kid"].(string)

		if pubkey, ok := m.jwtPublicKeys[keyid]; ok {
			return pubkey, nil
		}

		set, err := jwk.Fetch(r.Context(), m.config.JwksUri)
		if err != nil {
			return nil, err
		}

		if key, ok := set.LookupKeyID(keyid); ok {
			var pubkey rsa.PublicKey
			key.Raw(&pubkey)
			m.jwtPublicKeys[keyid] = &pubkey
			return &pubkey, nil
		}

		return nil, fmt.Errorf("keyid not found")
	})
	if err != nil {
		logrus.Infof("Not able to parse correctly the JWT token: %v", err)
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	// getting the "claims", i.e. the payload of the JWT token

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		s.Set(keyTokenSub, claims["sub"])
		//			logrus.Infof("email: %s", claims["email"])
		//			logrus.Infof("userid: %s", claims["sub"])
		//			logrus.Infof("tenandid: %s", claims["tid"])
	} else {
		logrus.Infof("not able to get JWT claim correctly")
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	http.Redirect(w, r, next, http.StatusFound)
}

func extractPath(next string) string {
	n, err := url.Parse(next)
	if err != nil {
		return "/"
	}
	return n.Path
}

func newState() string {
	var p [16]byte
	_, err := rand.Read(p[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(p[:])
}
