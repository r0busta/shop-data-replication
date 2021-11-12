package tests_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

var cleaner dbcleaner.DbCleaner

var _ = BeforeSuite(func() {
	cleaner = prepareDatabase()
})

var _ = BeforeEach(func() {
	cleaner.Clean("product")
})

func prepareDatabase() dbcleaner.DbCleaner {
	migrations := &migrate.FileMigrationSource{
		Dir: "../../../migrations/mysql",
	}

	queryString := driver.MySQLBuildQueryString("user", "password", "shop", "localhost", 3306, "false")
	db, err := sql.Open("mysql", queryString)
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("error applying migrations: %s", err)
	}

	log.Printf("Applied %d migrations!\n", n)

	cleaner := dbcleaner.New()
	cleaner.SetEngine(engine.NewMySQLEngine(queryString))

	return cleaner
}

func TestMysql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysql Suite")
}
