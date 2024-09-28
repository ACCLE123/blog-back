package mapper

import (
	"blog/model"
	"gorm.io/gorm"
)

type BlogRepository struct {
	DB *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		DB: db,
	}
}

func (r *BlogRepository) FindAllBlogs() ([]model.Blog, error) {
	var blogs []model.Blog
	if err := r.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}
