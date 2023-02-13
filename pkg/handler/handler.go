package handler

import (
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/app"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/health"
)

// Setup initialize all the handler functions
func Setup(api *operations.GolangSkeletonAPI) {
	c := NewCRUD()

	// healthcheck
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(c.GetHealthcheck)

	// todo functions
	api.AppListTodosHandler = app.ListTodosHandlerFunc(c.ListTodos)
	api.AppGetTodoHandler = app.GetTodoHandlerFunc(c.GetTodo)
	api.AppCreateTodoHandler = app.CreateTodoHandlerFunc(c.CreateTodo)
	api.AppPutTodoHandler = app.PutTodoHandlerFunc(c.UpdateTodo)
	api.AppDeleteTodoHandler = app.DeleteTodoHandlerFunc(c.DeleteTodo)
}
