package repositories

import (
	"biller/models"
	"gorm.io/gorm"
	"strings"
)

type TagRepository struct {
	DB *gorm.DB
}

func InitTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		DB: db,
	}
}

func (r *TagRepository) GetTags(userID uint) RepositoryResult {
	var tags []models.Tag
	result := r.DB.Where("user_id", userID).Or("user_id", "NULL").Find(&tags)
	if result.Error != nil {
		return RepositoryResult{Result: result.Error}
	}
	return RepositoryResult{Result: tags}
}

func (r *TagRepository) CreateTag(tag *models.Tag) RepositoryResult {
	tag.Name = strings.ToLower(tag.Name)
	result := r.DB.Create(&tag)

	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}

	return RepositoryResult{Result: tag.ID}
}

func (r *TagRepository) DeleteTag(tagId string, userId string) RepositoryResult {
	result := r.DB.Delete(&models.Tag{}).Where("id = ? AND user_id = ?", tagId, userId)
	if result.Error != nil {
		return RepositoryResult{Error: result.Error}
	}
	return RepositoryResult{Result: result}
}
