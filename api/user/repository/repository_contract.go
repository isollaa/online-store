package repository

type RepositoryContract interface {
	LoginContract
	UserContract
}

type Repository struct{}

func NewRepository() RepositoryContract {
	return &Repository{}
}
