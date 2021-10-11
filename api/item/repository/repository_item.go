package repository

import (
	"online-store/api/item/entity"
	"online-store/lib/request"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemContract interface {
	GetList(db *gorm.DB, param request.List) (res []entity.Item, total int64, err error)
	Create(db *gorm.DB, input entity.Item) (err error)
	Update(db *gorm.DB, filter entity.Item, input entity.Item) (err error)
	Delete(db *gorm.DB, filter entity.Item) (err error)
}

func (repo Repository) GetList(db *gorm.DB, param request.List) (res []entity.Item, total int64, err error) {
	res = []entity.Item{}
	query := db.Model(entity.Item{}).Where("deleted_at IS NULL")

	page := param.PerPage * (param.Page - 1)
	if param.Search != "" {
		formattedTextSearch := "%%" + param.Search + "%%"
		query = query.Where("name LIKE ? OR quantity LIKE ?", formattedTextSearch, formattedTextSearch)
	}

	if err := query.Count(&total).Error; err != nil {
		return res, total, err
	}

	if total <= 0 {
		return res, total, nil
	}

	var column string
	switch param.Sort {
	case "name":
		column = "name"
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

func (repo Repository) Create(db *gorm.DB, input entity.Item) (err error) {
	return db.Create(&input).Error
}

func (repo Repository) Update(db *gorm.DB, filter entity.Item, input entity.Item) (err error) {
	//use select for update to lock selected data, so other transaction wont be access on the same time
	if err = db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(filter).
		Take(&filter).Error; err != nil {
		return
	}

	if err = db.
		Model(filter).
		Where(filter).
		Updates(map[string]interface{}{
			"name":     input.Name,
			"quantity": input.Quantity,
		}).Error; err != nil {
		return
	}

	return
}

func (repo Repository) Delete(db *gorm.DB, filter entity.Item) (err error) {
	return db.Model(filter).Where(filter).Update("deleted_at", time.Now()).Error
}
