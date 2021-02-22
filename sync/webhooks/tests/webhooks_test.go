package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/r0busta/shop-data-replication/handler"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage/database"
	"github.com/r0busta/shop-data-replication/sync/webhooks"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	goshopify "github.com/r0busta/go-shopify/v3"
	"github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
)

const apiSecret = "secret-key-42"

var _ = Describe("POST /on/products/create", func() {
	It("saves the new product", func() {
		product := goshopify.Product{
			ID:    9876543210,
			Title: "Shopify Experts Coffee Mug",
		}

		dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
		Expect(err).To(BeNil())
		db := database.New(dbConn)

		h := handler.NewDefaultHandler(db)
		s := webhooks.New(h, webhooks.WithVerifyRequests(false))
		router := s.(*webhooks.Service).SetupRouter()

		body, err := json.Marshal(product)
		Expect(err).To(BeNil())
		req, err := http.NewRequest("POST", "/api/v1/on/products/create", bytes.NewReader(body))
		Expect(err).To(BeNil())

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(200))

		actual, err := models.Products().All(context.Background(), dbConn)
		Expect(err).To(BeNil())
		Expect(len(actual)).To(Equal(1))

		Expect(actual[0].ID).To(Equal(int64(9876543210)))
		Expect(actual[0].Title.String).To(Equal("Shopify Experts Coffee Mug"))
	})
})

var _ = Describe("Verify Webhooks", func() {
	It("doesn't verify requests with a nil body", func() {
		req, err := http.NewRequest("POST", "", nil)
		Expect(err).To(BeNil())

		err = webhooks.VerifyRequest(req, apiSecret)
		Expect(err).To(Not(BeNil()))
	})

	It("verifies the request", func() {
		req, err := http.NewRequest("POST", "", bytes.NewBufferString(""))
		Expect(err).To(BeNil())

		req.Header.Add("X-Shopify-Hmac-Sha256", "Ni1IFKKXDbi6xXEBR/hBjKgWbI0/trwMcufJFOT0aLw=")

		err = webhooks.VerifyRequest(req, apiSecret)
		Expect(err).To(BeNil())
	})
})
