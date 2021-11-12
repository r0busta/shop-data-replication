package shop

import "github.com/r0busta/shop-data-replication/models"

// Client for interactions with the Shopify store.
type Client interface {
	ListAllProducts() ([]*models.Product, error)
}
