package rest

import (
	"fmt"
	"main/conf"

	"github.com/gin-gonic/gin"
)

func Init() error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/transform", transformHandler)
	r.POST("/transform", transformHandler)

	r.GET("/generate", generateHandler)
	r.POST("/generate", generateHandler)

	r.GET("/assist", assistHandler)
	r.POST("/assist", assistHandler)

	r.Run(conf.Config.Server.Host + ":" + fmt.Sprint(conf.Config.Server.Port))

	return nil

}
