package handler

import (
	"testing"

	"github.com/nzin/golang-skeleton/pkg/handler/db"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/app"
	"github.com/nzin/golang-skeleton/swagger_gen/restapi/operations/health"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestHealth(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		db, err := db.NewTestDB()
		assert.Nil(t, err)
		c := crud{
			db: db,
		}
		res := c.GetHealthcheck(health.GetHealthParams{})
		_, ok := res.(*health.GetHealthOK)
		assert.Equal(t, true, ok)
	})
}

func TestCRUD(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d, err := db.NewTestDB()
		assert.Nil(t, err)
		c := crud{
			db: d,
		}

		// Create

		res1 := c.CreateTodo(app.CreateTodoParams{
			Body: app.CreateTodoBody{
				Title: "foo",
				Body:  "bar",
			},
		})
		resp1, ok := res1.(*app.CreateTodoOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, int64(1), resp1.Payload.TodoID)

		var todo db.Todo
		err = d.First(&todo).Error
		assert.Nil(t, err)

		// Update

		res2 := c.UpdateTodo(app.PutTodoParams{
			Body: app.PutTodoBody{
				Title: "foo2",
				Body:  "bar2",
			},
			TodoID: 1,
		})
		_, ok = res2.(*app.PutTodoOK)
		assert.Equal(t, true, ok)

		err = d.First(&todo).Error
		assert.Nil(t, err)
		assert.Equal(t, "foo2", todo.Title)
		assert.Equal(t, "bar2", todo.Body)

		// Get

		res3 := c.GetTodo(app.GetTodoParams{
			TodoID: 1,
		})
		resp3, ok := res3.(*app.GetTodoOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, "foo2", *resp3.Payload.Title)
		assert.Equal(t, "bar2", *resp3.Payload.Body)

		// Delete

		res4 := c.DeleteTodo(app.DeleteTodoParams{
			TodoID: 1,
		})
		_, ok = res4.(*app.DeleteTodoOK)
		assert.Equal(t, true, ok)

		err = d.First(&todo).Error
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})

	t.Run("not happy path: read unknown record", func(t *testing.T) {
		d, err := db.NewTestDB()
		assert.Nil(t, err)
		c := crud{
			db: d,
		}

		// Update

		res1 := c.GetTodo(app.GetTodoParams{
			TodoID: 1,
		})
		_, ok := res1.(*app.GetTodoDefault)
		assert.Equal(t, true, ok)
	})

	t.Run("not happy path: update unknown record", func(t *testing.T) {
		d, err := db.NewTestDB()
		assert.Nil(t, err)
		c := crud{
			db: d,
		}

		// Update

		res1 := c.UpdateTodo(app.PutTodoParams{
			Body: app.PutTodoBody{
				Title: "foo2",
				Body:  "bar2",
			},
			TodoID: 1,
		})
		_, ok := res1.(*app.PutTodoDefault)
		assert.Equal(t, true, ok)
	})
}

func TestListTodos(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d, err := db.NewTestDB()
		assert.Nil(t, err)
		c := crud{
			db: d,
		}

		// Create 3 records

		res1 := c.CreateTodo(app.CreateTodoParams{
			Body: app.CreateTodoBody{
				Title: "foo",
				Body:  "bar",
			},
		})
		resp1, ok := res1.(*app.CreateTodoOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, int64(1), resp1.Payload.TodoID)

		res2 := c.CreateTodo(app.CreateTodoParams{
			Body: app.CreateTodoBody{
				Title: "foo2",
				Body:  "bar2",
			},
		})
		resp2, ok := res2.(*app.CreateTodoOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, int64(2), resp2.Payload.TodoID)

		res3 := c.CreateTodo(app.CreateTodoParams{
			Body: app.CreateTodoBody{
				Title: "foo3",
				Body:  "bar3",
			},
		})
		resp3, ok := res3.(*app.CreateTodoOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, int64(3), resp3.Payload.TodoID)

		// Delete

		res4 := c.DeleteTodo(app.DeleteTodoParams{
			TodoID: 2,
		})
		_, ok = res4.(*app.DeleteTodoOK)
		assert.Equal(t, true, ok)

		// List

		res5 := c.ListTodos(app.ListTodosParams{})
		resp5, ok := res5.(*app.ListTodosOK)
		assert.Equal(t, true, ok)
		assert.Equal(t, 2, len(resp5.Payload.Todos))
	})
}
