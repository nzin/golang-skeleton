package main

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
)

func main() {
	env.Parse(&Config)

	keycloak := NewKeycloak(Config.KeycloakUsername, Config.KeycloakPassword, Config.KeycloakAddress)
	for i := 0; i < 10; i++ {
		err := keycloak.CreateClient(Config.OpenIDClientID, Config.OpenIDClientSecret, Config.OpenIDClientUrl)
		if err == nil {
			fmt.Println("Keycloak initialized successfully")
			os.Exit(0)
		}
		time.Sleep(10 * time.Second)
	}
	fmt.Println("not able to initialize Keycloak")
}
