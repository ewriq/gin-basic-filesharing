package main

import (
	"github.com/gin-gonic/gin"
	"anonim-fileshare/Handlers"
)

func main(){
	r := gin.Default()

	r.GET("/",Handlers.Home)
	r.POST("/upload",Handlers.Upload)
	r.Static("/dw/", "./shared")

	r.Run("127.0.0.1:5000")
}