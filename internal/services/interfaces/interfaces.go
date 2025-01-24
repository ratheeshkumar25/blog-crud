package interfaces

import "github.com/ratheeshkumar25/blog-crud-api/internal/model"

type BlogServiceInter interface {
	Create(blog *model.BlogPost) error
	GetAll() ([]model.BlogPost, error)
	GetByID(id uint) (*model.BlogPost, error)
	Update(blog *model.BlogPost) error
	Delete(id uint) error
}
