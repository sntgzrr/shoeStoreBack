package repositories

import (
	"errors"
	"storeApiRest/database"
	"storeApiRest/models"
)

func CreateUser(user models.User) error {
	query := `INSERT INTO users(user_full_name, user_email, user_password)
				VALUES ($1, $2, $3)`
	db := database.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.UserFullName, user.UserEmail, user.UserPassword)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}

func ReadUsers() (models.Users, error) {
	var users models.Users
	query := `SELECT user_id, user_full_name, user_email, user_password, user_created_at, user_updated_at
				FROM users`
	db := database.GetConnection()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		rows.Scan(
			&user.UserID,
			&user.UserFullName,
			&user.UserEmail,
			&user.UserPassword,
			&user.UserCreatedAt,
			&user.UserUpdatedAt,
		)
		users = append(users, &user)
	}
	return users, nil
}

func UpdateUser(user models.User, userID int) error {
	query := `UPDATE users
				SET user_full_name = $1, user_email = $2, user_password = $3, user_updated_at = now()
					WHERE user_id = $4`
	db := database.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.UserFullName, user.UserEmail, user.UserPassword, userID)
	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("more than 1 row was affected")
	}
	return nil
}
