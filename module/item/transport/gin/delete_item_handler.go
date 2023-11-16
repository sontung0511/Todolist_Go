package ginitem

import (
	"Todolist_Go/common"
	"Todolist_Go/module/item/biz"
	"Todolist_Go/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteItems(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewUDeleteItemBiz(store)

		if err := business.DeleteItemById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
