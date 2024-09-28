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

func (r *BlogRepository) FindBlogByID(id uint) (*model.Blog, error) {
	var blog model.Blog
	if err := r.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}
