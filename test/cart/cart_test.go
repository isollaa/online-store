package cart

import (
	"encoding/json"
	"online-store/api/cart/entity"
	"sync"
	"testing"
)

func payloadCreate(qty float64) []byte {
	c := entity.Cart{
		UserID:   1,
		ItemID:   1,
		Quantity: qty,
	}

	payload, _ := json.Marshal(&c)
	return payload
}

func TestCreateCart(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go Request((float64(1)), &wg)
	}

	wg.Wait()
}
