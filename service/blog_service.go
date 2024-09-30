package service

import (
	"blog/mapper"
	"blog/model"
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
	return s.Repo.FindAllBlogs()
}

func (s *BlogService) GetBlogByID(id uint) (*model.Blog, error) {
	blog, err := s.Repo.FindBlogByID(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *BlogService) AddBlog(blog *model.Blog) error {
	blog.CreatedAt = time.Now()
	return s.Repo.Create(blog)
}
