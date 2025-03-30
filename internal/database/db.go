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
	dbPath := path.Join(dbDir, "blx.db")
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.Mkdir(dbDir, os.FileMode(0766))
		if err != nil {
			return nil, err
		}
	}
	_, err := os.Stat(dbPath)
	dbExists := err == nil
	db, err := sql.Open("sqlite3", path.Join(dbDir, "blx.db?_foreign_keys=on"))

	err = migrateDB(db)
	if err != nil {
		return nil, err
	}
	if !dbExists {
		err = createTables(db)
		if err != nil {
			return nil, err
		}
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
		`create virtual table if not exists tasks_fts using fts5(name)`,
		`create virtual table if not exists projects_fts using fts5(name)`,
		`drop trigger if exists tasks_ai`,
		`drop trigger if exists tasks_au`,
		`drop trigger if exists tasks_ad`,
		`
		create trigger tasks_ai after insert on tasks
		begin
			insert into tasks_fts (rowid, name) values (new.id, new.name);
		end
		`,
		`
		create trigger tasks_au after update on tasks
		begin
			delete from tasks_fts where rowid = old.id;
			insert into tasks_fts (rowid,name) values(new.id, new.name);
		end
		`,
		`
		create trigger tasks_ad after delete on tasks
		begin
			delete from tasks_fts where rowid = old.id;
		end
		`,
		`drop trigger if exists projects_ai`,
		`drop trigger if exists projects_au`,
		`drop trigger if exists projects_ad`,
		`
		create trigger projects_ai after insert on projects
		begin
			insert into projects_fts (rowid, name) values (new.id, new.name);
		end
		`,
		`
		create trigger projects_au after update on projects
		begin
			delete from projects_fts where rowid = old.id;
			insert into projects_fts (rowid,name) values(new.id, new.name);
		end
		`,
		`
		create trigger projects_ad after delete on projects
		begin
			delete from projects_fts where rowid = old.id;
		end
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
