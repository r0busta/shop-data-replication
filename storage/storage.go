package storage

import (
	"github.com/r0busta/shop-data-replication/models"
)

type Storage interface {
	SaveProducts(products []*models.Product) error
	SaveProduct(product *models.Product) error

	// SaveProductVariants(variants []*models.ProductVariant) error
	// SaveProductVariant(variant *models.ProductVariant) error

	// SaveProductOptions(options []*models.ProductOption, values []*models.ProductOptionValue) error
	// SaveProductOption(option *models.ProductOption, values []*models.ProductOptionValue) error

	// SaveProductImages(productID int64, images []*models.Image) error
	// SaveProductImage(productID int64, image models.Image) error

	// SaveCollections(collections []*models.Collection) error
	// SaveCollection(collection *models.Collection) error

	// SaveCollectionProducts(collectionID int64, productIDs []int64) error
	// RemoveCollectionProducts(collectionID int64, productIDs []int64) error

	SaveOrders(orders []*models.Order) error
	SaveOrder(order *models.Order) error
}
