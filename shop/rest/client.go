package restShop

import (
	"log"
	"time"

	rest "github.com/r0busta/go-shopify/v3"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/shop"
	"github.com/volatiletech/null/v8"
)

const maxRemoteListLimit = 250

type Client struct {
	client *rest.Client
}

func New(restClient *rest.Client) shop.Client {
	return &Client{
		client: restClient,
	}
}

func (c *Client) ListAllProducts() ([]*models.Product, error) {
	opts := rest.ListOptions{Limit: maxRemoteListLimit}
	products, pagination, err := c.client.Product.ListWithPagination(opts)
	if err != nil {
		log.Panicln(err)
	}

	for pagination.NextPageOptions != nil {
		var p []rest.Product
		opts.PageInfo = pagination.NextPageOptions.PageInfo
		p, pagination, err = c.client.Product.ListWithPagination(opts)
		if e, ok := err.(rest.RateLimitError); ok {
			timeout := e.RetryAfter
			log.Printf("Warning! Rate limit error. Retrying after %d seconds", timeout)
			p, pagination, err = c.retryGetProductsPage(timeout, pagination.NextPageOptions.PageInfo)
			if err != nil {
				log.Panicln(err)
			}
		} else if err != nil {
			log.Panicln(err)
		}

		products = append(products, p...)
	}

	log.Printf("Loaded %d products\n", len(products))

	return convertProductsResponse(products)
}

func (c *Client) retryGetProductsPage(timeout int, pageInfo string) ([]rest.Product, *rest.Pagination, error) {
	time.Sleep(time.Duration(timeout) * time.Second)
	opts := rest.ListOptions{Limit: maxRemoteListLimit, PageInfo: pageInfo}
	return c.client.Product.ListWithPagination(opts)
}

func convertProductsResponse(products []rest.Product) ([]*models.Product, error) {
	res := []*models.Product{}
	for _, p := range products {
		res = append(res, &models.Product{
			ID:    p.ID,
			Title: null.StringFrom(p.Title),
		})
	}
	return res, nil
}
