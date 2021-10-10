package handler

import (
	"net/http"

	"online-store/api/user/entity"
	"online-store/lib/presenter"

	"github.com/gin-gonic/gin"

	validate "github.com/go-playground/validator/v10"
)

type LoginContract interface {
	Login(ctx *gin.Context)
}

func (h *Handler) Login(ctx *gin.Context) {
	param := entity.User{}
	if err := ctx.Bind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	v := validate.New()
	if err := v.Struct(param); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	token, err := h.Usecase.Login(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been login",
		"token":   token,
	})
}
