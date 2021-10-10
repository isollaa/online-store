package usecase

import (
	"fmt"
	"math"
	"online-store/api/cart/entity"
	"online-store/lib/presenter"
	"online-store/lib/request"
)

type CartContract interface {
	GetList(param request.List) (res []entity.Cart, meta presenter.Meta, err error)
	Create(param entity.Cart) (err error)
	Void(id int64) (err error)
}

func (u Usecase) GetList(param request.List) (res []entity.Cart, meta presenter.Meta, err error) {
	res, total, err := u.Repo.GetList(u.DB, param)
	if err != nil {
		return
	}

	meta = presenter.Meta{
		TotalData: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(param.PerPage))),
		Page:      param.Page,
		PerPage:   param.PerPage,
	}

	return
}

func (u Usecase) Create(param entity.Cart) (err error) {
	db := u.DB.Begin()

	item, err := u.Repo.GetItemForUpdate(db, param.ItemID)
	if err != nil {
		db.Rollback()
		return
	}

	if item.Quantity <= 0 {
		err = fmt.Errorf("quantity item with id %d is 0", param.ItemID)
		db.Rollback()
		return
	}

	param.ItemName = item.Name
	if err = u.Repo.Create(db, param); err != nil {
		db.Rollback()
		return
	}

	newQuantity := item.Quantity - param.Quantity
	if err = u.Repo.UpdateQuantityItem(db, param.ItemID, newQuantity); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}

func (u Usecase) Void(id int64) (err error) {
	db := u.DB.Begin()

	param, err := u.Repo.Get(db, id)
	if err != nil {
		db.Rollback()
		return
	}

	item, err := u.Repo.GetItemForUpdate(db, param.ItemID)
	if err != nil {
		db.Rollback()
		return
	}

	newQuantity := item.Quantity + param.Quantity
	if err = u.Repo.UpdateQuantityItem(db, param.ItemID, newQuantity); err != nil {
		db.Rollback()
		return
	}

	if err = u.Repo.Delete(db, entity.Cart{ID: id}); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}
