package handler

import (
	"fmt"
	"net/http"
	"online-store/api/item/entity"
	"online-store/lib/binder"
	"online-store/lib/presenter"
	"online-store/lib/validator"

	"github.com/gin-gonic/gin"

	validate "github.com/go-playground/validator/v10"
)

type ItemContract interface {
	// Get(ctx *gin.Context)
	GetList(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func (h *Handler) GetList(ctx *gin.Context) {
	param, err := binder.ValidateRequestList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}
	res, meta, err := h.Usecase.GetList(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.List{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_GET_LIST, h.Name),
		Meta:    meta,
		Data:    res,
	})
}

func (h *Handler) Create(ctx *gin.Context) {
	param := entity.Item{}
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

	if err := h.Usecase.Create(param); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_CREATE, h.Name),
	})
}

func (h *Handler) Update(ctx *gin.Context) {
	id, err := validator.PathInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	param := entity.Item{}
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

	if err := h.Usecase.Update(id, param); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_UPDATE, h.Name),
	})
}

func (h *Handler) Delete(ctx *gin.Context) {
	id, err := validator.PathInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	if err := h.Usecase.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, presenter.Default{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, presenter.Default{
		Message: fmt.Sprintf("%s %s", presenter.RESPONSE_SUCCESS_DELETE, h.Name),
	})
}
