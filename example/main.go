package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Patrick-Batenburg/nullify/null"
)

// Define a type that can have nullable fields, and
// will be used as our database result.
type User struct {
	ID                 null.UUID   `json:"id"`
	FirstName          null.String `json:"firstName"`
	MiddleName         null.String `json:"middleName"`
	LastName           null.String `json:"lastName"`
	Age                null.Int8   `json:"age"`
	Email              null.String `json:"email"`
	CreatedAt          null.Time   `json:"createdAt"`
	SomeOptionalID     null.UUID   `json:"someOptionalId"`
	SomeOptionalNumber null.Int64  `json:"someOptionalNumber"`
	SomeOptionalTime   null.Time   `json:"someOptionalTime"`
}

var (
	db *sql.DB
)

func init() {
	// Initialize SQLite database
	var err error
	db, err = sql.Open("sqlite3", ":memory:")

	if err != nil {
		panic(err)
	}
}

func main() {
	defer db.Close()

	// Example JSON input
	data := `
{
  "id": "09e9f207-d3b6-4fb0-8dd0-fb45025ebbdf",
  "firstName": "John",
  "middleName": null,
  "lastName": "Doe",
  "age": 1,
  "email": "johndoe@example.com",
  "createdAt": "2006-01-02T15:04:05-07:00"
}`

	user, err := convertToUserModel(data)
	// Sqlite3 driver only accepts int64 on types implementing the driver.Valuer interface.
	// Which is why we are choosing WithInt64Valuer option to convert the value to int64.
	user.Age = null.Int8From(user.Age.ValueOrZero(), null.WithInt64Valuer())
	if err != nil {
		log.Fatal(err)
	}

	err = createUser(user)

	if err != nil {
		log.Fatal(err)
	}

	userFromDB, err := getUserFromDB()

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJSON(userFromDB)
}

func convertToUserModel(data string) (user User, err error) {
	return user, json.Unmarshal([]byte(data), &user)
}

func createUser(user User) (err error) {
	// SQL statements
	insertStmt := `
INSERT INTO user (
        id,
        first_name,
        middle_name,
        last_name,
        age,
        email,
        created_at,
        some_optional_id,
        some_optional_number,
		some_optional_time
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Create user table
	_, err = db.Exec(`
CREATE TABLE user (
    id TEXT PRIMARY KEY,
    first_name TEXT,
    middle_name TEXT,
    last_name TEXT,
    age INTEGER,
    email TEXT,
    created_at TEXT,
    some_optional_id TEXT,
    some_optional_number INTEGER,
    some_optional_time TEXT
)`)

	if err != nil {
		return err
	}

	// Insert user into the database using the `null` types from `nullify/null`
	_, err = db.Exec(
		insertStmt,
		user.ID,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.Age,
		user.Email,
		user.CreatedAt,
		user.SomeOptionalID,
		user.SomeOptionalNumber,
		user.SomeOptionalTime,
	)

	return err
}

func getUserFromDB() (user User, err error) {
	selectStmt := `SELECT * FROM user`
	rows, err := db.Query(selectStmt)

	if err != nil {
		log.Fatal(err)
	}

	for {
		ok := rows.Next()

		if !ok {
			break
		}

		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.Age,
			&user.Email,
			&user.CreatedAt,
			&user.SomeOptionalID,
			&user.SomeOptionalNumber,
			&user.SomeOptionalTime,
		)

		if err != nil {
			return user, err
		}
	}

	return user, err
}

func prettyPrintJSON(value any) {
	data, err := json.MarshalIndent(value, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
