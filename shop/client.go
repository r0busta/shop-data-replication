package shop

import "github.com/r0busta/shop-data-replication/models"

type Client interface {
	ListAllProducts() ([]*models.Product, error)
}
