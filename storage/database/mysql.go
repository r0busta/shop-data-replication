package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/r0busta/shop-data-replication/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Database struct {
	connection *sql.DB
}

func New(connection *sql.DB) *Database {
	return &Database{
		connection: connection,
	}
}

func (s *Database) SaveProducts(products []*models.Product) error {
	for _, p := range products {
		err := s.SaveProduct(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Database) SaveProduct(product *models.Product) error {
	err := product.Upsert(context.Background(), s.connection, boil.Infer(), boil.Infer())
	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	return nil
}
