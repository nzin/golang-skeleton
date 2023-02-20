package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Keycloak struct {
	httpCaller KeycloakHttpCaller
}

func NewKeycloak(keycloakUsername, keycloakPassword, keycloakAddress string) *Keycloak {
	return &Keycloak{
		httpCaller: &KeycloakHttpCallerImpl{
			AdminUsername: keycloakUsername,
			AdminPassword: keycloakPassword,
			Address:       keycloakAddress,
		},
	}
}

// https://www.keycloak.org/docs-api/15.0/rest-api/#_clientrepresentation
type RealmRepresentation struct {
	Id      string `json:"id"`
	Realm   string `json:"realm"`
	Enabled bool   `json:"enabled"`
}

type ClientRepresentation struct {
	ID                      string   `json:"id"`
	ClientID                string   `json:"clientId"`
	ClientAuthenticatorType string   `json:"clientAuthenticatorType"`
	Enabled                 bool     `json:"enabled"`
	Name                    string   `json:"name"`
	Protocol                string   `json:"protocol"`
	PublicClient            bool     `json:"publicClient"`
	Secret                  string   `json:"secret"`
	RootUrl                 string   `json:"rootUrl"`
	RedirectUris            []string `json:"redirectUris"`
}

// CreateTenant returns tenant_id, portal_client_id, portal_client_secret
func (k *Keycloak) CreateClient(clientId, clientSecret, clientRootUrl string) error {
	// create client

	id := uuid.New()
	cr := ClientRepresentation{
		ID:                      id.String(),
		ClientID:                clientId,
		ClientAuthenticatorType: "client-secret",
		Enabled:                 true,
		Name:                    "todo-app",
		Protocol:                "openid-connect",
		PublicClient:            false,
		Secret:                  clientSecret,
		RootUrl:                 clientRootUrl,
		RedirectUris:            []string{clientRootUrl + "/*"},
	}

	body, statuscode, err := k.httpCaller.DoHttp("POST", "/auth/admin/realms/master/clients", cr)
	if err != nil {
		return fmt.Errorf("not able to call /auth/admin/realms/master/clients (%d): %s", statuscode, string(body))
	}

	if statuscode != 201 {
		return fmt.Errorf("not able to create client for realm master (%d): %s", statuscode, string(body))
	}

	return nil
}
