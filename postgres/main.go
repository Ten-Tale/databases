package main

import (
	"database/sql"
	databaseworker "databases/postgres/databaseWorker"
	"fmt"
)

func main() {
	db := databaseworker.OpenDbConnection(" ")

	defer func() {
		db := databaseworker.OpenDbConnection(" ")

		db.Exec("DROP DATABASE university_group")

		databaseworker.CloseDbConnection(db)
	}()
	_, err := db.Exec("CREATE DATABASE university_group")

	checkError(err)

	databaseworker.CloseDbConnection(db)

	db = databaseworker.OpenDbConnection("university_group")
	defer databaseworker.CloseDbConnection(db)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id CHAR(7),
		first_name VARCHAR(20),
		last_name VARCHAR(20)
	)`)

	defer func() {
		_, err = db.Exec("DROP TABLE students")

		checkError(err)
	}()

	checkError(err)

	checkError(err)

	printStudents(db)

	_, err = db.Exec(`INSERT INTO students(id, first_name, last_name)
		VALUES('19B0544', 'Andrey', 'Shkunov')`)

	checkError(err)

	printStudents(db)

	_, err = db.Exec(`UPDATE students
	SET last_name = 'Brukhanov'
	WHERE id = '19B0544'`)

	checkError(err)

	printStudents(db)

	_, err = db.Exec(`DELETE FROM students
	WHERE id = '19B0544'`)

	checkError(err)

	printStudents(db)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// not in the database worker as the function is specific
func printStudents(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM students")

	checkError(err)

	var (
		id        string
		firstName string
		lastName  string
	)

	counter := 0

	defer rows.Close()

	for rows.Next() {
		counter++
		err := rows.Scan(&id, &firstName, &lastName)

		checkError(err)

		fmt.Println("\n", id, firstName, lastName)
	}
	err = rows.Err()

	checkError(err)

	if counter == 0 {
		fmt.Println("\nThere are no rows in this query")
	}
}
