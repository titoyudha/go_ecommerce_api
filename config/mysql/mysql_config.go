package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (server Server) MysqlConnection(config Config) {
	var err error

	config.DBHost = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBPort = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connecting to database")
	}
	fmt.Println("Connected to mysql")
}
