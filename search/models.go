package search

type ResponseSearch struct {
	Products []Product `json:"products"`
	Facets   struct {
		MinPrice int `json:"min_price"`
		MaxPrice int `json:"max_price"`
		Brands   []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"brands"`
		Size []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"size"`
	} `json:"facets"`
	Total int `json:"total"`
}

type Product struct {
	Name         string `json:"name"`
	Date         string `json:"date"`
	Image        string `json:"image"`
	URL          string `json:"url"`
	Price        string `json:"price"`
	SpecialPrice string `json:"special_price"`
	//InStoreOnly  bool        `json:"in_store_only"`
	//Manufacturer string      `json:"manufacturer"`
	//NewsToDate   string      `json:"news_to_date"`
	//Sku          string      `json:"sku"`
	//Status       int         `json:"status"`
	//Stock        int         `json:"stock"`
	//SellState    interface{} `json:"sell_state"`
}
