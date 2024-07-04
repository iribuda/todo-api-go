package category

import (
	"database/sql"
	"fmt"
	"guthub.com/iribuda/todo-api-go/pkg/models"
)

type CategoryRepositoryImpl struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) *CategoryRepositoryImpl{
	return &CategoryRepositoryImpl{db: db}
}

func (cr *CategoryRepositoryImpl) GetAllCategories() ([]*models.Category, error){
	categories := make([]*models.Category, 0)
	rows, err := cr.db.Query("SELECT * FROM category")
	if (err != nil){
		return nil, fmt.Errorf("all categories: %v", err)
	}
	defer rows.Close()
	for rows.Next(){
		category := new(models.Category)
		err = rows.Scan(&category.ID, &category.Name)
		if (err != nil){
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}


func (cr *CategoryRepositoryImpl) GetCategoryByID(id int) (*models.Category, error){
	row := cr.db.QueryRow("SELECT * FROM category WHERE categoryId = ?", id)

	category := new(models.Category)
	if err := row.Scan(&category.ID, &category.Name); err != nil{
		return category, fmt.Errorf("get category: %v", err)
	}
	if category.ID == 0{
		return nil, fmt.Errorf("category not found")
	}
	return category, nil
}

func (cr *CategoryRepositoryImpl) CreateCategory(category models.Category) error{
	_, err := cr.db.Exec("INSERT INTO category (name) VALUES (?)", category.Name)
	if err != nil{
		return err
	}
	return nil
}