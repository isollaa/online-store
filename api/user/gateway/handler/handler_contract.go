package handler

import "online-store/api/user/usecase"

type HandlerContract interface {
	LoginContract
	UserContract
}

type Handler struct {
	Usecase usecase.UsecaseContract
	Name    string
}

func NewHandler(uc usecase.UsecaseContract) HandlerContract {
	return &Handler{
		Usecase: uc,
		Name:    "User",
	}
}
