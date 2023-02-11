# REST API service skeleton

## Status

This step:

- has 1 endpoint: an health-check endpoint (for kubernetes)
- use github.com/caarlos0/env to deal with env variables
- logs in json to the output (using logrus)


## Content

- It use go-swagger (https://github.com/go-swagger/go-swagger)
- The tricky part is to setup a http framework (gin / negroni / …). See
  - https://goswagger.io/use/middleware.html
  - Or check how golang-skeleton works:
    - https://github.com/nzin/golang-skeleton/blob/step1/swagger_gen/restapi/configure_golang-skeleton.go 
    - https://github.com/nzin/golang-skeleton/blob/step1/pkg/config/middleware.go 
- it add one /api/v1/healthcheck endpoint that returns 200
- it is dockerized (see Dockerfile)
- it use logrus middleware with a JSON format by default
- all environment variables are in 1 golang struct, managed by github.com/caarlos0/env

## Tips

if you install swagger via

```
go install github.com/go-swagger/go-swagger/cmd/swagger@v0.26.0
```

dont forget to add ~/go/bin to your path:
```
export PATH=$PATH:~/go/bin
```

## Code explanation

if you checkout step1 branch, you need to understand the code:

### Swagger endpoint implementation

First, you can install swagger tools:

```
make deps
```

Thew the `swagger/` directory contains the swagger definition. and the main file is the `swagger/index.yaml` file:

```
...
paths:
  /health:
    $ref: ./health.yaml
...
```

You have here the definition of the first endpoint (`/api/v1/health`), and in `swagger/health.yaml`:

```
get:
  tags:
    - health
  operationId: getHealth
```

you find the `operationId`, the function that will deliver the content/payload of the endpoint

To generate the stub/skeleton you need to run

```
make gen
```

It will generate/update the `swagger_gen` directory, and in particular the `swagger_gen/restapi/configure_golang_skeleton.go` is interesting, especially line 36:

```
import (
  ...
	"github.com/nzin/golang-skeleton/pkg/handler"
  ...
  )
...
	handler.Setup(api)
...
```

And in `pkg/handler/handler.go`, this is where we tell swagger, where to find the implementation of the REST API endpoints

```

// Setup initialize all the handler functions
func Setup(api *operations.GolangSkeletonAPI) {
	c := NewCRUD()

	// healthcheck
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(c.GetHealthcheck)
}
```

As you can see, there is a `CRUD` object, which contains the real business functions. This `CRUD` is defined in `pkg/handler/crud.go` 

```

// CRUD is the CRUD interface
type CRUD interface {
	// healthcheck
	GetHealthcheck(health.GetHealthParams) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	return &crud{}
}

type crud struct {}

func (c *crud) GetHealthcheck(params health.GetHealthParams) middleware.Responder {
	return health.NewGetHealthOK().WithPayload(&models.Health{Status: "OK"})
}
```

So if you want to add new REST endpoints, you will need to
- update the `swagger/index.yaml` (plus a second yaml file with the get/post/put/delete/... definition)
- update the `pkg/handler/handler.go` to tell swagger the real implementation
- update the `pkg/handler/crud.go` to implement the endpoint function

### Swagger http middleware

The second interesting part is (again) in the  `swagger_gen/restapi/configure_golang_skeleton.go`, especially line 60-62 :

```
import (
  ...
	"github.com/nzin/golang-skeleton/pkg/config"
  ...
  )

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return config.SetupGlobalMiddleware(handler)
}
```

and in `pkg/config/middleware.go` : 

```
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

	n.Use(setupRecoveryMiddleware())

	n.UseHandler(handler)

	return n
}```

It is telling swagger to use the negroni http framework, and it is adding some middlewares to
- allow gzip compression
- adding automatic logging (using logrus)
- adding a failback/recovery in case of panic() function
