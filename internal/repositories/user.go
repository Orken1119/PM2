package repositories

import (
	"context"
	"time"

	models "github.com/Orken1119/PM2/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) models.UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(c context.Context, user models.UserRequest) (int, error) {
	var userID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO users(
		email, password, roleid,created_at, balance)
		VALUES ($1, $2, $3, $4, $5) returning id;`

	err := ur.db.QueryRow(c, userQuery, user.Email, user.Password.Password, 2, currentTime, 0).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (ur *UserRepository) GetUserByEmail(c context.Context, email string) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, roleid, created_at, balance FROM users where email = $1`
	row := ur.db.QueryRow(c, query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.Balance)

	if err != nil {
		return user, err
	}
	return user, nil

}

func (ur *UserRepository) GetUserByID(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, phone_number, roleid, created_at, balance FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.Balance)

	if err != nil {
		return user, err
	}
	return user, nil

}

func (ur *UserRepository) GetProfile(c context.Context, userID int) (models.UserProfile, error) {
	user := models.UserProfile{}

	query := `SELECT id,user_name, email, balance FROM users WHERE id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Balance)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) EditPersonalData(c context.Context, userID int, name string, email string, phoneNumber string, birthday string) error {
	query := `UPDATE users
	SET 
		user_name = $1,
		email = $2
	WHERE 
	    id = $5;
	`
	_, err := ur.db.Exec(c, query, name, email, phoneNumber, birthday, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) AddUser(c context.Context, name string, balance int) (int, error) {
	var userID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO users(
		email, password, roleid,created_at, balance)
		VALUES ($1, $2, $3, $4, $5) returning id;`

	err := ur.db.QueryRow(c, userQuery, "", "", 2, currentTime, balance).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}