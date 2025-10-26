package tests

import (
	"go-test/backend/domain/validators"
	"go-test/backend/handlers"
	"go-test/backend/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupReadRouter() (*gin.Engine, repository.ItemsStorage) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("itemtype", validators.ValidateItemType)
		v.RegisterValidation("itemstatus", validators.ValidateItemStatus)
		v.RegisterValidation("sortcode", validators.ValidateSortCode)
	}

	s := repository.NewStore()
	handler := handlers.NewItemsHandler(s)
	r.GET("/items", handler.GetAll)
	r.GET("/items/:guid", handler.GetByGUID)
	r.POST("/items", handler.Create)
	r.PUT("/items/:guid", handler.Update)
	r.DELETE("/items/:guid", handler.Delete)
	return r, s
}
