package repository

import (
	"database/sql"

	"github.com/yourusername/simple-todo-golang/internal/models"
)

type UserRespositoryImpl struct {
	db *sql.DB
}

func NewUserRespository(db *sql.DB) *UserRespositoryImpl {
	return &UserRespositoryImpl{db: db}
}

func (repo *UserRespositoryImpl) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRespositoryImpl) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1`
	row := repo.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRespositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := repo.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRespositoryImpl) UpdateUser(user *models.User) (*models.User, error) {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := repo.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRespositoryImpl) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	return err
}

func (repo *UserRespositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = $1`
	row := repo.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
