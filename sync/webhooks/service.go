package webhooks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	goshopify "github.com/r0busta/go-shopify/v3"
	"github.com/r0busta/graphql"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/sync/handler"
	log "github.com/sirupsen/logrus"
	sqlNull "github.com/volatiletech/null/v8"
)

const routeGroupPrefix = "/api/v1"

type Topic int

const (
	TopicProductsCreate Topic = iota
	TopicProductsUpdate
	TopicProductsDelete
)

var topicHandlerPaths = map[Topic]string{
	TopicProductsCreate: "/on/products/create",
	TopicProductsUpdate: "/on/products/update",
	TopicProductsDelete: "/on/products/delete",
}

// Shopify webhook handler. Syncs updates received via webhooks to the database.
type Service struct {
	handler handler.Handler

	port             string
	shopifyAPISecret string

	verifyRequests bool
}

func New(h handler.Handler, opts ...Option) *Service {
	s := &Service{
		handler:        h,
		verifyRequests: true,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type Option func(s *Service)

func WithPort(v string) Option {
	return func(s *Service) {
		s.port = v
	}
}

func WithShopifyAPISecret(v string) Option {
	return func(s *Service) {
		s.shopifyAPISecret = v
	}
}

func WithVerifyRequests(v bool) Option {
	return func(s *Service) {
		s.verifyRequests = v
	}
}

func (s *Service) SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group(routeGroupPrefix)

	if s.verifyRequests {
		api.Use(s.VerificationRequired)
	} else {
		log.Warnln("Danger Zone: unauthenticated requests are allowed")
	}

	{
		api.POST(topicHandlerPaths[TopicProductsCreate], s.onProductCreate)
		api.POST(topicHandlerPaths[TopicProductsUpdate], s.onProductUpdate)
		api.POST(topicHandlerPaths[TopicProductsDelete], s.onProductDelete)

		api.Any("/", s.notFound)
	}
	r.Any("/", s.notFound)

	return r
}

func (s *Service) ProvisionSubscriptions(shopClient *graphql.Client, callbackBaseURL string) error {
	sub := Subscriptions{
		shopClient:      shopClient,
		callbackBaseURL: callbackBaseURL,
	}

	return sub.ProvisionSubscriptions([]Topic{
		TopicProductsCreate,
		TopicProductsUpdate,
		TopicProductsDelete,
	})
}

func (s *Service) Run() error {
	r := s.SetupRouter()
	port := fmt.Sprintf(":%s", s.port)

	return r.Run(port)
}

func (s *Service) onProductCreate(c *gin.Context) {
	product := &goshopify.Product{}

	err := c.ShouldBindJSON(product)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	err = s.handler.OnProductCreate(convertProductRequest(product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *Service) onProductUpdate(c *gin.Context) {
	product := &goshopify.Product{}
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	err = s.handler.OnProductUpdate(convertProductRequest(product))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *Service) onProductDelete(c *gin.Context) {
	c.String(http.StatusInternalServerError, "Not implemented")
}

func (s *Service) notFound(c *gin.Context) {
	c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func convertProductRequest(product *goshopify.Product) *models.Product {
	return &models.Product{
		ID:    product.ID,
		Title: sqlNull.StringFrom(product.Title),
	}
}

func (s *Service) VerificationRequired(c *gin.Context) {
	if err := VerifyRequest(c.Request, s.shopifyAPISecret); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}
}

func VerifyRequest(req *http.Request, secret string) error {
	if req.Body == nil {
		return fmt.Errorf("received nil request body")
	}

	actual := []byte(req.Header.Get("X-Shopify-Hmac-Sha256"))
	mac := hmac.New(sha256.New, []byte(secret))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	mac.Write(body)
	expected := []byte(
		base64.StdEncoding.EncodeToString(mac.Sum(nil)),
	)

	if hmac.Equal(actual, expected) {
		return nil
	}

	return fmt.Errorf("signature not verified")
}
