package storage

import (
	"Todolist_Go/module/item/model"
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "Deleted"

	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).
		Error; err != nil {
		return err
	}
	return nil
}
