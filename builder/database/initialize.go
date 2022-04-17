package database

import "github.com/jmoiron/sqlx"

func InitializeDB(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE database_migrations (
			version    TEXT NOT NULL,
			applied_at TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	if _, err = db.Exec(`INSERT INTO database_migrations (version, applied_at) VALUES ('2022-04-17T15:00:00Z', strftime('%Y-%m-%dT%H:%M:%SZ', 'now'));`); err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE modpack (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			version TEXT NOT NULL,

			mc_version TEXT NOT NULL,

			forge TEXT NULL,
			forge_version TEXT NULL,

			fabric_loader_version TEXT NULL

			UNIQUE(name)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
