package repository

import (
	"api-iotdoor/helper"
	"api-iotdoor/model/structure"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"
)

type UsersRepositoryInterface interface {
	SignUp(ctx context.Context, tx *sql.Tx, users structure.Users) structure.Users
	SignIn(ctx context.Context, tx *sql.Tx, username string, password string) (structure.Users, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (structure.Users, error)
	Update(ctx context.Context, tx *sql.Tx, users structure.Users) structure.Users
	Delete(ctx context.Context, tx *sql.Tx, user structure.Users)
}

type UsersImpl struct {
}

func NewUsersRepository() UsersRepositoryInterface {
	return &UsersImpl{}
}

var hash = sha256.New()

func (repository *UsersImpl) SignUp(ctx context.Context, tx *sql.Tx, user structure.Users) structure.Users {

	SQL := "INSERT INTO users(username, password, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	// Get the current time
	now := time.Now()

	hash.Write([]byte(user.Password))

	result, err := tx.ExecContext(ctx, SQL, user.Name, hex.EncodeToString(hash.Sum(nil)), user.Name, now, now)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UsersImpl) SignIn(ctx context.Context, tx *sql.Tx, username string, password string) (structure.Users, error) {

	SQL := "SELECT * FROM users WHERE username = ? LIMIT 1"

	rows, err := tx.QueryContext(ctx, SQL, username)

	helper.PanicIfError(err)

	user := structure.Users{}

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		//hash input password
		hash.Write([]byte(password))
		hashedInput := hex.EncodeToString(hash.Sum(nil))

		// Verifying password
		if hashedInput != user.Password {
			return structure.Users{}, errors.New("Password Salah")
		}

		return user, nil
	} else {
		return user, errors.New("category is not found")
	}
}

func (repository *UsersImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (structure.Users, error) {
	SQL := "select * from category where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, userId)

	helper.PanicIfError(err)

	user := structure.Users{}
	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.CreatedAt, &user.UpdatedAt)

		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("category is not found")
	}
}

func (repository *UsersImpl) Update(ctx context.Context, tx *sql.Tx, user structure.Users) structure.Users {
	SQL := "UPDATE users SET username = ?, password = ?, name = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Name, time.Now(), user.Id)

	helper.PanicIfError(err)

	return user
}

func (repository *UsersImpl) Delete(ctx context.Context, tx *sql.Tx, user structure.Users) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)

	helper.PanicIfError(err)
}
