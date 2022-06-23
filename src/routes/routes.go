package routes

import (
	"challenge/src/controller"
	"challenge/src/database"
	"challenge/src/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(cors.Default())

	db := database.DB
	fileService := service.FileServiceFactory()
	importService := service.ImportServiceFactory(db)
	transactionService := service.TransactionServiceFactory(db, importService)
	formController := controller.FormControllerFactory(fileService, transactionService)

	r.POST("/form", formController.Upload)

	r.Run(":3333")
}
