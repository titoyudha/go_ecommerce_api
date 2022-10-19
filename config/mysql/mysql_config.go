package mysql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/titoyudha/go_ecommerce_api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Items{}, &model.Cart{}, &model.Payment{})

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
