package controller

import (
	"challenge/src/service"

	"github.com/gin-gonic/gin"
)

type ImportController struct {
	importService *service.ImportService
}

func ImportControllerFactory(importService *service.ImportService) *ImportController {
	return &ImportController{importService}
}

func (c *ImportController) FindAll(ctx *gin.Context) {
	imports := c.importService.FindAll()

	ctx.JSON(200, imports)
}
