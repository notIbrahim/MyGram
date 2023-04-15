package helper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const keys = "Golang$6"

func GenKeys(ID uint, email string, Password string) string {
	FetchData := jwt.MapClaims{
		"ID":       int(ID),
		"email":    email,
		"Password": Password,
	}

	ParseTokenizer := jwt.NewWithClaims(jwt.SigningMethodHS256, FetchData)
	SignatureToken, _ := ParseTokenizer.SignedString([]byte(keys))
	return SignatureToken
}

func VerficationTokenizer(Context *gin.Context) (interface{}, error) {
	ErrResponse := errors.New("Procced With Sign-in")
	HeaderToken := Context.Request.Header.Get("Authorization")
	KeyBearer := strings.HasPrefix(HeaderToken, "Bearer")

	if !KeyBearer {
		return nil, ErrResponse
	}

	// _ this check if _, ok
	Tokenizer := strings.Split(HeaderToken, " ")[1]
	FinalToken, _ := jwt.Parse(Tokenizer, func(Token *jwt.Token) (interface{}, error) {
		if _, ok := Token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrResponse
		}
		return []byte(keys), nil
	})

	if _, ok := FinalToken.Claims.(*jwt.MapClaims); !ok && !FinalToken.Valid {
		return nil, ErrResponse
	}

	return FinalToken.Claims.(jwt.MapClaims), nil
}
