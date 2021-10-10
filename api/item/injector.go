package item

import (
	"online-store/api/item/gateway/handler"
	"online-store/api/item/repository"
	"online-store/api/item/usecase"
	"online-store/config"
)

func ApiHandler() handler.HandlerContract {
	repo := repository.NewRepository()
	usecase := usecase.NewUsecase(config.Orm, repo)
	return handler.NewHandler(usecase)
}
