package search

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var EURtoRUB float64 = 65

func requestDo(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("status code: %d - %s", res.StatusCode, res.Status)
	}
	return ioutil.ReadAll(res.Body)
}

// encode64 receive message and return base64 encoded string
func encode64(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

// converts int to string
func intToString(i int) string {
	return strconv.Itoa(i)
}

func stringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func toRUB(i float64) float64 {
	return i * EURtoRUB
}