package graphqlShop

import (
	"fmt"
	"strconv"

	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/shop"
	"github.com/volatiletech/null/v8"

	shopify "github.com/r0busta/go-shopify-graphql-model/graph/model"
	graphql "github.com/r0busta/go-shopify-graphql/v3"
)

type Client struct {
	client *graphql.Client
}

func New(graphqlClient *graphql.Client) shop.Client {
	return &Client{
		client: graphqlClient,
	}
}

const LIST_ALL_PRODUCTS = `
{
	products{
		edges{
			node{
				legacyResourceId
				title
			}
		}
	}
}
`

func (c *Client) ListAllProducts() ([]*models.Product, error) {
	data := []*shopify.Product{}
	err := c.client.BulkOperation.BulkQuery(LIST_ALL_PRODUCTS, &data)
	if err != nil {
		return []*models.Product{}, err
	}

	return convertProductsResponse(data)
}

func convertProductsResponse(products []*shopify.Product) ([]*models.Product, error) {
	res := []*models.Product{}
	for _, p := range products {
		id, err := strconv.ParseInt(p.LegacyResourceID.String, 10, 64)
		if err != nil {
			return []*models.Product{}, fmt.Errorf("error parsing product legacy ID: %s", err)
		}
		res = append(res, &models.Product{
			ID:    id,
			Title: null.StringFromPtr(p.Title.Ptr()),
		})
	}
	return res, nil
}
