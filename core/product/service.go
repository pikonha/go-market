package product

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Writer interface {
	Delete(id int) error
	Store(*Product) error
	Update(*Product) error
}

type Reader interface {
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
}

type UseCase interface {
	Writer
	Reader
}

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAll() ([]*Product, error) {
	var products []*Product

	rows, err := s.db.Query("select id, name, price, type from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Type)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

func (s *Service) Get(id int) (*Product, error) {
	var product Product

	stmt, err := s.db.Prepare("select id, name, price, type from products where id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Type)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Service) Delete(id int) error {
	return nil
}

func (s *Service) Store(product *Product) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into products(id, name, price,type) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&product.ID, &product.Name, &product.Price, &product.Type)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(*Product) error {
	return nil
}
