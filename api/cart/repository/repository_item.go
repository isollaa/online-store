package repository

import (
	"fmt"
	"online-store/api/cart/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemContract interface {
	GetItemForUpdate(db *gorm.DB, id int64) (res entity.Item, err error)
	UpdateQuantityItem(db *gorm.DB, id int64, quantity float64) (err error)
}

func (repo Repository) GetItemForUpdate(db *gorm.DB, id int64) (res entity.Item, err error) {
	if err = db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(entity.Item{ID: id}).
		Take(&res).Error; err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("item with id %d not found", id)
		return
	}

	return
}

func (repo Repository) UpdateQuantityItem(db *gorm.DB, id int64, quantity float64) (err error) {
	filter := entity.Item{
		ID: id,
	}
	if err = db.
		Model(filter).
		Where(filter).
		Update("quantity", quantity).Error; err != nil {
		return
	}

	return
}
