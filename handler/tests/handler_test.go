package tests

import (
	"context"
	"database/sql"

	"github.com/r0busta/shop-data-replication/handler"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage/database"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

var _ = Describe("On product create", func() {
	It("inserts the product in the database", func() {
		product := &models.Product{
			ID:    9876543210,
			Title: null.StringFrom("Shopify Experts Coffee Mug"),
		}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())
		db := database.New(dbConn)
		Expect(err).To(BeNil())

		h := handler.NewDefaultHandler(db)
		err = h.OnProductCreate(product)
		Expect(err).To(BeNil())

		actual, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(1))

		Expect(actual[0].ID).To(Equal(int64(9876543210)))
		Expect(actual[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))
	})
})
