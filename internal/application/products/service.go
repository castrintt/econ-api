package products

import "context"

// contract
type Service interface {
	GetProducts(context context.Context) error
}

// contract
type Repository interface {
}

// struct
type service struct {
	repository Repository
}

// constructor
func NewService(repository Repository) *service {
	return &service{repository: repository}
}
