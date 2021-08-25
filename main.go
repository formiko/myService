package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)
import "myService/controller"

func init() {
	logFile, err := os.OpenFile("log/myService_" + time.Now().Format("20060102150405") + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
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
	router.StaticFS("/file", gin.Dir("file", true))
	router.Run(":80")
}
