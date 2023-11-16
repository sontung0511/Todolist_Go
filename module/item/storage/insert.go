package storage

import (
	"Todolist_Go/module/item/model"
	"context"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
