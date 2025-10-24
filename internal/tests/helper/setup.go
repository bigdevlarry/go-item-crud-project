package helper

import (
	"go-test/internal/handlers"
	"go-test/internal/models/validators"
	"go-test/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupReadRouter() (*gin.Engine, store.ItemsStorage) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register custom validators (same as main.go)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("itemtype", validators.ValidateItemType)
		v.RegisterValidation("itemstatus", validators.ValidateItemStatus)
	}

	s := store.NewStore()
	handler := handlers.NewItemsHandler(s)
	r.GET("/items", handler.GetAll)
	r.GET("/items/:guid", handler.GetByGUID)
	r.POST("/items", handler.Create)
	r.PUT("/items/:guid", handler.Update)
	r.DELETE("/items/:guid", handler.Delete)
	return r, s
}
