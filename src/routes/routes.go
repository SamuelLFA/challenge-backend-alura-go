package routes

import (
	"challenge/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(cors.Default())

	formController := controller.NewFormController()

	r.POST("/form", formController.Upload)

	r.Run(":3333")
}
