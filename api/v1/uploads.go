package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	orm "go_restful_api_playground/database"
	Models "go_restful_api_playground/models"
	"net/http"
	"path/filepath"
)

// UploadApp godoc
// @Summary Create Upload File
// @Schemes
// @Description Create Upload File
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {string} string "{"message": "Your Upload Success"}"
// @Router /upload [post]
// @Security BearerAuth
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	UUIDv4 := uuid.Must(uuid.NewV4()).String()
	extension := filepath.Ext(file.Filename)
	newFilename := UUIDv4 + extension
	err := c.SaveUploadedFile(file, "./media/"+newFilename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
		})
		return
	}

	getUser := GetUser(c)

	// 強制轉型 interface to struct
	user, _ := getUser.(UserInfoBody)

	var existUser Models.User
	if result := orm.Db.First(&existUser, user.ID); result.Error != nil {
		return
	}

	modelFile := Models.File{UserId: existUser.ID, FileName: newFilename}
	if result := orm.Db.Create(&modelFile); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Upload Success",
	})

}
