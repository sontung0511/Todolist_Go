package biz

import (
	"Todolist_Go/module/item/model"
	"golang.org/x/net/context"
)

// Handler[di ve json tra ve cho client] -> Biz [thuc hien logic can thiet] -> Repository[tong hop va tranform du lieu] -> Storage[tang giao tiep du lieu]

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}
func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
