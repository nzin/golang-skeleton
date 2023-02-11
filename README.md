![golang](./docs/golang.png)

# The Golang skeleton

This repository is a REST API app golang skeleton.

This skeleton use different tools:

- [go-swagger](https://github.com/go-swagger/go-swagger)
- [gorm.io](https://gorm.io/) for SQL database
- [logrus](https://github.com/sirupsen/logrus) for logging
- [caarlos0/env](github.com/caarlos0/env) for environment variables parsing
- [datadog](https://pkg.go.dev/gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer) for microservice tracing
- [negroni](https://github.com/urfave/negroni) for the http framework
- [VueJS 2.x](https://vuejs.org/) when there is a UI

# The skeleton

This repository is divided into 6 steps representing different technologies put in place.

For each step, there is a corresponding git branch (step1, step2, ...) with a/the solution

Steps:
- [1- go-swagger backend service](./docs/step1.md)
- [2- Database storage](./docs/step2.md)
- [3- Unit Tests](./docs/step3.md)
- [4- Microservice tracing](./docs/step4.md)
- [5- VueJS UI](./docs/step5.md)
- [6- Authentication](./docs/step6.md)

