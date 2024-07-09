package service

type ProductDto struct {
	Name       string   `json:"name"`
	Price      float32  `json:"price"`
	Currencies []string `json:"currencies"`
}
