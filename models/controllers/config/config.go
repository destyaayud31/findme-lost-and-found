package config

import (
	"fmt"

	"go_api_destya/models"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

var DB *gorm.DB

func LoadConfig() {
	viper.AddConfigPath(".")    // cari file .env di root
	viper.SetConfigName(".env") // nama file tanpa ekstensi
	viper.SetConfigType("env")  // tipe file

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Error reading .env file:", err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		logrus.Fatal("Error unmarshalling config:", err)
	}

	logrus.Println("Load server successfully")
}

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	DB = db

	db.AutoMigrate(&models.Author{}, &models.Book{})

	logrus.Println("Database connected")
}
