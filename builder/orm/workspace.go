package orm

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/oxodao/gradient-builder/model"
)

type WorkspaceORM struct {
	db *sqlx.DB
}

func (o *WorkspaceORM) UpsertWorkspace(name string) (*model.Workspace, error) {
	ws := &model.Workspace{}

	row := o.db.QueryRowx("SELECT * FROM workspace WHERE name = ?", name)
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			row := o.db.QueryRowx("INSERT INTO workspace (name) VALUES (?) RETURNING id", name)
			if row.Err() != nil {
				return nil, row.Err()
			}

			ws.LauncherName = name
			err := row.Scan(ws.ID)
			if err != nil {
				return nil, err
			}

			return ws, nil
		}

		return nil, row.Err()
	}

	return nil, nil
}
