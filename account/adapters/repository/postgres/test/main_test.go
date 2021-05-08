package pgrepo_test

import (
	"os"
	"testing"

	irepository "github.com/idirall22/crypto_app/account/adapters/repository"
	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/jmoiron/sqlx"
)

var (
	repoTest irepository.IRepository

	dbTest *sqlx.DB
)

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../../../../")
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect(cfg.RepositoryConfig())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbTest = db
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	repoTest = pgrepo.NewPostgresRepo(db)

	os.Exit(m.Run())
}
