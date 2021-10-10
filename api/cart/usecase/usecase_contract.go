package usecase

import (
	"online-store/api/cart/repository"

	"gorm.io/gorm"
)

type UsecaseContract interface {
	CartContract
}

type Usecase struct {
	DB   *gorm.DB
	Repo repository.RepositoryContract
}

func NewUsecase(db *gorm.DB, repo repository.RepositoryContract) UsecaseContract {
	return &Usecase{
		DB:   db,
		Repo: repo,
	}
}
