package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	urlBase   = "https://www.thegoodwillout.com/"
	urlSearch = urlBase + "/search/full.php"
)

const (
	defaultLanguage    = "en"
	defaultNumProducts = 99
)

// Search receive query string
// returns csv with results, total count and error
func Search(query string, page int) (string, int, error) {
	products, err := search(query, page)
	if err != nil {
		return "", 0, err
	}
	csvString := processProduct(products.Products...)
	return csvString, products.Total, nil
}

// search receive query string
// returns ResponseSearch
func search(query string, page int) (ResponseSearch, error) {
	var resp ResponseSearch
	qs := buildQueryString(query, page)
	req := buildRequest(qs)
	b, err := requestDo(req)
	if err != nil {
		return resp, err
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}

// buildQueryString receive search query and page
// return base64 encoded query string
func buildQueryString(q string, pageNum int) string {
	b := new(bytes.Buffer)
	b.WriteString("q=" + encode64(q) + "&")
	b.WriteString("index=" + encode64(defaultLanguage) + "&")
	b.WriteString("i=" + encode64(intToString(defaultNumProducts)) + "&")
	b.WriteString("p=" + encode64(intToString(pageNum)))
	return encode64(b.String())
}

// buildRequest return request
func buildRequest(queryString string) *http.Request {
	req, _ := http.NewRequest("GET", urlSearch, nil)
	q := req.URL.Query()
	q.Add("qs", queryString)
	req.URL.RawQuery = q.Encode()
	return req
}

// processProduct receive Product slice,
// return this info in csv format: name, url, image, price, specialPrice(1 if it is)
func processProduct(products ...Product) string {
	b := new(bytes.Buffer)
	for _, product := range products {
		b.WriteString(product.Name + ",")
		b.WriteString(urlBase + product.URL + ",")
		b.WriteString(product.Image + ",")
		if product.SpecialPrice != "" {
			b.WriteString(processPrice(product.SpecialPrice) + ",1")
		} else {
			b.WriteString(processPrice(product.Price))
		}
		b.WriteString("\n")
	}
	return b.String()
}

// processPrice receive price string, change currency and return string
func processPrice(priceString string) string {
	price, err := stringToFloat(priceString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// multiple price to currency rate
		price = toRUB(price)
	}
	return fmt.Sprintf("%f", price)
}
