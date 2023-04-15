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

func CreateComment(ResponseContext *gin.Context) {
	// Creating Comment need some joins because i have to create
	// Table User and Photo
	// Bring User_ID and Photo_ID

	// Bug Issues :
	// Revamp Gorm of Photo
	// 1. Adding is_active set because its weird having multiple photo but not actually what they uses
	type Joined struct {
		UserID uint
		ID     uint
	}
	DB, _ := database.Connect()

	// Both having same User_ID
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	UserID := uint(Userdata["ID"].(float64))
	Temps := Joined{}
	Comment := entity.Comment{}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	QueryCheck := DB.Debug().Find(&Comment).Error

	if QueryCheck != nil {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Message": "Comment Not Found",
		})
		// it means error happend
		return
	}

	Joins := DB.Debug().Model(&entity.User{}).Select("users.id, photos.user_id, photos.id ").Joins("left join photos on users.id = photos.user_id").Where("user_id = ? ", UserID).Take(&Temps)

	if ReceivedContent == APP {
		ResponseContext.ShouldBindJSON(&Comment)
	} else {
		ResponseContext.ShouldBind(&Comment)
	}

	Comment.User_ID = Temps.UserID
	Comment.Photo_ID = Temps.ID

	// Now adding some into database
	QueryCreate := DB.Debug().Create(&Comment)

	if Joins.Error == nil && QueryCreate.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Result": "Data",
			"Data":   Comment,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Failed to Insert",
		})
		return
	}
}

func GetOneComment(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	Comments := []entity.Comment{}
	// Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	// UserID := uint(Userdata["ID"].(float64))
	Reference, _ := strconv.Atoi(ResponseContext.Param("ID"))
	QueryFetch := DB.Where("id = ?", Reference).Find(&Comments)

	if QueryFetch.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Comments,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "Comment not found",
		})
	}
}

func GetAllComments(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	Comments := []entity.Comment{}
	Result := []map[string]interface{}{}
	// Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	// UserID := uint(Userdata["ID"].(float64))
	QueryFetch := DB.Debug().Select("id", "user_id", "message").Find(&Comments).Find(&Result)

	if QueryFetch.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "Comment not found",
		})
	}
}

func UpdateComment(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram

	// Lead Into ID Photo INT but for some reason Foreign Key will be such a problem
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	UserID := uint(Userdata["ID"].(float64))
	Comments := entity.Comment{}
	Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Comments)
	} else {
		ResponseContext.ShouldBind(&Comments)
	}

	// Bug Check
	QueryCheck := DB.Model(&Comments).Where("User_ID = ?", UserID).Where("ID = ?", Params).Updates(entity.Comment{Message: Comments.Message}).Take(&Result)
	// QueryCheck := DB.Debug().Where("photo_id = ?", Params).Find(&Photo).Take(&Result)

	if QueryCheck.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusAccepted, gin.H{
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

func DeleteComment(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	UserID := uint(Userdata["ID"].(float64))
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	Comments := entity.Comment{}
	// Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Comments)
	} else {
		ResponseContext.ShouldBind(&Comments)
	}

	// Bug Check
	QueryCheck := DB.Debug().Where("ID = ?", Params).Where("User_ID = ?", UserID).Delete(&Comments)

	if QueryCheck.RowsAffected == 1 {
		DB.Exec("ALTER TABLE photos AUTO_INCREMENT = 1")
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Result": "Comment deleted successfully",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Failed to Delete",
		})
	}
}
