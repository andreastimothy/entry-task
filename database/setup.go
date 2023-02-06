package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() error {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	info := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=require",
		os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err = gorm.Open(postgres.Open(info), &gorm.Config{Logger: getLogger()})

	if err != nil {
		panic(err)
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
