package models

import()

type Category struct{
	ID 		int
	Name 	string
}

type CategoryRepository interface{
	GetAllCategories() ([]*Category, error)
	GetCategoryByID(id int) (Category, error)
	CreateCategory(Category) error
}