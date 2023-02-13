package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateTodo(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		id, err := CreateTodo(db, "foo", "bar")
		assert.Nil(t, err)

		var todo Todo
		db.First(&todo)
		assert.Equal(t, id, todo.ID)
		assert.Equal(t, "foo", todo.Title)
		assert.Equal(t, "bar", todo.Body)
	})
}

func TestDeleteTodo(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		id, err := CreateTodo(db, "foo", "bar")
		assert.Nil(t, err)

		err = DeleteTodo(db, id)
		assert.Nil(t, err)

		var todo Todo
		err = db.First(&todo).Error
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})

	t.Run("not happy path: no record", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		err = DeleteTodo(db, 1)
		assert.Equal(t, nil, err)
	})
}

func TestUpdateReadTodo(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		id, err := CreateTodo(db, "foo", "bar")
		assert.Nil(t, err)

		err = UpdateTodo(db, id, "foo2", "bar2")
		assert.Nil(t, err)

		todo, err := ReadTodo(db, id)
		assert.Nil(t, err)

		assert.Equal(t, "foo2", todo.Title)
		assert.Equal(t, "bar2", todo.Body)
	})

	t.Run("not happy path: no record", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		err = UpdateTodo(db, 1, "foo", "bar")
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestListTodos(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		db, err := NewTestDB()
		assert.Nil(t, err)

		id, err := CreateTodo(db, "foo", "bar")
		assert.Nil(t, err)

		id2, err := CreateTodo(db, "foo2", "bar2")
		assert.Nil(t, err)

		todos, err := ListTodos(db)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(todos))
		assert.Equal(t, id, todos[0].ID)
		assert.Equal(t, id2, todos[1].ID)
		assert.Equal(t, "foo2", todos[1].Title)
		assert.Equal(t, "bar2", todos[1].Body)
	})
}
