package config

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	negronitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/urfave/negroni"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// ServerShutdown is a callback function that will be called when
// we tear down the golang-skeleton server
func ServerShutdown() {
	if Config.DatadogTraceEnabled {
		tracer.Stop()
	}
}

// SetupGlobalMiddleware setup the global middleware
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	n := negroni.New()

	if Config.MiddlewareGzipEnabled {
		n.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if Config.MiddlewareVerboseLoggerEnabled {
		middleware := negronilogrus.NewMiddlewareFromLogger(logrus.StandardLogger(), "google-skeleton")

		for _, u := range Config.MiddlewareVerboseLoggerExcludeURLs {
			middleware.ExcludeURL(u)
		}

		n.Use(middleware)
	}

	if Config.DatadogTraceEnabled {
		tracer.Start()
		n.Use(negronitrace.Middleware(negronitrace.WithServiceName("appconfigr")))
	}

	if Config.CORSEnabled {
		n.Use(cors.New(cors.Options{
			AllowedOrigins:   Config.CORSAllowedOrigins,
			AllowedHeaders:   Config.CORSAllowedHeaders,
			ExposedHeaders:   Config.CORSExposedHeaders,
			AllowedMethods:   Config.CORSAllowedMethods,
			AllowCredentials: Config.CORSAllowCredentials,
		}))
	}

	n.Use(sessions.Sessions("golang", cookiestore.New([]byte(Config.SessionNonce))))

	if Config.OpenIDAuthEnabled {
		m, err := setupKeycloakOpenidMiddleware()
		if err != nil {
			panic(err)
		}
		n.Use(m)
	}

	n.Use(&negroni.Static{
		Dir:       http.Dir("./browser/golang-challenge-ui/dist/"),
		Prefix:    Config.WebPrefix,
		IndexFile: "index.html",
	})

	n.Use(setupRecoveryMiddleware())

	n.UseHandler(handler)

	return n
}

type recoveryLogger struct{}

func (r *recoveryLogger) Printf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func (r *recoveryLogger) Println(v ...interface{}) {
	logrus.Errorln(v...)
}

func setupRecoveryMiddleware() *negroni.Recovery {
	r := negroni.NewRecovery()
	r.Logger = &recoveryLogger{}
	return r
}
