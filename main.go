package main

import (
	"fmt"
	"log"

	"go-url-shortener/internal/controllers"
	"go-url-shortener/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

var urlController controllers.UrlController

func init() {
	env, env_err := godotenv.Read()

	if env_err != nil {
		log.Fatal("couldn't load env", env_err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env["POSTGRES_URL"],
		env["POSTGRES_USER"],
		env["POSTGRES_PASSWORD"],
		env["POSTGRES_DB_NAME"],
		env["POSTGRES_PORT"],
	)

	connection, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

	if migrationErr := connection.AutoMigrate(&model.Url{}); migrationErr != nil {
		log.Fatal("failed to migrate database: ", migrationErr)
	}

	db = connection
}

func main() {
	app := gin.Default()

	urlController = controllers.NewUrlController(db)

	app.GET("/:slug", urlController.FindUrlBySlug)
	app.POST("/shorten", urlController.Create)

	app.Run(":1040")
}
