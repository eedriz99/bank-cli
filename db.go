package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func db_connect(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic("FAILED TO ENABLE FOREIGN KEY: " + err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
        name TEXT NOT NULL,
        phone TEXT NOT NULL UNIQUE,
        acct_no TEXT NOT NULL,
        acct_type TEXT NOT NULL,
		bal REAL,
        lim REAL,
		PRIMARY KEY (acct_no)
    );`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE
    NOT EXISTS transactions (
        acct TEXT NOT NULL,
        src TEXT,
        dest TEXT NOT NULL,
        amount REAL NOT NULL,name string, phone_num string, acct_type string
		trn_type TEXT NOT NULL,
        is_success INTEGER,
        created_at TEXT DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%S', 'now')),
        FOREIGN KEY (acct) REFERENCES accounts(acct_no) ON DELETE CASCADE ON UPDATE CASCADE
    );`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Database created successfully!")
	return db, nil
}

func db_create_acct(a *account, db *sql.DB) {
	query := fmt.Sprintf("INSERT INTO accounts (name, phone, acct_no, acct_type, bal, lim) VALUES (%s, %s, %s, %s, %f, %f);", a.name, a.phone_num, a.acct_num, a.acct_type, a.bal, a.limit)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully!")
}
