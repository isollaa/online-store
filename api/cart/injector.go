package cart

import (
	"online-store/api/cart/gateway/handler"
	"online-store/api/cart/repository"
	"online-store/api/cart/usecase"
	"online-store/config"
)

//inject repo n usecase into handler
func ApiHandler() handler.HandlerContract {
	repo := repository.NewRepository()
	usecase := usecase.NewUsecase(config.Orm, repo)
	return handler.NewHandler(usecase)
}
