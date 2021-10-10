package repository

type RepositoryContract interface {
	CartContract
	ItemContract
}

type Repository struct{}

func NewRepository() RepositoryContract {
	return &Repository{}
}
