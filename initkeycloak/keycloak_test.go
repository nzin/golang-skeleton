package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type KeycloakHttpCallerMock struct {
	callback func(method string, path string, bodydata interface{}) ([]byte, int, error)
}

func (k *KeycloakHttpCallerMock) DoHttp(method string, path string, bodydata interface{}) ([]byte, int, error) {
	return k.callback(method, path, bodydata)
}

func TestKeycloakCreateTenant(t *testing.T) {
	t.Run("happy code path: creating tenant", func(t *testing.T) {
		mock := KeycloakHttpCallerMock{
			callback: func(method string, path string, bodydata interface{}) ([]byte, int, error) {
				if path == "/auth/admin/realms/master/clients" && method == "POST" {
					cr := bodydata.(ClientRepresentation)
					assert.Equal(t, "clientid", cr.ClientID)
					assert.Equal(t, "clientsecret", cr.Secret)
					assert.Equal(t, "http://localhost:18000", cr.RootUrl)

					return nil, 201, nil
				}
				return nil, 500, nil
			},
		}

		k := &Keycloak{
			httpCaller: &mock,
		}
		err := k.CreateClient("clientid", "clientsecret", "http://localhost:18000")
		assert.Nil(t, err)
	})

	t.Run("not happy code path: return 500 instead of creating tenant", func(t *testing.T) {
		mock := KeycloakHttpCallerMock{
			callback: func(method string, path string, bodydata interface{}) ([]byte, int, error) {
				return nil, 500, nil
			},
		}

		k := &Keycloak{
			httpCaller: &mock,
		}
		err := k.CreateClient("clientid", "clientsecret", "http://localhost:18000")
		assert.NotNil(t, err)
	})
}
