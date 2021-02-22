package tests

import (
	"context"
	"database/sql"

	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage/database"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

var _ = Describe("Save products", func() {
	It("saves all products in the database", func() {
		products := []*models.Product{{
			ID:    9876543210,
			Title: null.StringFrom("Shopify Experts Coffee Mug"),
		}, {
			ID:    9223372036854775807,
			Title: null.StringFrom("Shopify Experts Hoodie"),
		}}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())

		db := database.New(dbConn)
		err = db.SaveProducts(products)
		Expect(err).To(BeNil())

		actual, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(2))

		Expect(actual[0].ID).To(Equal(int64(9876543210)))
		Expect(actual[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))

		Expect(actual[1].ID).To(Equal(int64(9223372036854775807)))
		Expect(actual[1].Title.String).To(Equal("Shopify Experts Hoodie"))
	})
})

var _ = Describe("Save product", func() {
	It("inserts the new product into the database", func() {
		product := &models.Product{
			ID:    9876543210,
			Title: null.StringFrom("Shopify Experts Coffee Mug"),
		}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())

		db := database.New(dbConn)

		err = db.SaveProduct(product)
		Expect(err).To(BeNil())

		actual, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(1))

		Expect(actual[0].ID).To(Equal(int64(9876543210)))
		Expect(actual[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))
	})

	It("updates the existing product in the database", func() {
		product := &models.Product{
			ID:    9876543210,
			Title: null.StringFrom("Shopify Experts Coffee Mug"),
		}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())

		db := database.New(dbConn)

		err = db.SaveProduct(product)
		Expect(err).To(BeNil())

		actual, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(1))

		// Update
		product.Title = null.StringFrom("Shopify Experts Coffee Mug Large")
		product.Vendor = null.StringFrom("Blackbird Commerce")
		err = db.SaveProduct(product)
		Expect(err).To(BeNil())

		actual, err = models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(1))

		Expect(actual[0].ID).To(Equal(int64(9876543210)))
		Expect(actual[0].Title.String).To(Equal("Shopify Experts Coffee Mug Large"))
	})
})
