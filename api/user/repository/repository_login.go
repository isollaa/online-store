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

//inject tb user if it doesn't have user with username: "admin", password: "admin"
func (repo Repository) InjectAdmin(db *gorm.DB) (err error) {
	users := []entity.User{}
	user := entity.User{
		Username: "admin",
		Password: base64.StdEncoding.EncodeToString([]byte("admin")),
	}

	if err = db.
		Select("id").
		Where(user).
		Where("deleted_at IS NULL").
		Limit(1).
		Find(&users).Error; err != nil {
		return
	}

	if len(users) <= 0 {
		//do inser ignore to prefent multiply create
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
