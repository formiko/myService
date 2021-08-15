package main

import "github.com/gin-gonic/gin"
import "myService/controller"

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/time")
	{
		v1.GET("", controller.StampToTime)
	}
	router.Run(":80")
}
