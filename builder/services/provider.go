package services

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/oxodao/gradient-builder/database"
)

var GET *Provider = nil

type Provider struct {
	DB       *sqlx.DB
	Modrinth *Modrinth
}

func Init(workspacePath string) error {
	isNewWorkspace := false
	if _, err := os.Stat(workspacePath + "/workspace.db"); os.IsNotExist(err) {
		isNewWorkspace = true
	}

	db, err := sqlx.Open("sqlite3", workspacePath+"/workspace.db")
	if err != nil {
		return err
	}

	if isNewWorkspace {
		err = database.InitializeDB(db)
	}

	GET = &Provider{
		DB:       db,
		Modrinth: &Modrinth{},
	}

	return nil
}
