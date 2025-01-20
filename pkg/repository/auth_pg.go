package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	app "github.com/rommomm123321/go-test-backend"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.User) (app.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, email, password) values ($1, $2, $3) ON CONFLICT (email) DO NOTHING RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&user.Id); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthPostgres) GetUser(email, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
