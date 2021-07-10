package product_infra

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/picolloo/go-market/product/domain"
)


type PostgresProductRepository struct {
  db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
  return &PostgresProductRepository{
    db: db,
  }
}

func (self *PostgresProductRepository) GetAll() ([]*product_domain.Product, error) {
  var products []*product_domain.Product

  rows, err := self.db.Query("select id, name, price, type from products")
  if err != nil {
    return nil, err
  }

  defer rows.Close()
  for rows.Next() {
    var p product_domain.Product
    err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Type)
    if err != nil {
      return nil, err
    }
    products = append(products, &p)
  }

  return products, nil
}

func (self *PostgresProductRepository) Get(id int) (*product_domain.Product, error) {
  var product product_domain.Product

  stmt, err := self.db.Prepare("select id, name, price, type from products where id = $1")
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

func (self *PostgresProductRepository) Delete(id int) error {
  tx, err := self.db.Begin()
  if err != nil {
    return err
  }

  stmt, err := tx.Prepare("delete from products where id = $1")
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

func (self *PostgresProductRepository) Store(product *product_domain.Product) error {
  tx, err := self.db.Begin()
  if err != nil {
    return err
  }
  stmt, err := tx.Prepare("insert into products(name, price,type) values ($1,$2,$3)")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(&product.Name, &product.Price, &product.Type)

  if err != nil {
    tx.Rollback()
    return err
  }

  tx.Commit()
  return nil
}

func (self *PostgresProductRepository) Update(product *product_domain.Product) error {
  tx, err := self.db.Begin()
  if err != nil {
    return err
  }
  stmt, err := tx.Prepare("update products set name=$1, price=$2, type=$3 where id=$4")
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
