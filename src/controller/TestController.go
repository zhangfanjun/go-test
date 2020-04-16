package controller

import (
	"github.com/gin-gonic/gin"
	"test/src/service"
	"test/src/util"
)

func TestController(engine *gin.Engine) {
	//请求的接口和处理body合并
	engine.GET("/test", func(ctx *gin.Context) {
		util.Response(200, "ok", nil, ctx)
	})
	engine.GET("/test/data/:id", service.GetDtataById)
}
