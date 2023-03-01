package persistence

import (
	"go-minitwit/src/application"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDbConnection() *gorm.DB {
	dsn := "host=minitwit_db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	if db == nil{
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		
		if err != nil {
			panic("Failed to connect to database")
		}
		if db != nil{
			return db
		}
	}

	return db
}

func ConfigurePersistence() {
	db := GetDbConnection()

	applyMigrations(db)
	seed(db)
}

func applyMigrations(db *gorm.DB) {
	db.AutoMigrate(&application.User{})
	db.AutoMigrate(&application.Message{})
}
