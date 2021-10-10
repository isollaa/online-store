package usecase

import (
	"online-store/api/user/repository"

	"gorm.io/gorm"
)

type UsecaseContract interface {
	LoginContract
	UserContract
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
