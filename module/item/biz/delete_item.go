package biz

import (
	"Todolist_Go/module/item/model"
	"context"
)

type DeleteItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}
type DeleteItemBiz struct {
	store DeleteItemStore
}

func NewUDeleteItemBiz(store DeleteItemStore) *DeleteItemBiz {
	return &DeleteItemBiz{store: store}
}
func (biz *DeleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status == "Deleted" {
		return model.ErrItemIsDeleted
	}
	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return nil
	}
	return nil
}
