package mysql

import (
	"database/sql"
	"github.com/anggardagasta/mini_wallet/models"
	"github.com/anggardagasta/mini_wallet/service"
)

type serviceUsersRepository struct {
	DB *sql.DB
}

func NewServiceUsersRepository(db *sql.DB) service.IServiceUsersRepository {
	return serviceUsersRepository{DB: db}
}

func (repo serviceUsersRepository) GetUserByUsernameOrEmail(form models.FormRegister) (result models.GetUserScanner, err error) {
	q := `SELECT id, username, email, password FROM users WHERE username = ? OR email = ?`
	rows, err := repo.DB.Query(q, form.Username, form.Email)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Username, &result.Email, &result.Password); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) GetUserByUsername(username string) (result models.GetUserScanner, err error) {
	q := `SELECT id, username, email, password FROM users WHERE username = ?`
	rows, err := repo.DB.Query(q, username)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Username, &result.Email, &result.Password); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) GetUserByEmail(email string) (result models.GetUserScanner, err error) {
	q := `SELECT id, username, email, password FROM users WHERE email = ?`
	rows, err := repo.DB.Query(q, email)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Username, &result.Email, &result.Password); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) InsertUser(input models.FormRegister) (id int64, err error) {
	stmt, err := repo.DB.Prepare(`INSERT INTO users (username, email, password) VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(input.Username, input.Email, input.Password)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
