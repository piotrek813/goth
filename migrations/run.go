package migrations

import (
	"fmt"
	"piotrek813/goth/db"
)

func createTable(table string, definition string) {
	db := db.GetDB()

	_, err := db.Exec(fmt.Sprintf(` CREATE TABLE IF NOT EXISTS %v ( %v ) `, table, definition))

	if err != nil {
		panic(err)
	}
}

func RunMigrations() {
	createTable("Users", ` 
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			login TEXT ,
			password TEXT,
			UNIQUE(id)`)

	createTable("Posts", ` 
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			title VARCHAR(255) NOT NULL,
			body TEXT NOT NULL,
			UNIQUE (id)
			FOREIGN KEY(user_id) REFERENCES Users(id)`)

	createTable("Collections", `
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			name VARCHAR(255) NOT NULL,
			UNIQUE (id) `)

	createTable("Collections_Inputs", `
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			collection_id INTEGER,
			text_input_id INTEGER,
			number_input_id INTEGER,
			textarea_input_id INTEGER,

			FOREIGN KEY(text_input_id ) REFERENCES TextAreaInputs(id),
			FOREIGN KEY(number_input_id ) REFERENCES NumberInputs(id),
			FOREIGN KEY(textarea_input_id) REFERENCES TextInputs(id),
			FOREIGN KEY(collection_id) REFERENCES Collections(id),
			UNIQUE (id) `)

	// -- Inputs --

	createTable("TextInputs", `
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			name VARCHAR(255) NOT NULL,
			value TEXT,
			UNIQUE (id) `)

	createTable("TextAreaInputs", `
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			name VARCHAR(255) NOT NULL,
			value TEXT,
			UNIQUE (id) `)

	createTable("NumberInputs", `
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME NOT NULL DEFAULT current_timestamp,
			name VARCHAR(255) NOT NULL,
			value INTEGER,
			UNIQUE (id) `)
}
