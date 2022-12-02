package database

import (
	"github.com/mvpoyatt/go-api/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	DbName   string `mapstructure:"dbname"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	HostName string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

var Db *gorm.DB

func Connect(configs DbConfig) {
	dsn := configs.UserName + ":" + configs.Password + "@tcp(" + configs.HostName + ":" + configs.Port + ")/" + configs.DbName + "?parseTime=true"

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Panicf("Failed to connect to database: %w", err)
	} else {
		logger.Log.Info("Databse connection succeeded")
	}

	// Automigrate all types to keep database up to date
	Db.AutoMigrate(&User{})
}
