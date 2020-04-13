package models

import (
	"fmt"

	"github.com/WTBacon/todo-app-gin/config"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func FetchAllTodo(todo *[]Todo) error {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func FetchSingleTodo(todo *Todo, id string) error {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) error {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo) error {
	fmt.Println(*todo)
	config.DB.Model(todo).Update(todo)
	return nil
}

func DeleteTodo(todo *Todo, id string) error {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
