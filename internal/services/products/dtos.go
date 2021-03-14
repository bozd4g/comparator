package products

type Dto struct {
	Name    string     `json:"name"`
	Product ProductDto `json:"product"`
}

type MultipleDto struct {
	Name     string       `json:"name"`
	Total    float64      `json:"total"`
	Products []ProductDto `json:"products"`
}

type ProductDto struct {
	Link  string  `json:"link"`
	Price float64 `json:"price"`
}
