package product_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/picolloo/go-market/core/product"
)

func clearDb(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from products")
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func TestStore(t *testing.T) {
	p := &product.Product{
		ID:    1,
		Name:  "product",
		Price: 10.5,
		Type:  product.ProductType(product.Food),
	}
	db, err := sql.Open("sqlite3", "../../data/product.db")
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados %s", err.Error())
	}
	defer db.Close()
	err = clearDb(db)
	if err != nil {
		t.Fatalf("Erro ao limpar banco de dados %s", err.Error())
	}
	service := product.NewService(db)
	err = service.Store(p)
	if err != nil {
		t.Fatalf("Erro ao inserir produto %s", err.Error())
	}
	saved, err := service.Get(p.ID)
	if err != nil {
		t.Fatalf("Erro ao buscar produto %s", err.Error())
	}
	if saved.ID != p.ID {
		t.Errorf("Dados inválidos. Esperando ID: %d, recebido: %d", p.ID, saved.ID)
	}
	if saved.Name != p.Name {
		t.Errorf("Dados inválidos. Esperando Name: %s, recebido: %s", p.Name, saved.Name)
	}
	if saved.Price != p.Price {
		t.Errorf("Dados inválidos. Esperando Price: %f, recebido: %f", p.Price, saved.Price)
	}
	if saved.Type != p.Type {
		t.Errorf("Dados inválidos. Esperando Type: %s, recebido: %s", p.Type.String(), saved.Type.String())
	}
}

func TestUpdate(t *testing.T) {
	p := &product.Product{
		ID:    1,
		Name:  "product",
		Price: 10.5,
		Type:  product.ProductType(product.Food),
	}
	db, err := sql.Open("sqlite3", "../../data/product.db")
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados %s", err.Error())
	}
	defer db.Close()
	err = clearDb(db)
	if err != nil {
		t.Fatalf("Erro ao limpar banco de dados %s", err.Error())
	}
	service := product.NewService(db)
	err = service.Store(p)
	if err != nil {
		t.Fatalf("Erro ao inserir produto %s", err.Error())
	}
	updated := &product.Product{
		ID:    1,
		Name:  "p2",
		Price: 20,
		Type:  product.ProductType(product.Toy),
	}
	err = service.Update(updated)
	if err != nil {
		t.Fatalf("Erro ao atualizar produto %s", err.Error())
	}
	saved, err := service.Get(p.ID)
	if err != nil {
		t.Fatalf("Erro ao buscar produto %s", err.Error())
	}
	if saved.ID != updated.ID {
		t.Errorf("Dados inválidos. Esperando ID: %d, recebido: %d", updated.ID, saved.ID)
	}
	if saved.Name != updated.Name {
		t.Errorf("Dados inválidos. Esperando Name: %s, recebido: %s", updated.Name, saved.Name)
	}
	if saved.Price != updated.Price {
		t.Errorf("Dados inválidos. Esperando Price: %f, recebido: %f", updated.Price, saved.Price)
	}
	if saved.Type != updated.Type {
		t.Errorf("Dados inválidos. Esperando Type: %s, recebido: %s", updated.Type.String(), saved.Type.String())
	}
}

func TestDelete(t *testing.T) {
	p := &product.Product{
		ID:    1,
		Name:  "product",
		Price: 50.4,
		Type:  product.ProductType(product.Electronic),
	}
	db, err := sql.Open("sqlite3", "../../data/product.db")
	clearDb(db)
	if err != nil {
		t.Fatalf("Erro ao limpar banco de dados %s", err.Error())
	}
	service := product.NewService(db)
	err = service.Store(p)
	if err != nil {
		t.Fatalf("Erro ao inserir produto %s", err.Error())
	}
	err = service.Delete(p.ID)
	if err != nil {
		t.Fatalf("Erro ao deletar produto %s", err.Error())
	}
	saved, _ := service.Get(p.ID)
	if saved != nil {
		t.Fatalf("Erro ao deletar produto %d", p.ID)
	}
}

func TestGet(t *testing.T) {
	p := &product.Product{
		ID:    1,
		Name:  "product",
		Price: 50.4,
		Type:  product.ProductType(product.Electronic),
	}
	db, err := sql.Open("sqlite3", "../../data/product.db")
	clearDb(db)
	if err != nil {
		t.Fatalf("Erro ao limpar banco de dados %s", err.Error())
	}
	service := product.NewService(db)
	err = service.Store(p)
	if err != nil {
		t.Fatalf("Erro ao inserir produto %s", err.Error())
	}
	_, err = service.Get(p.ID)
	if err != nil {
		t.Fatalf("Erro ao buscar produto %s", err.Error())
	}
}

func TestGetAll(t *testing.T) {
	ps := []*product.Product{
		{
			ID:    1,
			Name:  "product",
			Price: 50.4,
			Type:  product.ProductType(product.Electronic),
		},
		{
			ID:    2,
			Name:  "product",
			Price: 100.2,
			Type:  product.ProductType(product.Electronic),
		},
	}
	db, err := sql.Open("sqlite3", "../../data/product.db")
	clearDb(db)
	if err != nil {
		t.Fatalf("Erro ao limpar banco de dados %s", err.Error())
	}
	service := product.NewService(db)

	for _, p := range ps {
		err = service.Store(p)
		if err != nil {
			t.Fatalf("Erro ao inserir produto %s", err.Error())
		}
	}

	products, err := service.GetAll()
	if err != nil {
		t.Fatalf("Erro ao buscar todos os produtos %s", err.Error())
	}
	if len(products) != len(ps) {
		t.Fatalf("Numero de produtos encontrados inválido. Esperado %d, recebido %d", len(ps), len(products))
	}
}
