package queries

import (
	"MyGram/app/entity"
	"MyGram/pkg/database"
	"MyGram/pkg/helper"
	"net/http"
	"strconv"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// This Problem will even weirder because Auto_Increment Bust up

// All I need is Authorization First and Gun it

func GetAll(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	Photo := []entity.Photo{}
	Result := []map[string]interface{}{}
	QueryCheck := DB.Debug().Select("id", "user_id", "title", "caption", "photo_url").Find(&Photo).Find(&Result)

	if QueryCheck.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "Photo not found",
		})
	}
}

func GetOne(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	Photo := []entity.Photo{}
	// Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	// UserID := uint(Userdata["ID"].(float64))
	Reference, _ := strconv.Atoi(ResponseContext.Param("ID"))
	QueryFetch := DB.Where("id = ?", Reference).Find(&Photo)

	if QueryFetch.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Photo,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "Photo not found",
		})
	}
}

func CreatePhoto(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	Params := uint(Userdata["ID"].(float64))
	Photo := entity.Photo{}
	Photo.User_ID = Params
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Photo)
	} else {
		ResponseContext.ShouldBind(&Photo)
	}

	// Bug Check
	QueryCheck := DB.Debug().Create(&Photo)

	if QueryCheck.RowsAffected == 1 {

		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": gin.H{
				"Title":    Photo.Title,
				"Caption":  Photo.Caption,
				"PhotoURL": Photo.PhotoURL,
			},
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Record Failed to Insert",
		})
	}
}

func UpdatePhoto(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram

	// Lead Into ID Photo INT but for some reason Foreign Key will be such a problem
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	UserID := uint(Userdata["ID"].(float64))
	Photo := entity.Photo{}
	Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Photo)
	} else {
		ResponseContext.ShouldBind(&Photo)
	}

	// Bug Check
	QueryCheck := DB.Model(&Photo).Where("User_ID = ?", UserID).Where("ID = ?", Params).Updates(entity.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Take(&Result)
	// QueryCheck := DB.Debug().Where("photo_id = ?", Params).Find(&Photo).Take(&Result)

	if QueryCheck.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Record Failed to Update",
		})
	}
}

func DeletePhoto(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	UserID := uint(Userdata["ID"].(float64))
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	Photo := entity.Photo{}
	// Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Photo)
	} else {
		ResponseContext.ShouldBind(&Photo)
	}

	// Bug Check
	QueryCheck := DB.Debug().Where("ID = ?", Params).Where("User_ID = ?", UserID).Delete(&Photo)

	if QueryCheck.RowsAffected == 1 {
		DB.Exec("ALTER TABLE photos AUTO_INCREMENT = 1")
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Result": "Photo Delete Success",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Failed to Delete",
		})
	}
}
