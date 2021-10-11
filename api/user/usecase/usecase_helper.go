package usecase

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func (u *Usecase) generateToken(id int64) (string, error) {
	token := jwt.MapClaims{
		"user_id":    id,
		"expired_at": time.Now().Add(10 * time.Minute).Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	signedToken, err := sign.SignedString([]byte(viper.GetString("signing_key")))
	if err != nil {
		return "", err
	}
	return signedToken, err
}
