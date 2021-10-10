package repository

import (
	"encoding/base64"
	"fmt"
	"online-store/api/user/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LoginContract interface {
	InjectAdmin(db *gorm.DB) (err error)
	Validate(db *gorm.DB, input *entity.User) (err error)
}

func (repo Repository) InjectAdmin(db *gorm.DB) (err error) {
	users := []entity.User{}
	user := entity.User{
		Username: "admin",
		Password: base64.StdEncoding.EncodeToString([]byte("admin")),
	}

	if err = db.
		Select("id").
		Where(user).
		Limit(1).
		Find(&users).Error; err != nil {
		return
	}

	if len(users) <= 0 {
		if err = db.
			Clauses(clause.Insert{Modifier: "IGNORE"}).
			Create(&user).Error; err != nil {
			return
		}
	}

	return
}

func (repo Repository) Validate(db *gorm.DB, input *entity.User) (err error) {
	if err = db.
		Where(input).
		Take(&input).Error; err != nil {
		err = fmt.Errorf("invalid username / password")
		return
	}

	return
}
