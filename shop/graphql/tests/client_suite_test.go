package tests

import (
	"testing"

	"github.com/r0busta/shop-data-replication/shop"
	graphqlShop "github.com/r0busta/shop-data-replication/shop/graphql"
	"github.com/r0busta/shop-data-replication/shop/graphql/tests/mock_v3"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	shopifyGraphql "github.com/r0busta/go-shopify-graphql/v3"
)

const (
	storeApi = "https://shop.myshopify.com/admin/api/2021-04/graphql.json"
)

var client shop.Client
var mockBulkOperationService *mock_v3.MockBulkOperationService

func TestGraphqlClient(t *testing.T) {
	client = newMockGraphqlClient(t)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphql Client Suite")
}

func newMockGraphqlClient(t *testing.T) shop.Client {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBulkOperationService = mock_v3.NewMockBulkOperationService(ctrl)

	shopifyClient := &shopifyGraphql.Client{
		BulkOperation: mockBulkOperationService,
	}
	return graphqlShop.New(shopifyClient)
}
