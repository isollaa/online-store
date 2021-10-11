package usecase

import (
	"encoding/base64"
	"online-store/api/user/entity"
	"online-store/config"
)

type LoginContract interface {
	Login(user entity.User) (res string, err error)
}

func (u Usecase) Login(user entity.User) (res string, err error) {
	//encode inputed password cause stored password is pares in base64
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))

	if err = u.Repo.InjectAdmin(u.DB); err != nil {
		return
	}

	if err = u.Repo.Validate(u.DB, &user); err != nil {
		return
	}

	res, err = u.generateToken(user.ID, config.Token)
	if err != nil {
		return
	}

	return
}
