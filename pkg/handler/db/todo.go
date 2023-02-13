package db

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Title string
	Body  string `sql:"type:text"`
}

func ListTodos(db *gorm.DB) ([]*Todo, error) {
	var list []*Todo
	err := db.Find(&list).Error
	return list, err
}

func CreateTodo(db *gorm.DB, title, body string) (uint, error) {
	t := Todo{
		Title: title,
		Body:  body,
	}
	err := db.Create(&t).Error
	return t.ID, err
}

func ReadTodo(db *gorm.DB, todoid uint) (*Todo, error) {
	where := Todo{}
	where.ID = todoid
	var t Todo
	err := db.Where(&where).First(&t).Error
	return &t, err
}

func UpdateTodo(db *gorm.DB, todoid uint, title, body string) error {
	where := Todo{}
	where.ID = todoid
	var t Todo
	err := db.Where(&where).First(&t).Error
	if err != nil {
		return err
	}
	t.Title = title
	t.Body = body
	return db.Save(&t).Error
}

func DeleteTodo(db *gorm.DB, todoid uint) error {
	where := Todo{}
	where.ID = todoid
	return db.Unscoped().Where(&where).Delete(&where).Error
}
