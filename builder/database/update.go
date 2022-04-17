package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func UpdateDatabaseIfRequired(db *sqlx.DB) {
	fmt.Println("@TODO: migrations")
}
