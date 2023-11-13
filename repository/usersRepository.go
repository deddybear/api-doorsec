package repository

import (
	"api-doorsec/helper"
	"api-doorsec/model/structure"
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
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) error
	Update(ctx context.Context, tx *sql.Tx, users structure.Users) structure.Users
	Delete(ctx context.Context, tx *sql.Tx, user structure.Users)
}

type UsersImpl struct {
}

func NewUsersRepository() UsersRepositoryInterface {
	return &UsersImpl{}
}

func (repository *UsersImpl) SignUp(ctx context.Context, tx *sql.Tx, user structure.Users) structure.Users {
	var hash = sha256.New()
	SQL := "INSERT INTO users(username, password, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	// Get the current time
	now := time.Now()

	hash.Write([]byte(user.Password))

	result, err := tx.ExecContext(ctx, SQL, user.Username, hex.EncodeToString(hash.Sum(nil)), user.Name, now, now)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UsersImpl) SignIn(ctx context.Context, tx *sql.Tx, username string, password string) (structure.Users, error) {
	var hash = sha256.New()
	SQL := "SELECT * FROM users WHERE username = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()
	user := structure.Users{}

	if rows.Next() {

		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		//hash input password
		hash.Write([]byte(password))

		hashedInput := hex.EncodeToString(hash.Sum(nil))

		// Verifying password
		if hashedInput != user.Password {
			return structure.Users{}, errors.New("wrong username or password")
		}

		return user, nil
	} else {
		return user, errors.New("user account is not found")
	}
}

func (respository *UsersImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) error {
	SQL := "SELECT COUNT(id) FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	var count int
	if rows.Next() {
		rows.Scan(&count)

		helper.PanicIfError(err)

		if count > 0 {
			return errors.New("username has been take, please user another username")
		} else {
			return nil
		}

	} else {
		return nil
	}
}

func (repository *UsersImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (structure.Users, error) {
	SQL := "select * from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

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
