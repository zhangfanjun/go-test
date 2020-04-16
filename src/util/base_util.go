package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/src/model/vo"
	"time"
)

func Response(code int32, msg string, content interface{}, context *gin.Context) {
	responseVO := vo.ResponseVO{Code: code, UserMsg: msg, Timestamp: time.Now().Unix() * 1000, Content: content}
	context.PureJSON(http.StatusOK, &responseVO)
	context.Writer.Flush()
}
