package Response

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user/app/Models"
	"user/config"
)

type LoginResponseData struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func LoginResponse(user Models.User) (*LoginResponseData, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = "USER"
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	signedToken, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	res := &LoginResponseData{
		Message: "Login Success",
		Token:   signedToken,
	}

	return res, nil
}
