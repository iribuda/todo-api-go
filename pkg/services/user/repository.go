package user

import (
	"database/sql"
	"fmt"

	"guthub.com/iribuda/todo-api-go/pkg/models"
)

type UserRepositoryImpl struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepositoryImpl{
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error){
	row := ur.db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	u, err := scanRowIntoUser(row)
	if err != nil {
		return nil, err
	} 

	if u.UserID == 0 {
		fmt.Println(err)

		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func (ur *UserRepositoryImpl) GetUserByID(id int) (*models.User, error){
	row := ur.db.QueryRow("SELECT * FROM user WHERE userId = ?", id)

	u, err := scanRowIntoUser(row)
	if err != nil {
		return nil, err
	}

	if u.UserID == 0 {
		fmt.Println(err)

		return nil, fmt.Errorf("user not found")
	}

	return u, nil

}
func (ur *UserRepositoryImpl) CreateUser(user models.User) error{
	_, err := ur.db.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(row *sql.Row)(*models.User, error){
	user := new(models.User)
    if err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password); err != nil {
        return user, fmt.Errorf("get user: %v", err)
    }
    return user, nil
}
