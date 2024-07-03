package user

import (
	"database/sql"
	"fmt"
)

type UserRepositoryImpl struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepositoryImpl{
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (*User, error){
	row := ur.db.QueryRow("SELECT * FROM users WHERE email = ?", email)

		u, err := scanRowIntoUser(row)
		if err != nil {
			return nil, err
		}

	if u.UserID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func (ur *UserRepositoryImpl) GetUserByID(id int) (*User, error){
	return nil, nil

}
func (ur *UserRepositoryImpl) CreateUser(user User) error{
	_, err := ur.db.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(row *sql.Row)(*User, error){
	user := new(User)
    if err := row.Scan(&user.Username, &user.Email, &user.Password); err != nil {
        return user, fmt.Errorf("get user: %v", err)
    }
    return user, nil
}