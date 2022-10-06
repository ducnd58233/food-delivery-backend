package main

import (
	"food-delivery/components"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load file .env
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalln("Could not load .env file")
	}

	// Connect to db
	mySqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	if !ok {
		log.Fatalln("Missing MySQL connection string.")
	}

	dsn := mySqlConnStr
	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errCon != nil {
		log.Fatalln(errCon)
	}

	log.Println("Connected:", db)

	router := gin.Default()
	// router.Use(middleware.Recover())

	appCtx := component.NewAppContext(db)

	mainRoute(router, appCtx)

	router.Run(":6000")
}
