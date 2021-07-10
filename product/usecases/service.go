package product_usecase

import (
  "github.com/picolloo/go-market/product/domain"
)

type Service struct {
  repo product_domain.Repository
}

func NewService(repo product_domain.Repository) Service {
  return Service{
    repo: repo,
  }
}

func (s *Service) Delete(id int) error {
  err := s.repo.Delete(id)
  if err != nil {
    return err
  }
  return nil
}

func (s *Service) Store(p *product_domain.Product) error {
  err := s.repo.Store(p)
  if err != nil {
    return err
  }
  return nil
}

func (s *Service) Update(p *product_domain.Product) error {
  err := s.repo.Update(p)
  if err != nil {
    return err
  }
  return nil
}

func (s *Service) GetAll() ([]*product_domain.Product, error) {
  products, err := s.repo.GetAll()
  if err != nil {
    return nil, err
  }
  return products, nil
}

func (s *Service) Get(id int) (*product_domain.Product, error) {
  product, err := s.repo.Get(id)
  if err != nil {
    return nil, err
  }
  return product, nil
}
