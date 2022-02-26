package cmd

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// CreateTable creates a user table.
func CreateTable() error {
	sqlCommand := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username varchar(20), password VARCHAR(60), description TEXT, age INTEGER, contact TEXT, firstName TEXT, lastName TEXT, class TEXT)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand)
	return err
}

// AddNewUser insert a new user in the sql database from an Account struct.
func AddNewUser(user Account) error {
	sqlCommand := `INSERT INTO users (username, password, description, age, contact, firstName, lastName, class) VALUES (?,?,?,?,?,?,?,?)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand, user.Username, user.Password, user.Description, user.Age, user.Contact, user.FirstName, user.LastName, user.Class)

	return err
}

// GetUser gets an user from the database from an username.
func GetUser(username string) (user User, err error) {
	sqlCommand := `SELECT username, description, age, contact, firstName, lastName, class FROM users WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	err = DB.QueryRowContext(ctx, sqlCommand, username).Scan(&user.Username, &user.Description, &user.Age, &user.Contact, &user.FirstName, &user.LastName, &user.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("no row for %s", username)
		}
		return user, err
	}

	return user, nil
}

// GetAllUsers returns all users from the datbase in an User struct.
func GetAllUsers() (users []User, err error) {
	sqlCommand := `SELECT username, description, age, contact, firstName, lastName, class FROM users`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	rows, err := DB.QueryContext(ctx, sqlCommand)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Username, &user.Description, &user.Age, &user.Contact, &user.FirstName, &user.LastName, &user.Class)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdatePassword updates account information in the database.
func UpdatePassword(password string, username string) error {
	sqlCommand := `UPDATE users SET password = ? WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand, username, password)

	return err
}

// UpdateUser updates user information in the database.
func UpdateUser(user User) error {
	sqlCommand := `UPDATE users SET description = ?, age = ?, contact = ?, firstName = ?, lastName = ?, class = ? WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand, user.Description, user.Age, user.Contact, user.FirstName, user.LastName, user.Class, user.Username)

	return err
}

// UpdateUsername updates the username in the database.
func UpdateUsername(newUsername string, oldUsername string) error {
	sqlCommand := `UPDATE users SET username = ? WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand, newUsername, oldUsername)

	return err
}

// DeleteUser deletes an user from the database by the username.
func DeleteAccount(username string) error {
	sqlCommand := `DELETE FROM users WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, sqlCommand, username)

	return err
}

// CheckUsername checks if an user already exists in the database. Returns true if an user already exists.
func CheckUsername(username string) (exists bool, err error) {
	sqlCommand := `SELECT username FROM users WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	err = DB.QueryRowContext(ctx, sqlCommand, username).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

// GetPassword returns the password for a user.
func GetPassword(username string) (password string, err error) {
	sqlCommand := `SELECT password FROM users WHERE username = ?`

	ctx, cancelfunc := context.WithTimeout(context.Background(), SqlTimeOut)
	defer cancelfunc()

	err = DB.QueryRowContext(ctx, sqlCommand, username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no row for %s", username)
		}
		return "", err
	}

	return password, err
}
