package binder

import (
	"online-store/lib/request"
	"online-store/lib/validator"

	"github.com/gin-gonic/gin"
)

func ValidateRequestList(ctx *gin.Context) (request.List, error) {
	page, err := validator.SetDefaultQueryParamNumber(ctx, "page", "1")
	if err != nil {
		return request.List{}, err
	}
	perPage, err := validator.SetDefaultQueryParamNumber(ctx, "per_page", "10")
	if err != nil {
		return request.List{}, err
	}

	return request.List{
		Search:  ctx.Query("search"),
		Sort:    ctx.Query("sort"),
		Order:   ctx.Query("order"),
		PerPage: perPage,
		Page:    page,
	}, nil
}
