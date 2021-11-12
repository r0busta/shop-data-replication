package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Database struct {
	connection *sql.DB
}

var _ storage.Storage = &Database{}

func New(connection *sql.DB) *Database {
	return &Database{
		connection: connection,
	}
}

func (d *Database) SaveProducts(products []*models.Product) error {
	for _, p := range products {
		err := d.SaveProduct(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Database) SaveProduct(product *models.Product) error {
	err := product.Upsert(context.Background(), d.connection, boil.Infer(), boil.Infer())
	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	return nil
}

func (d *Database) SaveProductVariants(variants []*models.ProductVariant) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveProductVariant(variant *models.ProductVariant) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveProductOptions(options []*models.ProductOption, values []*models.ProductOptionValue) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveProductOption(option *models.ProductOption, values []*models.ProductOptionValue) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveProductImages(productID int64, images []*models.Image) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveProductImage(productID int64, image models.Image) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveCollections(collections []*models.Collection) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveCollection(collection *models.Collection) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveCollectionProducts(collectionID int64, productIDs []int64) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) RemoveCollectionProducts(collectionID int64, productIDs []int64) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveOrders(orders []*models.CustomerOrder) error {
	return fmt.Errorf("not implemented")
}

func (d *Database) SaveOrder(order *models.CustomerOrder) error {
	return fmt.Errorf("not implemented")
}
