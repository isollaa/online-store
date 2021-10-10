package handler

import "online-store/api/item/usecase"

type HandlerContract interface {
	ItemContract
}

type Handler struct {
	Usecase usecase.UsecaseContract
	Name    string
}

func NewHandler(uc usecase.UsecaseContract) HandlerContract {
	return &Handler{
		Usecase: uc,
		Name:    "Item",
	}
}
