package service

import (
	"blog/mapper"
	"blog/model"
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
