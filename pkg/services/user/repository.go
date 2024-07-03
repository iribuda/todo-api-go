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
	// row := ur.db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	rows, err := ur.db.Query("SELECT * FROM `user` WHERE email = ?", email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// u, err := scanRowIntoUser(row)
	// if err != nil {
	// 	return nil, err
	// }

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
		fmt.Println(err)

			return nil, err
		}
	}

	if u.UserID == 0 {
		fmt.Println(err)

		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
func (ur *UserRepositoryImpl) GetUserByID(id int) (*models.User, error){
	return nil, nil

}
func (ur *UserRepositoryImpl) CreateUser(user models.User) error{
	_, err := ur.db.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// func scanRowIntoUser(row *sql.Row)(*User, error){
// 	user := new(User)
//     if err := row.Scan(&user.Username, &user.Email, &user.Password); err != nil {
//         return user, fmt.Errorf("get user: %v", err)
//     }
//     return user, nil
// }

func scanRowIntoUser(rows *sql.Rows)(*models.User, error){
		user := new(models.User)
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
	
		return user, nil
	}