package tests_test

//go:generate mockgen -destination mock_v3/bulk_operation_service.go github.com/r0busta/go-shopify-graphql/v3 BulkOperationService
import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	shopifyGraphql "github.com/r0busta/go-shopify-graphql/v3"
	"github.com/r0busta/shop-data-replication/shop"
	graphqlShop "github.com/r0busta/shop-data-replication/shop/graphql"
	"github.com/r0busta/shop-data-replication/shop/graphql/tests/mock_v3"
)

var (
	client                   shop.Client
	mockBulkOperationService *mock_v3.MockBulkOperationService
)

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
