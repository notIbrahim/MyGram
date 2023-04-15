package middleware

import (
	"MyGram/app/entity"
	"MyGram/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(Path string) gin.HandlerFunc {
	return func(ReceivedContext *gin.Context) {
		DB, _ := database.Connect()
		Param, err := strconv.Atoi(ReceivedContext.Param("ID"))
		if err != nil {
			ReceivedContext.JSON(http.StatusBadRequest, gin.H{
				"Error":   http.StatusBadRequest,
				"Message": "Missing ID Parameter or Invalid",
			})
			return
		}
		Userdata := ReceivedContext.MustGet("UserData").(jwt.MapClaims)
		UserID := uint(Userdata["ID"].(float64))

		switch Path {
		case "users":
			// Bring Entity Users
			UserEntity := entity.User{}
			QueryChecks := DB.Where("id = ?", Param).Take(&UserEntity)
			if QueryChecks.RowsAffected != 1 && UserEntity.IDModels.ID != uint(Param) {
				ReceivedContext.JSON(http.StatusUnauthorized, gin.H{
					"Message": "Either you are unauthorized",
					"Error":   err.Error(),
				})
				return
			}
		case "socials":
			UserEntity := entity.Social{}
			QueryChecks := DB.Where("user_id = ?", Param).Take(&UserEntity)
			if QueryChecks.RowsAffected != 1 && UserEntity.IDModels.ID != UserID {
				ReceivedContext.JSON(http.StatusUnauthorized, gin.H{
					"Message": "Either you are unauthorized",
					"Error":   err.Error(),
				})
				return
			}
		case "comments":
			UserEntity := entity.User{}
			QueryChecks := DB.Where("id = ?", Param).Take(&UserEntity)
			if QueryChecks.RowsAffected != 1 && UserEntity.IDModels.ID != UserID {
				ReceivedContext.JSON(http.StatusUnauthorized, gin.H{
					"Message": "Either you are unauthorized",
					"Error":   err.Error(),
				})
				return
			}
		case "photos":
			UserEntity := entity.Photo{}
			QueryChecks := DB.Where("user_id = ?", Param).Take(&UserEntity)
			if QueryChecks.RowsAffected != 1 && UserEntity.IDModels.ID != UserID {
				ReceivedContext.JSON(http.StatusUnauthorized, gin.H{
					"Message": "Either you are unauthorized",
					"Error":   err.Error(),
				})
				return
			}

		default:
			ReceivedContext.JSON(http.StatusNotFound, gin.H{
				"Error": "Path not found",
			})
			return
		}
	}
}
