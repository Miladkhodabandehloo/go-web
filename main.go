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
	db, err := gorm.Open("postgres", "host=db_server port=5432 user=postgres dbname=app password=my-secret sslmode=disable")
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
