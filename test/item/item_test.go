package item

import (
	"online-store/api/item/entity"
	"online-store/test"
	"testing"
)

//customize token if expired by relogin response token
var token = test.TOKEN_ONE_MONTH_EXP

func TestCreateCart(t *testing.T) {
	url := "http://localhost:3000/item"
	item := entity.Item{
		Name:     "item_testing",
		Quantity: 1000,
	}

	test.RequestPost(item, url, token)
}
