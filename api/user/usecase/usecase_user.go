package usecase

import (
	"encoding/base64"
	"math"
	"online-store/api/user/entity"
	"online-store/lib/presenter"
	"online-store/lib/request"
)

type UserContract interface {
	GetList(param request.List) (res []entity.User, meta presenter.Meta, err error)
	Create(param entity.User) (err error)
	Update(id int64, param entity.User) (err error)
	Delete(id int64) (err error)
}

func (u Usecase) GetList(param request.List) (res []entity.User, meta presenter.Meta, err error) {
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

func (u Usecase) Create(param entity.User) (err error) {
	db := u.DB.Begin()
	//encode inputed password cause stored password is pares in base64
	param.Password = base64.StdEncoding.EncodeToString([]byte(param.Password))

	if err = u.Repo.Create(db, param); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}

func (u Usecase) Update(id int64, param entity.User) (err error) {
	db := u.DB.Begin()
	//encode inputed password cause stored password is pares in base64
	param.Password = base64.StdEncoding.EncodeToString([]byte(param.Password))

	if err = u.Repo.Update(db, entity.User{ID: id}, param); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}

func (u Usecase) Delete(id int64) (err error) {
	db := u.DB.Begin()

	if err = u.Repo.Delete(db, entity.User{ID: id}); err != nil {
		db.Rollback()
		return
	}

	return db.Commit().Error
}
