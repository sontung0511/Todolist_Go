package model

import (
	"Todolist_Go/common"
	"errors"
	"strings"
)

var (
	ErrTittleCannotEmpty = errors.New("tittle cannot be empty")
	ErrItemIsDeleted     = errors.New("da deleted")
)

type TodoItem struct {
	common.SQLModel
	Tittle      string `json:"tittle" gorm:"column:tittle;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Tittle      string `json:"tittle" gorm:"column:tittle;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (i *TodoItemCreation) Validate() error {
	i.Tittle = strings.TrimSpace(i.Tittle)
	if i.Tittle == "" {
		return ErrTittleCannotEmpty
	}
	return nil
}
func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Tittle      string `json:"tittle" gorm:"column:tittle;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
