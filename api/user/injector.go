package user

import (
	"online-store/api/user/gateway/handler"
	"online-store/api/user/repository"
	"online-store/api/user/usecase"
	"online-store/config"
)

func ApiHandler() handler.HandlerContract {
	repo := repository.NewRepository()
	usecase := usecase.NewUsecase(config.Orm, repo)
	return handler.NewHandler(usecase)
}
