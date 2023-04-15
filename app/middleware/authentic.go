package middleware

import (
	"MyGram/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(Context *gin.Context) {
		ReceivedToken, err := helper.VerficationTokenizer(Context)
		_ = ReceivedToken
		if err != nil {
			Context.JSON(http.StatusUnauthorized, gin.H{
				"Status":  http.StatusUnauthorized,
				"Message": err.Error(),
			})
			return
		}
		Context.Set("UserData", ReceivedToken)
		Context.Next()
	}
}
