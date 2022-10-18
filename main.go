package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/titoyudha/go_ecommerce_api/config/mysql"
	rdb "github.com/titoyudha/go_ecommerce_api/config/redis"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	var db *gorm.DB

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	fmt.Println("Server is running well")
	server.Run()

	mysql.MysqlConnection(db)
	rdb.ClientConnection()

}
