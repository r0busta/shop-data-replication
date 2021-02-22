package main

import (
	"database/sql"
	"os"

	"github.com/r0busta/shop-data-replication/handler"
	graphqlShop "github.com/r0busta/shop-data-replication/shop/graphql"
	"github.com/r0busta/shop-data-replication/storage/database"
	"github.com/r0busta/shop-data-replication/sync/webhooks"

	"github.com/joho/godotenv"
	shopify "github.com/r0busta/go-shopify-graphql/v3"
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
	db := database.New(dbConn)

	if os.Getenv("PROVISION_DATA") != "" {
		log.Debugln("Bootstraping the storage with initial shop data")
		c := graphqlShop.New(shopifyClient)
		products, err := c.ListAllProducts()
		if err != nil {
			log.Fatalln(err)
		}

		err = db.SaveProducts(products)
		if err != nil {
			log.Fatalln(err)
		}
		log.Debugln("Data bootstrapped")
	}

	// Sync service
	h := handler.NewDefaultHandler(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := webhooks.New(h, webhooks.WithPort(port),
		webhooks.WithShopifyApiSecret(os.Getenv("STORE_API_SECRET")),
	)

	err = s.(*webhooks.Service).ProvisionSubscriptions(shopifyClient.GraphQLClient(), os.Getenv("STORE_WEBHOOKS_CALLBACK_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Debugln("Starting sync service")
	log.Fatalln(s.Run())
}
