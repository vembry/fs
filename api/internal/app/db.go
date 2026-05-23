package app

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func NewDb(sqliteDdlSchema string) *sql.DB {
	conn, err := sql.Open("sqlite", "/db/fs.db")
	if err != nil {
		log.Fatalf("error on opening connection to sqlite. err=%v", err)
	}

	sqlitePostStart(conn, sqliteDdlSchema)

	return conn
}

func sqlitePostStart(conn *sql.DB, ddlSchema string) {
	if _, err := conn.Exec("PRAGMA journal_mode=WAL; PRAGMA foreign_keys=ON;"); err != nil {
		log.Fatal("Failed to set PRAGMAs:", err)
	}

	// create tables
	if _, err := conn.Exec(ddlSchema); err != nil {
		log.Fatal("Failed to initialize database tables:", err)
	}
}
