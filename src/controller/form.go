package controller

import (
	"challenge/src/model"
	"challenge/src/service"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormController struct {
	fileService        *service.FileService
	transactionService *service.TransactionService
}

func FormControllerFactory(fileService *service.FileService, transactionService *service.TransactionService) *FormController {
	return &FormController{fileService, transactionService}
}

func (c *FormController) Upload(ctx *gin.Context) {
	handler, fileBytes, err := c.readFile(ctx)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Failed to read file"))
		return
	}

	var models []model.Transaction
	if err = c.fileService.ParseLines(string(fileBytes), &models); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Failed to parse file"))
		return
	}

	if err = c.transactionService.Save(models); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("%s Uploaded!", handler.Filename))
}

func (*FormController) readFile(ctx *gin.Context) (*multipart.FileHeader, []byte, error) {
	req := ctx.Request
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		return nil, nil, err
	}

	log.Printf("name: %s", handler.Filename)
	log.Printf("size: %d", handler.Size)

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	return handler, fileBytes, nil
}
