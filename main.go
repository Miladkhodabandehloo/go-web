package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log *logrus.Logger
var db *gorm.DB

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD")))

	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(db)
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/favicon.ico", gin.WrapH(http.NotFoundHandler()))
	router.StaticFS("/statics/", gin.Dir("static", true))
	log.Fatalln(router.Run(":8000"))
}

func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hello world."})
	return
}
