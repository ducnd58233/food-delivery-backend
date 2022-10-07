package main

import (
	"food-delivery/components"
	"food-delivery/components/uploadprovider"
	middleware "food-delivery/middlewares"
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

	s3Bucketname, ok := os.LookupEnv("S3_BUCKET_NAME")
	if !ok {
		log.Fatalln("Missing S3 Bucket Name string.")
	}

	s3Region, ok := os.LookupEnv("S3_REGION")
	if !ok {
		log.Fatalln("Missing S3 Region string.")
	}

	s3APIKey, ok := os.LookupEnv("S3_ACCESS_KEY")
	if !ok {
		log.Fatalln("Missing S3 API Key string.")
	}

	s3SecretKey, ok := os.LookupEnv("S3_SECRET_KEY")
	if !ok {
		log.Fatalln("Missing S3 Secret Key string.")
	}

	s3Domain, ok := os.LookupEnv("S3_DOMAIN")
	if !ok {
		log.Fatalln("Missing S3 Domain string.")
	}

	s3Provider := uploadprovider.NewS3Provider(s3Bucketname,s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := component.NewAppContext(db, s3Provider)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))

	mainRoute(router, appCtx)

	router.Run(":6000")
}
