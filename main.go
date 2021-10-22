package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/formiko/myService/controller/moonController"
	"os"
	"time"
	"github.com/formiko/myService/controller"
)

func init() {
	// todo 0777 ?
	err := os.Mkdir("log", 0777)
	if err != nil {
		fmt.Printf("mkdir fail with err:[%+v]\n", err)
	}
	logFile, err := os.OpenFile("log/myService_" + time.Now().Format("20060102150405") + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:[%+v]\n", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/time")
	{
		v1.GET("", controller.StampToTime)
	}
	moonGroup := router.Group("moon")
	{
		moonGroup.GET("/admire", moonController.Admire)
		moonGroup.GET("/change", moonController.ChangeMoon)
	}
	router.Run(":80")
}
