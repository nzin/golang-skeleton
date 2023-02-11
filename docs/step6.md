# Authentication

## Status

Adding an openid connect authentication

(suggestion: use keycloak locally/in docker-compose)

## Content

- You need to install and configure [keycloak](https://www.keycloak.org/) locally
- This code use https://github.com/coreos/go-oidc to talk to Keycloak

## Code explanation

if you checkout step6 branch, you need to understand the code:

### backend

a new middleware has been added (`pkg/config/middleware_openidconnect.go`) that
- will redirect to an IDP (keycloak) if there is no session (or the session expired)
- will get back a JWT token (`pkg/config/middleware_openidconnect.go` line 193)
- will extract some information (the `sub` and store it into the session)

### docker-compose

A docker-compose.yaml has been created. To use it:
```
docker-compose build
docker-compose run
```

and you can then connect to http://localhost:18000 (and identify your self with `admin`/`Pa55w0rd`)

To initiate keycloak with the correct openidconnect client id/secret, a golang application has been created in `initkeycloak/` directory. It is mostly for "demo" purpose, and it initialize Keycloak with a new client app
