package repositories

import (
	"biller/models"
	"gorm.io/gorm"
	"strings"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func InitCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) GetCategories() RepositoryResult {
	var categories []models.Category
	result := r.DB.Find(&categories)
	if result.Error != nil {
		return RepositoryResult{Result: result.Error}
	}
	return RepositoryResult{Result: categories}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) RepositoryResult {
	category.Name = strings.ToLower(category.Name)
	result := r.DB.Create(&category)

	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}

	return RepositoryResult{Result: category.ID}
}

func (r *CategoryRepository) DeleteCategory(categoryId string) RepositoryResult {
	result := r.DB.Delete(&models.Tag{}, categoryId)
	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}
	return RepositoryResult{Result: result}
}
