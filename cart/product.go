package cart

type Product struct {
	Id string
	Name string
	Price float64
}

type CProduct struct {
	Product
	Quantity int
}

