package db

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		err := fmt.Errorf("unable to connect to postgres : %s", DBStatus)
		fmt.Printf("[ERROR] %s\n", err.Error())
	}
	return db
}

type Resolver struct {
	DB *pg.DB
}
