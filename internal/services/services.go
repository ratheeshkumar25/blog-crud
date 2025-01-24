package services

import (
	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	int "github.com/ratheeshkumar25/blog-crud-api/internal/repositories/interfaces"
	inter "github.com/ratheeshkumar25/blog-crud-api/internal/services/interfaces"
)

type BlogService struct {
	Repo int.BlogRepoInter
}

// Create implements interfaces.BlogServiceInter.
func (b *BlogService) Create(blog *model.BlogPost) error {
	return b.Repo.Create(blog)
}

// GetAll implements interfaces.BlogServiceInter.
func (b *BlogService) GetAll() ([]model.BlogPost, error) {
	return b.Repo.GetAll()
}

// GetByID implements interfaces.BlogServiceInter.
func (b *BlogService) GetByID(id uint) (*model.BlogPost, error) {
	return b.Repo.GetByID(id)
}

// Update implements interfaces.BlogServiceInter.
func (b *BlogService) Update(blog *model.BlogPost) error {
	return b.Repo.Update(blog)
}

// Delete implements interfaces.BlogServiceInter.
func (b *BlogService) Delete(id uint) error {
	return b.Repo.Delete(id)
}

func NewBlogServices(repo int.BlogRepoInter) inter.BlogServiceInter {
	return &BlogService{
		Repo: repo,
	}
}
