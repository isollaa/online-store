package usecase

import (
	"online-store/api/item/repository"

	"gorm.io/gorm"
)

type UsecaseContract interface {
	ItemContract
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
