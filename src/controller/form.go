package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type formController struct{}

func NewFormController() *formController {
	return &formController{}
}

func (c formController) Upload(ctx *gin.Context) {
	req := ctx.Request
	req.ParseMultipartForm(32 << 20)
	file, handler, _ := req.FormFile("file")

	log.Printf("name: %s", handler.Filename)
	log.Printf("size: %d", handler.Size)

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(fileBytes))

	ctx.String(http.StatusOK, fmt.Sprintf("%s Uploaded!", handler.Filename))
}
