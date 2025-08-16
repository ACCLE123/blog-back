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

func (r *BlogRepository) FindAllBlogs(tx *gorm.DB) ([]model.Blog, error) {
	var blogs []model.Blog
	if err := tx.Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepository) FindBlogByID(tx *gorm.DB, id uint) (*model.Blog, error) {
	var blog model.Blog
	if err := tx.First(&blog, id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *BlogRepository) UpdateBlogByID(tx *gorm.DB, blog *model.Blog) error {
	return tx.Model(&model.Blog{}).Where("id = ?", blog.ID).Updates(blog).Error
}

func (r *BlogRepository) Create(tx *gorm.DB, blog *model.Blog) error {
	return tx.Create(blog).Error
}

func (r *BlogRepository) DeleteBlogByID(tx *gorm.DB, id uint) error {
	return tx.Delete(&model.Blog{}, id).Error
}
