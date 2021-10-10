package repository

type RepositoryContract interface {
	ItemContract
}

type Repository struct{}

func NewRepository() RepositoryContract {
	return &Repository{}
}
