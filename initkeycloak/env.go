package main

// Config is the whole configuration of the app
var Config = struct {
	KeycloakUsername   string `env:"GOLANG_SKELETON_KEYCLOAK_USERNAME"`
	KeycloakPassword   string `env:"GOLANG_SKELETON_KEYCLOAK_PASSWORD"`
	KeycloakAddress    string `env:"GOLANG_SKELETON_KEYCLOAK_ADDRESS" envDefault:"http://localhost:8080"`
	OpenIDClientID     string `env:"GOLANG_SKELETON_OPENID_CLIENT_ID"`
	OpenIDClientSecret string `env:"GOLANG_SKELETON_OPENID_CLIENT_SECRET"`
	OpenIDClientUrl    string `env:"GOLANG_SKELETON_OPENID_CLIENT_URL" envDefault:"http://localhost:18000"`
}{}
