package graphqlshop

import (
	"fmt"
	"strconv"

	shopify "github.com/r0busta/go-shopify-graphql-model/graph/model"
	graphql "github.com/r0busta/go-shopify-graphql/v3"
	"github.com/r0busta/shop-data-replication/models"
	"github.com/volatiletech/null/v8"
)

type Client struct {
	client *graphql.Client
}

func New(graphqlClient *graphql.Client) *Client {
	return &Client{
		client: graphqlClient,
	}
}

const listAllProducts = `
{
	products{
		edges{
			node{
				id
				legacyResourceId
				vendor
				title
				descriptionHtml
				totalInventory
				priceRangeV2 {
					minVariantPrice {
						amount
						currencyCode
					}
					maxVariantPrice {
						amount
						currencyCode
					}
				}
				images(first: 20) {
					edges {
						node {
							transformedSrc
						}
					}
				}
				hasOnlyDefaultVariant
				options {
					name
					values
					position
				}
				variants(first: 20) {
					edges {
						node {
							id
							legacyResourceId
							price
							compareAtPrice
							inventoryQuantity
							availableForSale
							selectedOptions {
								name
								value
							}
						}
					}
				}
			}
		}
	}
}
`

func (c *Client) ListAllProducts() ([]*models.Product, error) {
	data := []*shopify.Product{}

	err := c.client.BulkOperation.BulkQuery(listAllProducts, &data)
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
			return []*models.Product{}, fmt.Errorf("error parsing product legacy ID: %w", err)
		}

		res = append(res, &models.Product{
			ID:    id,
			Title: null.StringFromPtr(p.Title.Ptr()),
		})
	}

	return res, nil
}
