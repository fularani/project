package db

import(
	"log"
	"project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

func Init() *gorm.DB {
	
	dbname :="postgres"
	username :="postgres"
	password :="postgres"
	host :="0.0.0.0"
    port :="5432"
	sslMode :="disable"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, username, password, dbname, port, sslMode,
	)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})

    return db
}