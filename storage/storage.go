package storage

import (
	"github.com/r0busta/shop-data-replication/models"
)

type Storage interface {
	SaveProducts(products []*models.Product) error
	SaveProduct(product *models.Product) error
}
