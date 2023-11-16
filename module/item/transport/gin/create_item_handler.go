package ginitem

import (
	"Todolist_Go/common"
	"Todolist_Go/module/item/biz"
	"Todolist_Go/module/item/model"
	"Todolist_Go/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData model.TodoItemCreation
		if err := c.ShouldBind(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)
		if err := business.CreateNewItem(c.Request.Context(), &itemData); err != nil {
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))
	}
}
