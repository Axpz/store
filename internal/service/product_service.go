package service

import (
	"fmt"

	"github.com/Axpz/store/internal/storage"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	store storage.StoreInterface
}

func NewProductService(store storage.StoreInterface) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (s *ProductService) CreateProduct(c *gin.Context, product *storage.Product) error {
	return fmt.Errorf("not implemented")
}

func (s *ProductService) GetProduct(c *gin.Context, id string) (*storage.Product, error) {
	product, err := s.store.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) GetProducts(c *gin.Context) ([]storage.Product, error) {
	products, err := s.store.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(c *gin.Context, product *storage.Product) error {
	return fmt.Errorf("not implemented")
}

func (s *ProductService) DeleteProduct(c *gin.Context, id string) error {
	return fmt.Errorf("not implemented")
}
