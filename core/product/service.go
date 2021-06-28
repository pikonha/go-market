package product

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UseCase interface {
	Delete(id int) error
	Store(*Product) error
	Update(*Product) error
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
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
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Store(product *Product) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into products(name, price,type) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&product.Name, &product.Price, &product.Type)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (s *Service) Update(product *Product) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update products set name=?, price=?, type=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&product.Name, &product.Price, &product.Type, &product.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
