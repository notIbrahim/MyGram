package queries

import (
	"MyGram/app/entity"
	"MyGram/pkg/database"
	"MyGram/pkg/helper"
	"net/http"
	"strconv"

	"strings"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Authorization First

// MyGram Socials godoc
// @Summary MyGram Function Social Media
// @Description MyGram Social Media function for check one status of a user without login known as incognito
// @Tags SocialQueries
// @Consume ({mpfd, json})
// @Produce json
// @Param ID formData int true "Social Media ID it means what iterations are you looking for"
// @Description
// @Success 302
// @Failure 404
// @Router /socials/{ID} [get]
func GetStatus(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	Socials := []entity.Social{}
	Param, _ := strconv.Atoi(ResponseContext.Param("ID"))
	Result := []map[string]interface{}{}
	QueryCheck := DB.Debug().Where("id = ?", Param).Find(&Socials).Find(&Result)

	if QueryCheck.RowsAffected != 0 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "All of Social Media not found",
		})
	}
}

// MyGram Socials godoc
// @Summary MyGram Function Social Media
// @Description MyGram Social Media function for check all user status who been registered without login known as incognito
// @Tags SocialQueries
// @Consume ({mpfd, json})
// @Produce json
// @Description
// @Success 302
// @Failure 404
// @Router /socials/ [get]
func GetStatusAll(ResponseContext *gin.Context) {
	DB, _ := database.Connect()
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	Social := []entity.Social{}
	Result := []map[string]interface{}{}
	QueryCheck := DB.Debug().Find(&Social).Find(&Result)

	if QueryCheck.RowsAffected == 1 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error":   "Record Not Found",
			"Message": "Social Media Not Found",
		})
	}
}

// MyGram Socials godoc
// @Summary MyGram Function Social Media
// @Description MyGram Social Media function for Create User but need to login if this user are really real which user have been authorize
// @Tags SocialQueries
// @Consume ({mpfd, json})
// @Produce json
// @Param Name formData string true "What your name in this social accounts should be"
// @Param LinkURL formData string true "This Link URL your input will be combine your name and this style known use as LinkedIn User Link"
// @Description
// @Success 201
// @Failure 400
// @Router /socials/Add [post]

func CreateSocialMedia(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	Params := uint(Userdata["ID"].(float64))
	Socials := entity.Social{}
	Socials.User_ID = Params
	Result := []map[string]interface{}{}

	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Socials)
	} else {
		ResponseContext.ShouldBind(&Socials)
	}
	ResultString := []string{Socials.Name, Socials.Social_Media_URL}
	FinalString := strings.Join(ResultString, "-")
	Socials.Social_Media_URL = FinalString

	QueryFinder := DB.Debug().Where("user_id = ?", Params).Find(&Socials)

	if QueryFinder.RowsAffected == 1 {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Message": "You can't have multiple accounts with same UserID",
		})
		return
	}

	// Bug Check
	QueryCheck := DB.Debug().Create(&Socials).Find(&Result)

	if QueryCheck.RowsAffected != 0 {

		ResponseContext.JSON(http.StatusCreated, gin.H{
			"Data": Result,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Record Failed to Insert",
		})
	}
}

// MyGram Socials godoc
// @Summary MyGram Function Social Media
// @Description MyGram Social Media function for Update of a user but needed as Authorization and Authorization in order to do so
// @Tags SocialQueries
// @Consume ({mpfd, json})
// @Produce json
// @Param Name formData string true "What your name in this social accounts should be"
// @Param LinkURL formData string true "This Link URL your input will be combine your name and this style known use as LinkedIn User Link"
// @Success 202
// @Failure 400
// @Router /socials/Edit/{ID} [put]

func UpdateSocialMedia(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram

	// Lead Into ID Socials INT but for some reason Foreign Key will be such a problem
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	UserID := uint(Userdata["ID"].(float64))
	Socials := entity.Social{}
	Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Socials)
	} else {
		ResponseContext.ShouldBind(&Socials)
	}
	ResultString := []string{Socials.Name, Socials.Social_Media_URL}
	FinalString := strings.Join(ResultString, "-")
	Socials.Social_Media_URL = FinalString

	// Bug Check
	QueryCheck := DB.Model(&Socials).Where("User_ID = ?", UserID).Where("ID = ?", Params).Updates(entity.Social{Name: Socials.Name, Social_Media_URL: Socials.Social_Media_URL}).Take(&Result)
	// QueryCheck := DB.Debug().Where("Socials_id = ?", Params).Find(&Socials).Take(&Result)

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

// MyGram Socials godoc
// @Summary MyGram Function Social Media
// @Description MyGram Social Media function for Delete of a user
// @Tags SocialQueries
// @Consume ({mpfd, json})
// @Produce json
// @Param ID formData int true "Social Media ID it means what iterations are you looking for"
// @Description
// @Success 202
// @Failure 400
// @Router /socials/Delete/{ID} [delete]

func DeleteSocialMedia(ResponseContext *gin.Context) {
	Userdata := ResponseContext.MustGet("UserData").(jwt.MapClaims)
	DB, _ := database.Connect()
	ReceivedContext := helper.GetContentTypeOf(ResponseContext)
	// 	// If Model going to be Brand Twitter then its easy then if it Brand Instagram i need to authorize but technically
	// 	// if it link then its almost same as Twitter and Instagram
	UserID := uint(Userdata["ID"].(float64))
	Params, _ := strconv.Atoi(ResponseContext.Param("ID"))
	Socials := entity.Social{}
	// Result := map[string]interface{}{}
	if ReceivedContext == APP {
		ResponseContext.ShouldBindJSON(&Socials)
	} else {
		ResponseContext.ShouldBind(&Socials)
	}

	// Bug Check
	QueryCheck := DB.Debug().Where("ID = ?", Params).Where("User_ID = ?", UserID).Delete(&Socials)

	if QueryCheck.RowsAffected == 1 {
		DB.Exec("ALTER TABLE Socialss AUTO_INCREMENT = 1")
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Result": "Account deleted",
		})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid Request",
			"Message": "Failed to Delete",
		})
	}
}

// This Problem will even weirder because Auto_Increment Bust up

// All I need is Authorization First and Gun it
