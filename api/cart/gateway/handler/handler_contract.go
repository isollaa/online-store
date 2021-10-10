package handler

import "online-store/api/cart/usecase"

type HandlerContract interface {
	CartContract
}

type Handler struct {
	Usecase usecase.UsecaseContract
	Name    string
}

func NewHandler(uc usecase.UsecaseContract) HandlerContract {
	return &Handler{
		Usecase: uc,
		Name:    "Cart",
	}
}
