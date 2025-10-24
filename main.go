package main

import (
	"go-test/internal/handlers"
	"go-test/internal/models/validators"
	"go-test/internal/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("itemtype", validators.ValidateItemType)
		if err != nil {
			log.Fatal("Failed to register itemtype validator:", err)
		}
		err = v.RegisterValidation("itemstatus", validators.ValidateItemStatus)
		if err != nil {
			log.Fatal("Failed to register itemstatus validator:", err)
		}
	}

	s := store.NewStore()
	h := handlers.NewItemsHandler(s)

	r.GET("/items", h.GetAll)
	r.GET("/items/:guid", h.GetByGUID)
	r.POST("/items", h.Create)
	r.PUT("/items/:guid", h.Update)
	r.DELETE("/items/:guid", h.Delete)

	err := r.Run()
	if err != nil {
		return
	}
}
