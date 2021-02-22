package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var fixtures *testfixtures.Loader
var cleaner dbcleaner.DbCleaner

var _ = BeforeSuite(func() {
	fixtures, cleaner = prepareDatabase()
})

var _ = BeforeEach(func() {
	cleaner.Clean("product")
	err := fixtures.Load()
	if err != nil {
		log.Fatalf("error creating fixtures: %s", err)
	}
})

func prepareDatabase() (*testfixtures.Loader, dbcleaner.DbCleaner) {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../migrations/mysql",
	}

	queryString := driver.MySQLBuildQueryString("user", "password", "test", "localhost", 3306, "false")
	db, err := sql.Open("mysql", queryString)
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("error applying migrations: %s", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("fixtures"),
	)
	if err != nil {
		log.Fatalf("error creating fixtures: %s", err)
	}

	err = fixtures.Load()
	if err != nil {
		log.Fatalf("error loading fixtures: %s", err)
	}

	cleaner := dbcleaner.New()
	cleaner.SetEngine(engine.NewMySQLEngine(queryString))

	return fixtures, cleaner
}

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}
