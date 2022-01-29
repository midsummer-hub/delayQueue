package httpClient

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// SendPostRequest TODO
func SendPostRequest(url, data string) (err error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("post err:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if strings.ToLower(string(body)) != "SUCCESS" {
		err = errors.New("post resule is not success ,try to rePost")
		log.Println("post resule is not success err:", string(body))
	}
	return
}
