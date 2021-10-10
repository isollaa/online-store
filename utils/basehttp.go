package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HTTPRequest(request *http.Request, timeout string) (error, []byte, int) {
	duration, _ := time.ParseDuration(timeout)

	client := &http.Client{
		Timeout: duration,
	}
	testResp, err := client.Do(request)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed connect service %v", err))
		return err, []byte(``), testResp.StatusCode
	}
	defer testResp.Body.Close()

	body, errRead := ioutil.ReadAll(testResp.Body)

	if errRead != nil {
		return err, []byte(``), testResp.StatusCode
	}

	return nil, body, testResp.StatusCode
}
