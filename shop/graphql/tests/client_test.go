package tests

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	shopify "github.com/r0busta/go-shopify-graphql-model/graph/model"
	"gopkg.in/guregu/null.v4"
)

var _ = Describe("Products", func() {
	It("lists all products", func() {
		result := []*shopify.Product{{
			LegacyResourceID: null.StringFrom("9876543210"),
			Title:            null.StringFrom("Shopify Experts Coffee Mug"),
		}, {
			LegacyResourceID: null.StringFrom("9223372036854775807"),
			Title:            null.StringFrom("Shopify Experts Hoodie"),
		}}
		mockBulkOperationService.
			EXPECT().
			BulkQuery(gomock.Any(), gomock.Any()).
			SetArg(1, result)

		products, err := client.ListAllProducts()
		Expect(err).To(BeNil())
		Expect(len(products)).To(Equal(2))

		Expect(products[0].ID).To(Equal(int64(9876543210)))
		Expect(products[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))

		Expect(products[1].ID).To(Equal(int64(9223372036854775807)))
		Expect(products[1].Title.String).To(Equal("Shopify Experts Hoodie"))
	})
})
