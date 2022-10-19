package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/titoyudha/go_ecommerce_api/config/mysql"
	rdb "github.com/titoyudha/go_ecommerce_api/config/redis"
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	fmt.Println("Server is running well")
	mysql.DBConnection()

	rdb.ClientConnection()
	server.Run()

}
