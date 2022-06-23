package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	config "go_restful_api_playground/configs"
	"mime/multipart"
	"net/http"
	"time"
)

var FUH = FileUploadHelper{}

type FileUploadHelper struct {
}

func (fmh *FileUploadHelper) ToLocal(c *gin.Context, input interface{}, newFilename string) (string, error) {
	localPath := "/media/" + newFilename
	formFile := input.(*multipart.FileHeader)
	err := c.SaveUploadedFile(formFile, "."+localPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
		})
		return "", err
	}
	return localPath, nil
}

func (fmh *FileUploadHelper) ToCloudinary(input interface{}) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(
		config.Cfg.Cloudinary.CloudName,
		config.Cfg.Cloudinary.ApiKey,
		config.Cfg.Cloudinary.ApiSecret,
	)
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.Cfg.Cloudinary.UploadFolder})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
