package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	shopify "github.com/r0busta/go-shopify-graphql/v3"
	graphqlShop "github.com/r0busta/shop-data-replication/shop/graphql"
	"github.com/r0busta/shop-data-replication/storage/database"
	"github.com/r0busta/shop-data-replication/sync/handler"
	"github.com/r0busta/shop-data-replication/sync/webhooks"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	env := os.Getenv("ENV")
	if env != "" {
		err := godotenv.Load(".env." + env + ".local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if env == "dev" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Init
	shopifyClient := shopify.NewDefaultClient()
	dbConn, err := sql.Open("mysql", driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false"))
	storage := database.New(dbConn)

	if os.Getenv("PROVISION_DATA") != "" {
		log.Debugln("Bootstraping the storage with initial shop data")

		c := graphqlShop.New(shopifyClient)
		products, err := c.ListAllProducts()
		if err != nil {
			log.Fatalln(err)
		}

		err = storage.SaveProducts(products)
		if err != nil {
			log.Fatalln(err)
		}

		log.Debugln("Data bootstrapped")
	}

	if os.Getenv("RUN_SYNC_SERVICE") != "" {
		// Sync service
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		h := handler.NewDefaultHandler(storage)

		syncService := webhooks.New(h, webhooks.WithPort(port),
			webhooks.WithShopifyAPISecret(os.Getenv("STORE_API_SECRET")),
		)

		err = syncService.ProvisionSubscriptions(shopifyClient.GraphQLClient(), os.Getenv("STORE_WEBHOOKS_CALLBACK_URL"))
		if err != nil {
			log.Fatalln(err)
		}

		log.Debugln("Starting the sync service")
		log.Fatalln(syncService.Run())
	}
}
