package biz

// bai tap lem them
type all struct {
	createItemBiz
	UpdateItemBiz
	DeleteItemBiz
	listItemBiz
	getItemBiz
}

func NewItemUseCase() *all {
	return &all{}
}
