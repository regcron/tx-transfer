package db

import (
	accountDto "com/txfer/bounded_contexts/account/dtos"
	transferDto "com/txfer/bounded_contexts/transfer/dtos"
	"com/txfer/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *configs.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	// DB = DB.Debug()
	MigrateSchema()
	fmt.Println("🚀 Connected Successfully to the Database")
}

func MigrateSchema() {
	log.Println("🚀 Migrating Database")
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(&accountDto.AccountDto{}, &transferDto.TransactionDto{})
	fmt.Println("👍 Migration complete")
}
