package cart

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"online-store/lib/conv"
	"sync"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoxNjMzODc4Mjg0LCJ1c2VyX2lkIjoxfQ.FblVj3d8lUsoDf3TUqq7qyOmL2XlpEpCvYTU3LC0A6E"

func Request(qty float64, wg *sync.WaitGroup) {
	defer wg.Done()
	url := "http://localhost:3000/cart"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("item_id", "1")
	_ = writer.WriteField("quantity", conv.Float64ToString(qty))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoxNjMzODg2NDE4LCJ1c2VyX2lkIjoxfQ.cHl8iKv6Mj5iInld9y7_5n3Lk8pVSqz49_bbOE1C2l0")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
