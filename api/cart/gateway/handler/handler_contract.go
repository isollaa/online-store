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
		//set static name to use as feature differentiator on response message
		Name: "Cart",
	}
}
