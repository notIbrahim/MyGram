package queries

import (
	"MyGram/app/entity"
	"MyGram/pkg/database"
	"MyGram/pkg/helper"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MyGram User Registration godoc
// @Summary User Registration
// @Description Well Obviously it User Registration
// @Tags tags
// @Consume ({mpfd, json})
// @Produce json
// @Param email formData string true "Your Email address registered"
// @Param password formData string true "Your password registered"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Router users/register
const APP = "application/json"

func Registered(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		panic(err)
	}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	_, _ = DB, ReceivedContent
	User := entity.User{}

	if ReceivedContent == APP {
		ResponseContext.ShouldBindJSON(&User)
	} else {
		ResponseContext.ShouldBind(&User)
	}

	QueryFind := DB.Where("email = ?", User.Email).Error

	if QueryFind != nil {
		ResponseContext.JSON(http.StatusFound, gin.H{
			"Message": "Email Already Used",
		})
	}

	if User.UniqueID == "" {
		User.UniqueID = helper.String(rand.Intn(20))
	}

	QueryCreate := DB.Debug().Create(&User)

	if QueryCreate.RowsAffected == 1 {
		ResponseContext.JSON(http.StatusCreated, gin.H{
			"Status":           http.StatusCreated,
			"Received Message": "Successfully created",
			"Message": gin.H{
				"Username": User.Username,
				"Email":    User.Email,
				"Password": User.Password,
			}})
		return
	} else {
		ResponseContext.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Error":   "Invalid JSON Data",
			"Message": err.Error(),
		})
		return
	}
}

func UserLogged(ResponseContext *gin.Context) {
	DB, err := database.Connect()
	if err != nil {
		panic(err)
	}
	ReceivedContent := helper.GetContentTypeOf(ResponseContext)
	_, _ = DB, ReceivedContent
	UserCheck := entity.User{}

	if ReceivedContent == APP {
		ResponseContext.ShouldBindJSON(&UserCheck)
	} else {
		ResponseContext.ShouldBind(&UserCheck)
	}

	// Result := map[string]interface{}{}
	PasswordEntity := helper.PasswordCheck([]byte(helper.LookupPassword(UserCheck.Password)), []byte(UserCheck.Password))
	// Query Check email are aaa@mail.com and some bcrypt stuff this is checking data again not taking data
	QueryCheck := DB.Debug().Where("email = ?", UserCheck.Email).Take(&UserCheck).Error
	if PasswordEntity && QueryCheck == nil {
		// Give them Tokenizer
		GenereateToken := helper.GenKeys(UserCheck.IDModels.ID, UserCheck.Email, UserCheck.Password)
		ResponseContext.JSON(http.StatusAccepted, gin.H{
			"Token": GenereateToken,
		})
		return
	} else {
		ResponseContext.JSON(http.StatusNotFound, gin.H{
			"Error": "Password and Email Does Not Match",
		})
		return
	}

	// Find Account Where

}
