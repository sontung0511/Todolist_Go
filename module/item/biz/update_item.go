package biz

import (
	"Todolist_Go/module/item/model"
	"context"
)

//find id

type UpdateItemStore interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}
type UpdateItemBiz struct {
	store UpdateItemStore
}

func NewUpdateItemBiz(store UpdateItemStore) *UpdateItemBiz {
	return &UpdateItemBiz{store: store}
}
func (biz *UpdateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status == "Deleted" {
		return model.ErrItemIsDeleted
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return nil
	}
	return nil
}
