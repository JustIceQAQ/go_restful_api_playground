package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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
	err := c.SaveUploadedFile(file, "./media/"+UUIDv4+extension)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your Upload Success",
	})

}
