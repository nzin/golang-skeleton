package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type KeycloakHttpCaller interface {
	DoHttp(method string, path string, bodydata interface{}) ([]byte, int, error)
}

type KeycloakHttpCallerImpl struct {
	AdminUsername string
	AdminPassword string
	Address       string
}

type AccessToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func (k *KeycloakHttpCallerImpl) getAccessToken() (string, error) {
	requestURL := fmt.Sprintf("%s/auth/realms/master/protocol/openid-connect/token", k.Address)

	data := url.Values{}
	data.Set("client_id", "admin-cli")
	data.Set("grant_type", "password")
	data.Set("username", k.AdminUsername)
	data.Set("password", k.AdminPassword)

	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("unable to get the access token (%d): %s", resp.StatusCode, string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var token AccessToken
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

// httpCall is here to do a generic http call. It will return body content, statuscode, error
func (k *KeycloakHttpCallerImpl) DoHttp(method string, path string, bodydata interface{}) ([]byte, int, error) {
	var reader io.Reader
	requestURL := fmt.Sprintf("%s%s", k.Address, path)

	if bodydata != nil {
		body, err := json.Marshal(bodydata)
		if err != nil {
			return nil, 0, err
		}
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, requestURL, reader)
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, 0, err
	}

	accesstoken, err := k.getAccessToken()
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", "Bearer "+accesstoken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 20,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
