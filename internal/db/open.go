package db

import (
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// An in-memory database belongs to its connection, so a pool of them is a
	// pool of different databases: the schema lands on whichever connection ran
	// it and every other one sees "no such table". Pinning to a single
	// connection makes ":memory:" behave like the file database it stands in
	// for in tests. Harmless in production, where the path is a real file and
	// the app runs one replica against one RWO volume anyway.
	db.SetMaxOpenConns(1)

	if _, err = db.Exec(`PRAGMA journal_mode=WAL; PRAGMA foreign_keys=ON;`); err != nil {
		return nil, err
	}
	if _, err = db.Exec(schema); err != nil {
		return nil, err
	}
	if _, err = db.Exec(`UPDATE proposal_leads SET created_at = datetime('now') WHERE created_at IS NULL OR created_at = '';`); err != nil {
		return nil, err
	}
	return db, nil
}
