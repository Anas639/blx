package database

import (
	"database/sql"
	"os"
	"path"
	"strconv"
)

const DB_VERSION = 1

func InitDB() (*sql.DB, error) {
	dbDir := path.Join(os.Getenv("HOME"), ".local/share/blx")
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.Mkdir(dbDir, os.FileMode(0766))
		if err != nil {
			return nil, err
		}
	}
	db, err := sql.Open("sqlite3", path.Join(dbDir, "blx.db?_foreign_keys=on"))

	err = migrateDB(db)
	if err != nil {
		return nil, err
	}
	err = createTables(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{`
		create table if not exists projects
		(id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
		)
		`,
		`
		create table if not exists tasks 
		(id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TIMESTAMP default CURRENT_TIMESTAMP,
		status TEXT CHECK(status IN('new','ongoing','paused','ended')) DEFAULT "new",
		project_id TEXT,
		FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
		)`,

		`
		create table if not exists task_sessions
		(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		start_time TIMESTAMP default CURRENT_TIMESTAMP,
		end_time TIMESTAMP,
		task_id TEXT NOT NULL,
		UNIQUE (id, task_id),
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
		)
		`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateDB(db *sql.DB) error {
	var currentVersion int
	res := db.QueryRow("pragma user_version;")
	if res.Err() != nil {
		return res.Err()
	}
	res.Scan(&currentVersion)
	_, err := db.Exec("pragma user_version = " + strconv.Itoa(DB_VERSION))

	return err
}
