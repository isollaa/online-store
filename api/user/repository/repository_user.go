package repository

import (
	"fmt"
	"online-store/api/user/entity"
	"online-store/lib/request"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserContract interface {
	GetList(db *gorm.DB, param request.List) (res []entity.User, total int64, err error)
	Create(db *gorm.DB, input entity.User) (err error)
	Update(db *gorm.DB, filter entity.User, input entity.User) (err error)
	Delete(db *gorm.DB, filter entity.User) (err error)
}

func (repo Repository) GetList(db *gorm.DB, param request.List) (res []entity.User, total int64, err error) {
	res = []entity.User{}
	query := db.Model(entity.User{}).Where("deleted_at IS NULL")

	page := param.PerPage * (param.Page - 1)
	if param.Search != "" {
		formattedTextSearch := "%%" + param.Search + "%%"
		query = query.Where("username LIKE ? OR password LIKE ?", formattedTextSearch, formattedTextSearch)
	}

	if err := query.Count(&total).Error; err != nil {
		return res, total, err
	}

	if total <= 0 {
		return res, total, nil
	}

	var column string
	switch param.Sort {
	case "username":
		column = "username"
	case "password":
		column = "password"
	}

	order := strings.ToUpper(param.Order)
	if order == "" {
		order = "ASC"
	}

	if column != "" {
		query = query.Order(column + " " + order)
	}

	if err := query.
		Limit(param.PerPage).
		Offset(page).
		Find(&res).Error; err != nil {
		return res, total, err
	}

	return
}

func (repo Repository) Create(db *gorm.DB, input entity.User) (err error) {
	return db.Create(&input).Error
}

func (repo Repository) Update(db *gorm.DB, filter entity.User, input entity.User) (err error) {
	if err = db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(filter).
		Where("deleted_at IS NULL").
		Take(&filter).Error; err != nil {
		err = fmt.Errorf("user with id %d not found", filter.ID)
		return
	}

	if err = db.
		Model(filter).
		Where(filter).
		Updates(map[string]interface{}{
			"username": input.Username,
			"password": input.Password,
		}).Error; err != nil {
		return
	}

	return
}

func (repo Repository) Delete(db *gorm.DB, filter entity.User) (err error) {
	db = db.Model(filter).Where(filter).Where("deleted_at IS NULL").Update("deleted_at", time.Now())
	if db.RowsAffected == 0 {
		err = fmt.Errorf("user with id %d is not exist or has been deleted", filter.ID)
		return
	}

	return db.Error
}
