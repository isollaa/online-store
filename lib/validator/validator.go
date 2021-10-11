package validator

import (
	"fmt"
	"online-store/lib/conv"
	"online-store/lib/presenter"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PathInt64(ctx *gin.Context, key string) (res int64, err error) {
	res = conv.StringToInt64(ctx.Param(key), 0)
	if res == 0 {
		err = fmt.Errorf("'%s' %s", key, presenter.RESPONSE_ERROR_INVALID_PARAM_NUMBER_NOT_EMPTY)
		return
	}

	return
}

func SetDefaultQueryParamNumber(ctx *gin.Context, key string, dval string) (int, error) {
	p := ctx.Query(key)
	if p == "" {
		p = dval
	}

	param, err := strconv.Atoi(p)
	if err != nil {
		return 0, fmt.Errorf("'%s' %s", key, presenter.RESPONSE_ERROR_INVALID_PARAM_NUMBER)
	}

	return param, nil
}
