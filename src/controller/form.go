package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type formController struct{}

func NewFormController() *formController {
	return &formController{}
}

func (c formController) Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Printf("name: %s", file.Filename)
	log.Printf("size: %d", file.Size)

	ctx.String(http.StatusOK, fmt.Sprintf("%s Uploaded!", file.Filename))
}
