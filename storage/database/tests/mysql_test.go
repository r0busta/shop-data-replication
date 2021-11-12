package tests_test

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage/database"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var _ = Describe("Save products", func() {
	It("saves all products in the database", func() {
		products := []*models.Product{{
			ID:             9876543210,
			Title:          null.StringFrom("Shopify Experts Coffee Mug"),
			BodyHTML:       null.StringFrom("<p>Lorem ipsum dolor sit amet</p>"),
			Handle:         "shopify-experts-coffee-mug",
			ProductType:    null.StringFrom("Mug"),
			Vendor:         null.StringFrom("Blackbird"),
			Status:         1, // active
			PublishedScope: "web",
			TemplateSuffix: null.StringFrom(""),
		}, {
			ID:             9223372036854775807,
			Title:          null.StringFrom("Shopify Experts Hoodie"),
			BodyHTML:       null.StringFrom("<p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.</p><p>laudantium, totam rem aperiam, eaque ipsa quae ab.</p>"),
			Handle:         "shopify-experts-hoodie",
			ProductType:    null.StringFrom("Hoodie"),
			Vendor:         null.StringFrom("Blackbird"),
			Status:         3, // draft
			PublishedScope: "global",
			TemplateSuffix: null.StringFrom("custom"),
		}}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())

		db := database.New(dbConn)
		err = db.SaveProducts(products)
		Expect(err).To(BeNil())

		got, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(got)).To(Equal(2))

		Expect(got[0].ID).To(Equal(int64(9876543210)))
		Expect(got[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))
		Expect(got[0].BodyHTML.String).To(Equal("<p>Lorem ipsum dolor sit amet</p>"))
		Expect(got[0].Handle).To(Equal("shopify-experts-coffee-mug"))
		Expect(got[0].ProductType.String).To(Equal("Mug"))
		Expect(got[0].Vendor.String).To(Equal("Blackbird"))
		Expect(got[0].Status).To(Equal(int8(1)))
		Expect(got[0].PublishedScope).To(Equal("web"))
		Expect(got[0].TemplateSuffix.String).To(Equal(""))

		Expect(got[1].ID).To(Equal(int64(9223372036854775807)))
		Expect(got[1].Title.String).To(Equal("Shopify Experts Hoodie"))
		Expect(got[1].BodyHTML.String).To(Equal("<p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque.</p><p>laudantium, totam rem aperiam, eaque ipsa quae ab.</p>"))
		Expect(got[1].Handle).To(Equal("shopify-experts-hoodie"))
		Expect(got[1].ProductType.String).To(Equal("Hoodie"))
		Expect(got[1].Vendor.String).To(Equal("Blackbird"))
		Expect(got[1].Status).To(Equal(int8(3)))
		Expect(got[1].PublishedScope).To(Equal("global"))
		Expect(got[1].TemplateSuffix.String).To(Equal("custom"))
	})
})

var _ = Describe("Save product", func() {
	It("inserts a new product into the database", func() {
		product := &models.Product{
			ID:    9876543210,
			Title: null.StringFrom("Shopify Experts Coffee Mug"),
		}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())

		db := database.New(dbConn)

		err = db.SaveProduct(product)
		Expect(err).To(BeNil())

		want := models.Product{}
		err = queries.Raw(`select * from product`).Bind(context.Background(), dbConn, &want)

		Expect(err).To(BeNil())

		Expect(want.ID).To(Equal(int64(9876543210)))
		Expect(want.Title.String).To(Equal("Shopify Experts Coffee Mug"))
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

		// Update
		product.Title = null.StringFrom("Shopify Experts Coffee Mug Large")
		product.Vendor = null.StringFrom("Blackbird Commerce")
		err = db.SaveProduct(product)
		Expect(err).To(BeNil())

		want := models.Product{}
		err = queries.Raw(`select * from product`).Bind(context.Background(), dbConn, &want)

		Expect(err).To(BeNil())

		Expect(want.ID).To(Equal(int64(9876543210)))
		Expect(want.Title.String).To(Equal("Shopify Experts Coffee Mug Large"))
	})
})
