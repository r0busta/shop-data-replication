# shop-data-replication

Replicate your Shopify data.
Use a SQL database and unleash relational database super powers.

## How it works

The sync service handles changes in the Shopify product data via webhooks and saves updates into the database.

On the startup, product data is, optionally, backfilled from Shopify to the database.

## Getting started

```bash
sqlboiler mysql
docker-compose up
sql-migrate up
go test ./...
go run .
```
