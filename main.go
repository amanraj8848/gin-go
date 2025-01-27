package main


import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/ping", pingRespone)

	router.Run(":8000")
}

func pingResponse(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}