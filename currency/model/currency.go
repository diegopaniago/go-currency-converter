package model

type Exchange struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Currency struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Exchange Exchange `json:"exchange"`
}
