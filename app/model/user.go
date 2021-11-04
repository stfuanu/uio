package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	ID        uint32    `db:"id" bson:"id,omitempty"` // Don't use Id, use UserID() instead for consistency with MongoDB
	FirstName string    `db:"first_name" bson:"first_name"`
	LastName  string    `db:"last_name" bson:"last_name"`
	Email     string    `db:"email" bson:"email"`
	Password  string    `db:"password" bson:"password"`
	StatusID  uint8     `db:"status_id" bson:"status_id"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   uint8     `db:"deleted" bson:"deleted"`
}

// UserStatus table contains every possible user status (active/inactive)
type UserStatus struct {
	ID        uint8     `db:"id" bson:"id"`
	Status    string    `db:"status" bson:"status"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   uint8     `db:"deleted" bson:"deleted"`
}

// UserID returns the user id
func (u *User) UserID() string {
	r := fmt.Sprintf("%v", u.ID)
	return r
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	// var errr error

	// Learnt here :
	// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/

	// sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	// sqlStatement := `SELECT id, password, status_id, first_name FROM user WHERE email=$1;`

	var user User

	db, err := sql.Open("mysql", "ding:Pass_stfu404@/auth")
	if err != nil {
		log.Println("SQL Driver Error", err) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// from here : https://golang.org/doc/database/querying

	// var db *sql.DB

	errr := db.QueryRow("SELECT id, password, status_id, first_name FROM user WHERE email = ?", email).Scan(&user.ID, &user.Password, &user.StatusID, &user.FirstName)

	// fmt.Println(rows, user, errr)

	// fmt.Println(errr, standardizeError(errr), sql.ErrNoRows)

	return user, errr

}

// UserCreate creates user
func UserCreate(firstName, lastName, email, password string) error {
	var err error

	db, err := sql.Open("mysql", "ding:Pass_stfu404@/auth")
	if err != nil {
		log.Println("SQL Driver Error", err) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// now := time.Now()
	_, err = db.Exec("INSERT INTO user (first_name, last_name, email, password) VALUES (?,?,?,?)", firstName, lastName, email, password)

	fmt.Println("from usercreate", err)
	return standardizeError(err)
}
