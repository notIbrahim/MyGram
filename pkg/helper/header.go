package helper

import "github.com/gin-gonic/gin"

func GetContentTypeOf(Context *gin.Context) string {
	return Context.Request.Header.Get("Content-Type")
}
