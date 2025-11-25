package api

import (
	"database/sql"
	"goinventory/dataservice"
	"goinventory/model"

	"github.com/IBM/sarama"
)

// Bizlogic Interface and Implementation
type IBizlogic interface {
	AddProduct(product model.Product) (model.Product, error)
	GetProduct(id int) (model.Product, error)
	ListProducts(search string) ([]model.Product, error)
}

// Bizlogic Implementation
type Bizlogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

// NewBizlogic creates a new instance of Bizlogic
func NewBizlogic(db *sql.DB, producer sarama.SyncProducer) *Bizlogic {
	return &Bizlogic{
		DB:       db,
		Producer: producer,
	}
}

// AddProduct adds a new product to the inventory
func (biz *Bizlogic) AddProduct(product model.Product) (model.Product, error) {
	newID, err := dataservice.AddProduct(biz.DB, product)
	if err != nil {
		return model.Product{}, err
	}

	createdProduct, err := dataservice.GetProduct(biz.DB, newID)
	if err != nil {
		return model.Product{}, err
	}

	return createdProduct, nil
}

// GetProduct retrieves a product by its ID
func (biz *Bizlogic) GetProduct(id int) (model.Product, error) {
	return dataservice.GetProduct(biz.DB, id)
}

// ListProducts lists products based on a search query
func (biz *Bizlogic) ListProducts(search string) ([]model.Product, error) {
	return dataservice.ListProducts(biz.DB, search)
}
