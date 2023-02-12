package handler

import (
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/health"
)

// Setup initialize all the handler functions
func Setup(api *operations.GolangSkeletonAPI) {
	c := NewCRUD()

	// healthcheck
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(c.GetHealthcheck)
}
