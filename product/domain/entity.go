package product_domain

type Product struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Price float32     `json:"price"`
	Type  ProductType `json:"type"`
}

type ProductType int

const (
	Food = iota + 1
	Toy
	Electronic
)

func (p ProductType) String() string {
	switch p {
	case Food:
		return "Food"
	case Toy:
		return "Toy"
	case Electronic:
		return "Electronic"
	default:
		return "Unknown"
	}
}
