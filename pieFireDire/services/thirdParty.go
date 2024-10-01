package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func callGetMeat() (string, error) {
	timeOut := 2 * time.Minute
	client := http.Client{
		Timeout: timeOut,
	}
	var req *http.Request
	var err error
	req, err = http.NewRequest(http.MethodGet, "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text", nil)
	if err != nil {
		log.Error("req baconipsum error ", err)
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("call baconipsum error ", err)
		return "", err
	}
	defer resp.Body.Close()
	var rawBody []byte
	rawBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read response body error ", err)
		return "", err
	}
	if resp.StatusCode != 200 {
		log.Error("status code is ", resp.StatusCode, " message ", string(rawBody))
		return "", errors.New(string(rawBody))
	}
	return string(rawBody), nil
}
