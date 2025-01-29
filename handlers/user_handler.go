package handlers

import(
	"net/http"
	db "www.github.com/amanraj8848/gin-go/database"
	"www.github.com/amanraj8848/gin-go/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var newuser models.User
	if err := c.ShouldBindJSON(&newuser); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.DB.Create(&newuser)
	c.JSON(http.StatusCreated, newuser)
}