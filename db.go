package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func newDB(dir, dbfile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	for _, sql := range createSQL {
		_, err = db.Exec(sql)
		if err != nil {
			db.Close()
			return nil, err
		}
	}
	return db, nil
}

var createSQL = []string{
	`CREATE TABLE IF NOT EXISTS emails
(id integer primary key autoincrement,
message_id text UNIQUE not null,
subject text not null,
filepath text not null,
folder_name text not null)`,

	`CREATE TABLE IF NOT EXISTS emails_references
(message text not null,
sibling text not null,
foreign key(message) references emails(message_id),
foreign key(sibling) references emails(message_id));   `,
}
