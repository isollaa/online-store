package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func RequestPost(payload interface{}, url string, token string) (r string) {
	method := "POST"

	b, _ := json.Marshal(payload)
	p := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, p)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	r = string(body)
	return
}

func RequestPostConcurrent(payload interface{}, url string, token string, wg *sync.WaitGroup) {
	defer wg.Done()
	method := "POST"

	b, _ := json.Marshal(payload)
	p := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, p)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	fmt.Println(body)
}
