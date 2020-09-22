package controller

import (
	"github.com/gin-gonic/gin"
	"goBackendExample/internal/controller/database"
	"goBackendExample/internal/model"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetPostgre()
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success"})
}

func GetAllUser(c *gin.Context) {
	db := database.GetPostgre()
	var users []model.User
	result := db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success", "result": result})
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	db := database.GetPostgre()
	var user model.User
	result := db.First(&user, userId)

	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success", "result": result})
}

func UpdateUser(c *gin.Context) {
	intId, _ := strconv.Atoi(c.Param("id"))
	id := uint(intId)
	cond := model.User{DbBasicModel: model.DbBasicModel{ID: id}}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetPostgre()
	result := db.Model(&cond).Updates(user)
	c.JSON(http.StatusOK, gin.H{"errCode": 0, "errMessage": "success", "result": result})
}
