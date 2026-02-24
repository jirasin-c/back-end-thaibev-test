package db

import (
	"fmt"
	"time"

	"thaibev-test/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var pgCON *sqlx.DB = nil

func Connection() *sqlx.DB {
	if pgCON == nil {
		con := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
			env.PG_HOST,
			env.PG_NAME,
			env.PG_PORT,
			env.PG_USER,
			env.PG_PASSWORD,
		)

		db, err := sqlx.Connect("postgres", con)

		if err != nil {
			fmt.Println("connect error:", err)
			return nil
		}

		db.SetMaxOpenConns(150)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(time.Hour)

		pgCON = db

	}
	return pgCON
}
