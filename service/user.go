package service

import (
	"log"
	"main/db"
	"main/helper"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	jwtData, _ := c.Get(helper.JwtIssuer)
	log.Println(jwtData)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "data": "1234"})
}

func RemoveFile(c *gin.Context) {
	id := c.Param("id")

	if _, err := helper.RemoveFile(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"success": true, "data": id})

}

func Upload(c *gin.Context) {
	// one file
	// file, err := c.FormFile("image")
	// upload, err := helper.SaveFile(c, file, err)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"success": false, "error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully" + upload})

	// Multiple file uploads
	form, _ := c.MultipartForm()
	images := form.File["images"]
	for _, file := range images {

		log.Println("File name", file.Filename)
		// handle each uploaded file
		// file contains information such as filename, size, and the actual file data
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}

func CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	collection, _ := db.Collection(db.UsersCollection)
	user := model.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error  df": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": result.InsertedID})
}
