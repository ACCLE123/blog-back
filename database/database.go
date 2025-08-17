package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	// dsn := "host=localhost dbname=db_blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=127.0.0.1 port=5432 user=root password=123456 dbname=db_blog sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established successfully")
	return DB
}
