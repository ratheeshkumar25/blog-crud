package repositories

import (
	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	inter "github.com/ratheeshkumar25/blog-crud-api/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type BlogRepository struct {
	DB *gorm.DB
}

// Create implements interfaces.BlogRepoInter.
func (b *BlogRepository) Create(blog *model.BlogPost) error {
	if err := b.DB.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}

// GetAll implements interfaces.BlogRepoInter.
func (b *BlogRepository) GetAll() ([]model.BlogPost, error) {
	var blogs []model.BlogPost
	if err := b.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetByID implements interfaces.BlogRepoInter.
func (b *BlogRepository) GetByID(id uint) (*model.BlogPost, error) {
	var blogs model.BlogPost
	if err := b.DB.First(&blogs, id).Error; err != nil {
		return nil, err
	}
	return &blogs, nil
}

// Update implements interfaces.BlogRepoInter.
func (b *BlogRepository) Update(blog *model.BlogPost) error {
	if err := b.DB.Save(&blog).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements interfaces.BlogRepoInter.
func (b *BlogRepository) Delete(id uint) error {
	if err := b.DB.Delete(&model.BlogPost{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewBlogRepository(db *gorm.DB) inter.BlogRepoInter {
	return &BlogRepository{
		DB: db,
	}
}
