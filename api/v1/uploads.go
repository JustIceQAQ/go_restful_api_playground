package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go_restful_api_playground/configs"
	orm "go_restful_api_playground/database"
	Models "go_restful_api_playground/models"
	utils "go_restful_api_playground/utils"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

func FilePathFix(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}

// UploadApp godoc
// @Summary List Files
// @Schemes
// @Description GET List Files
// @Tags Upload
// @Accept json
// @Produce json
// @Success 200 {array} Models.File
// @Router /files/ [get]
// @Security BearerAuth
func FileList(c *gin.Context) {
	var files []Models.File
	FilePathFix(c)

	if result := orm.Db.Find(&files); result.Error != nil {
		return
	}

	for index, file := range files {
		if !strings.Contains(file.Uri, "http") {
			files[index].Uri = FilePathFix(c) + file.Uri
			//file.Uri = FilePathFix(c) + file.Uri
		}
	}

	c.JSON(http.StatusOK, &files)
}

// UploadApp godoc
// @Summary Create Upload File
// @Schemes
// @Description Create Upload File
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Param filename formData string true "filename"
// @Success 200 {string} string "{"message": "Your Upload Success"}"
// @Router /files/ [post]
// @Security BearerAuth
func UploadFile(c *gin.Context) {
	var uploadUrl string
	var uploadErr error
	formFile, _ := c.FormFile("file")
	formFileName := c.PostForm("filename")
	UUIDv4 := uuid.Must(uuid.NewV4()).String()
	extension := filepath.Ext(formFile.Filename)
	uuidName := UUIDv4 + extension

	switch configs.Cfg.UploadTo {
	case "local":
		uploadUrl, uploadErr = utils.FUH.ToLocal(c, formFile, uuidName)
		if uploadErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": uploadErr.Error(),
			})
			return
		}
	case "cloudinary":
		openFile, _ := formFile.Open()
		uploadUrl, uploadErr = utils.FUH.ToCloudinary(openFile)
		if uploadErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": uploadErr.Error(),
			})
			return
		}

	}
	fmt.Println(uploadUrl)

	getUser := GetUser(c)

	// 強制轉型 interface to struct
	user, _ := getUser.(UserInfoBody)

	var existUser Models.User
	if result := orm.Db.First(&existUser, user.ID); result.Error != nil {
		return
	}

	modelFile := Models.File{UserId: existUser.ID, FileName: formFileName, Uri: uploadUrl}
	if result := orm.Db.Create(&modelFile); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Upload Success",
	})

}
