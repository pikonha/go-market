package product_domain

type Repository interface {
	Delete(id int) error
	Store(*Product) error
	Update(*Product) error
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
}

