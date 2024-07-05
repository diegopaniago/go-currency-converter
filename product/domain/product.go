package currency

type Currency struct {
	alias string
}

type Product struct {
	id         int
	name       string
	price      float64
	currencies []Currency
}
