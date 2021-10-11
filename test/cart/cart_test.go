package cart

//this package is used to test data race & make sure requester wont be able to create cart when selected item is empty

import (
	"online-store/api/cart/entity"
	"online-store/test"
	"sync"
	"testing"
)

//customize token if expired by relogin response token
var token = test.TOKEN_ONE_MONTH_EXP

//make sure item is exist if not, do test create item first
func TestCreateCart(t *testing.T) {
	url := "http://localhost:3000/cart"
	cart := entity.Cart{
		ItemID:   1,
		Quantity: 1,
	}

	var wg sync.WaitGroup

	//concurrent create cart on the same item & time
	for i := 1; i <= 150; i++ {
		wg.Add(1)
		go test.RequestPostConcurrent(cart, url, token, &wg)
	}

	wg.Wait()
}
