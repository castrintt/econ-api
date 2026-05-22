package products

import "context"

// contract
type Service interface {
	GetProducts(context context.Context) ([]any,error)
}

// struct
type service struct {
}

// constructor
func NewService() *service {
	return &service{}
}

// methods
func (s *service) GetProducts(context context.Context) ([]any,error) {
	return nil, nil
}
