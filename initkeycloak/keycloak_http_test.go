package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeycloakAccessToken(t *testing.T) {
	t.Run("happy code path: return correct access token", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/auth/realms/master/protocol/openid-connect/token" {
				assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
				r.ParseForm()
				assert.Equal(t, "admin", r.Form["username"][0])
				assert.Equal(t, "password", r.Form["password"][0])
				assert.Equal(t, "admin-cli", r.Form["client_id"][0])
				assert.Equal(t, "password", r.Form["grant_type"][0])
				w.WriteHeader(200)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{
					"access_token":"tokenvalue",
					"scope":"openid"
				}`)
				return
			}
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"missing values"}`)
		}))
		defer svr.Close()

		k := &KeycloakHttpCallerImpl{
			AdminUsername: "admin",
			AdminPassword: "password",
			Address:       svr.URL,
		}
		token, err := k.getAccessToken()
		assert.Nil(t, err)
		assert.Equal(t, "tokenvalue", token)
	})

	t.Run("not happy code path: return 500 instead of access token", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"missing values"}`)
		}))
		defer svr.Close()

		k := &KeycloakHttpCallerImpl{
			AdminUsername: "admin",
			AdminPassword: "password",
			Address:       svr.URL,
		}
		_, err := k.getAccessToken()
		assert.NotNil(t, err)
	})
}

func TestKeycloakGenericCall(t *testing.T) {
	t.Run("happy code path: test a GET call", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/foobar" && r.Method == "GET" {
				w.WriteHeader(200)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{"status":"ok"}`)
				return
			}
			if r.URL.Path == "/auth/realms/master/protocol/openid-connect/token" {
				assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
				r.ParseForm()
				assert.Equal(t, "admin", r.Form["username"][0])
				assert.Equal(t, "password", r.Form["password"][0])
				assert.Equal(t, "admin-cli", r.Form["client_id"][0])
				assert.Equal(t, "password", r.Form["grant_type"][0])
				w.WriteHeader(200)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{
					"access_token":"tokenvalue",
					"scope":"openid"
				}`)
				return
			}
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"missing values"}`)
		}))
		defer svr.Close()

		k := &KeycloakHttpCallerImpl{
			AdminUsername: "admin",
			AdminPassword: "password",
			Address:       svr.URL,
		}
		body, status, err := k.DoHttp("GET", "/foobar", nil)
		assert.Nil(t, err)
		assert.Equal(t, `{"status":"ok"}`, string(body))
		assert.Equal(t, 200, status)
	})

	t.Run("happy code path: test a POST call", func(t *testing.T) {
		type Payload struct {
			A string
			B int
		}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/foobar" && r.Method == "POST" {
				body, err := ioutil.ReadAll(r.Body)
				assert.Nil(t, err)
				var p Payload
				json.Unmarshal(body, &p)
				assert.Equal(t, "foo", p.A)
				assert.Equal(t, 10, p.B)

				w.WriteHeader(200)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{"status":"ok"}`)
				return
			}
			if r.URL.Path == "/auth/realms/master/protocol/openid-connect/token" {
				assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
				r.ParseForm()
				assert.Equal(t, "admin", r.Form["username"][0])
				assert.Equal(t, "password", r.Form["password"][0])
				assert.Equal(t, "admin-cli", r.Form["client_id"][0])
				assert.Equal(t, "password", r.Form["grant_type"][0])
				w.WriteHeader(200)
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprintf(w, `{
					"access_token":"tokenvalue",
					"scope":"openid"
				}`)
				return
			}
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"missing values"}`)
		}))
		defer svr.Close()

		k := &KeycloakHttpCallerImpl{
			AdminUsername: "admin",
			AdminPassword: "password",
			Address:       svr.URL,
		}

		p := Payload{
			A: "foo",
			B: 10,
		}
		body, status, err := k.DoHttp("POST", "/foobar", p)
		assert.Nil(t, err)
		assert.Equal(t, `{"status":"ok"}`, string(body))
		assert.Equal(t, 200, status)
	})

	t.Run("not happy code path: return 500 instead of access token", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"missing values"}`)
		}))
		defer svr.Close()

		k := &KeycloakHttpCallerImpl{
			AdminUsername: "admin",
			AdminPassword: "password",
			Address:       svr.URL,
		}
		_, err := k.getAccessToken()
		assert.NotNil(t, err)
	})
}
