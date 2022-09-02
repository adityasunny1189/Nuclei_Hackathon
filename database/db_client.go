package database

import (
	"log"

	"github.com/adityasunny1189/FitnFine/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}
	log.Println("connected to database")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserDetail{})
	DB.AutoMigrate(&models.Plan{})
	DB.AutoMigrate(&models.PlanDetail{})
	DB.AutoMigrate(&models.Exercise{})
	DB.AutoMigrate(&models.Goal{})
	DB.AutoMigrate(&models.Wallet{})
	log.Println("database migration completed")
}