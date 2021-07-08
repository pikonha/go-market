package product_use_case

import (
  "github.com/picolloo/go-market/product/domain"
)

type Service struct {
  repo *product_domain.Repository
}

func NewService(repo *product_domain.Repository) *Service {
  return &Service{
    repo: repo,
  }
}

func (self *Service) Delete(id int) error {
  err := self.repo.Delete(id)
  if err != nil {
    return err
  }
  return nil
}

func (self *Service) Store(p *product_domain.Product) error {
  err := self.repo.Store(p)
  if err != nil {
    return err
  }
  return nil
}

func (self *Service) Update(p *product_domain.Product) error {
  err := self.repo.Update(p)
  if err != nil {
    return err
  }
  return nil
}

func (self *Service) GetAll() ([]*product_domain.Product, error) {
  products, err := self.repo.GetAll()
  if err != nil {
    return nil, err
  }
  return products, nil
}

func (self *Service) Get(id int) (*product_domain.Product, error) {
  product, err := self.repo.Get(id)
  if err != nil {
    return nil, err
  }
  return product, nil
}
