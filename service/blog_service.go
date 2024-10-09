package service

import (
	"blog/mapper"
	"blog/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type BlogService struct {
	Repo *mapper.BlogRepository
}

func NewBlogService(repo *mapper.BlogRepository) *BlogService {
	return &BlogService{
		Repo: repo,
	}
}

func (s *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return s.Repo.FindAllBlogs(s.Repo.DB)
}

func (s *BlogService) GetBlogByID(id uint) (*model.Blog, error) {
	blog, err := s.Repo.FindBlogByID(s.Repo.DB, id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *BlogService) AddBlog(blog *model.Blog) error {
	blog.CreatedAt = time.Now()
	return s.Repo.Create(s.Repo.DB, blog)
}

func (s *BlogService) UpdateBlog(blog *model.Blog) error {
	return s.Repo.UpdateBlogByID(s.Repo.DB, blog)
}

func (s *BlogService) DeleteBlog(id uint) error {
	return s.Repo.DeleteBlogByID(s.Repo.DB, id)
}

func (s *BlogService) UpdateOrAddBlog(blog *model.Blog) error {
	tx := s.Repo.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	res, err := s.Repo.FindBlogByID(tx, blog.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if res != nil {
		if err := s.Repo.UpdateBlogByID(tx, blog); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := s.Repo.Create(tx, blog); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
