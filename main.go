package main

import (
	"fmt"

	"github.com/vlad2095/thegoodwillout/search"
)

// example usage
func main() {
	text := "Black t-shirt"
	page := 1
	res, count, err := search.Search(text, page)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
		fmt.Println("Total count:", count)
	}
}
