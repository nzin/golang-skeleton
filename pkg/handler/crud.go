package handler

import (
	"github.com/nzin/golang-skeleton/swagger_gen/models"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/app"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/health"
	"gorm.io/gorm"

	"github.com/nzin/golang-skeleton/pkg/handler/db"
	"github.com/go-openapi/runtime/middleware"
)

// CRUD is the CRUD interface
type CRUD interface {
	// healthcheck
	GetHealthcheck(health.GetHealthParams) middleware.Responder
	ListTodos(app.ListTodosParams) middleware.Responder
	CreateTodo(app.CreateTodoParams) middleware.Responder
	GetTodo(app.GetTodoParams) middleware.Responder
	UpdateTodo(app.PutTodoParams) middleware.Responder
	DeleteTodo(app.DeleteTodoParams) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	return &crud{
		db: db,
	}
}

type crud struct {
	db *gorm.DB
}

func (c *crud) GetHealthcheck(params health.GetHealthParams) middleware.Responder {
	return health.NewGetHealthOK().WithPayload(&models.Health{Status: "OK"})
}

func (c *crud) ListTodos(app.ListTodosParams) middleware.Responder {
	list, err := db.ListTodos(c.db)
	if err != nil {
		return app.NewListTodosDefault(500).WithPayload(
			ErrorMessage("cannot find todos: %v", err))
	}
	payload := app.ListTodosOKBody{
		Todos: make([]*models.Todo, len(list)),
	}
	for i, t := range list {
		payload.Todos[i] = &models.Todo{
			ID:    int64(t.ID),
			Title: &t.Title,
			Body:  &t.Body,
		}
	}
	return app.NewListTodosOK().WithPayload(&payload)
}

func (c *crud) CreateTodo(params app.CreateTodoParams) middleware.Responder {
	id, err := db.CreateTodo(c.db, params.Body.Title, params.Body.Body)
	if err != nil {
		return app.NewCreateTodoDefault(500).WithPayload(
			ErrorMessage("cannot create todo: %v", err))
	}
	payload := app.CreateTodoOKBody{
		TodoID: int64(id),
	}
	return app.NewCreateTodoOK().WithPayload(&payload)
}

func (c *crud) GetTodo(params app.GetTodoParams) middleware.Responder {
	t, err := db.ReadTodo(c.db, uint(params.TodoID))
	if err != nil {
		return app.NewGetTodoDefault(500).WithPayload(
			ErrorMessage("cannot get todo %d: %v", params.TodoID, err))
	}
	payload := models.Todo{
		ID:    int64(t.ID),
		Title: &t.Title,
		Body:  &t.Body,
	}
	return app.NewGetTodoOK().WithPayload(&payload)

}

func (c *crud) UpdateTodo(params app.PutTodoParams) middleware.Responder {
	err := db.UpdateTodo(c.db, uint(params.TodoID), params.Body.Title, params.Body.Body)
	if err != nil {
		return app.NewPutTodoDefault(500).WithPayload(
			ErrorMessage("cannot update todo %d: %v", params.TodoID, err))
	}
	return app.NewPutTodoOK()
}

func (c *crud) DeleteTodo(params app.DeleteTodoParams) middleware.Responder {
	err := db.DeleteTodo(c.db, uint(params.TodoID))
	if err != nil {
		return app.NewDeleteTodoDefault(500).WithPayload(
			ErrorMessage("cannot delete todo %d: %v", params.TodoID, err))
	}
	return app.NewDeleteTodoOK()
}
