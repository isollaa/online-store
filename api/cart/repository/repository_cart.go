package repository

import (
	"online-store/api/cart/entity"
	"online-store/lib/request"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CartContract interface {
	Get(db *gorm.DB, id int64) (res entity.Cart, err error)
	GetList(db *gorm.DB, param request.List) (res []entity.Cart, total int64, err error)
	Create(db *gorm.DB, input entity.Cart) (err error)
	Delete(db *gorm.DB, filter entity.Cart) (err error)
}

func (repo Repository) Get(db *gorm.DB, id int64) (res entity.Cart, err error) {
	if err = db.
		Where(entity.Cart{ID: id}).
		Take(&res).Error; err != nil {
		return
	}

	return
}

func (repo Repository) GetList(db *gorm.DB, param request.List) (res []entity.Cart, total int64, err error) {
	res = []entity.Cart{}
	query := db.Model(entity.Cart{}).Where("deleted_at IS NULL")

	page := param.PerPage * (param.Page - 1)
	if param.Search != "" {
		formattedTextSearch := "%%" + param.Search + "%%"
		query = query.Where("user_id LIKE ? OR item_name LIKE ? OR quantity LIKE ?", formattedTextSearch, formattedTextSearch, formattedTextSearch)
	}

	if err := query.Count(&total).Error; err != nil {
		return res, total, err
	}

	if total <= 0 {
		return res, total, nil
	}

	var column string
	switch param.Sort {
	case "user_id":
		column = "user_id"
	case "item_name":
		column = "item_name"
	case "quantity":
		column = "quantity"
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

func (repo Repository) Create(db *gorm.DB, input entity.Cart) (err error) {
	return db.Create(&input).Error
}

func (repo Repository) Delete(db *gorm.DB, filter entity.Cart) (err error) {
	return db.Model(filter).Where(filter).Update("deleted_at", time.Now()).Error
}
