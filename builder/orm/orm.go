package orm

import "github.com/jmoiron/sqlx"

var Workspace *WorkspaceORM = nil

func Init(db *sqlx.DB) {
	Workspace = &WorkspaceORM{db: db}
}
