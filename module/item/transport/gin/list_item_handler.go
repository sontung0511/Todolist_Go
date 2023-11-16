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

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var queryString struct {
			common.Paging
			model.Filter
		}
		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		queryString.Process()
		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)
		result, err := business.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Filter, queryString.Paging))
	}
}
