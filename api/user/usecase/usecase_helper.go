package usecase

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (u *Usecase) generateToken(id int64, str string) (string, error) {
	token := jwt.MapClaims{
		"user_id":    id,
		"expired_at": time.Now().Add(10 * time.Minute).Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	signedToken, err := sign.SignedString([]byte(str))
	if err != nil {
		return "", err
	}
	return signedToken, err
}
