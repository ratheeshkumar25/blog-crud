package services_test

import (
	"testing"
	"time"

	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	"github.com/ratheeshkumar25/blog-crud-api/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockBlogRepo struct{}

func (m *MockBlogRepo) Create(blog *model.BlogPost) error {
	return nil
}

func (m *MockBlogRepo) GetAll() ([]model.BlogPost, error) {
	return []model.BlogPost{
		{
			ID:          1,
			Title:       "Testing Blog",
			Description: "Test-Description",
			Body:        "Testing Body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockBlogRepo) GetByID(id uint) (*model.BlogPost, error) {
	return &model.BlogPost{
		ID: 1, Title: "Testing Blog",
		Description: "Test-Description",
		Body:        "Testing Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockBlogRepo) Update(blog *model.BlogPost) error {
	return nil
}
func (m *MockBlogRepo) Delete(id uint) error {
	return nil
}

func TestBlogService_Create(t *testing.T) {
	mockrepo := &MockBlogRepo{}
	blogSvc := services.NewBlogServices(mockrepo)
	newBlog := &model.BlogPost{
		Title:       "New Blog",
		Description: "New Description",
		Body:        "New Body",
	}
	err := blogSvc.Create(newBlog)

	assert.NoError(t, err)
}

func TestBlogService_GetAll(t *testing.T) {
	mockRepo := &MockBlogRepo{}
	blogService := services.NewBlogServices(mockRepo)

	blogs, err := blogService.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, blogs)
	assert.Equal(t, 1, len(blogs))
	assert.Equal(t, "Testing Blog", blogs[0].Title)
}

func TestBlogService_GetByID(t *testing.T) {
	mockRepo := &MockBlogRepo{}
	blogService := services.NewBlogServices(mockRepo)

	blog, err := blogService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, blog)
	assert.Equal(t, uint(1), blog.ID)
	assert.Equal(t, "Testing Blog", blog.Title)
}

func TestBlogService_Update(t *testing.T) {
	mockRepo := &MockBlogRepo{}
	blogService := services.NewBlogServices(mockRepo)

	updatedBlog := &model.BlogPost{
		ID:          1,
		Title:       "Updated Blog",
		Description: "Updated Description",
		Body:        "Updated Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := blogService.Update(updatedBlog)

	assert.NoError(t, err)
}

func TestBlogService_Delete(t *testing.T) {
	mockRepo := &MockBlogRepo{}
	blogService := services.NewBlogServices(mockRepo)

	err := blogService.Delete(1)

	assert.NoError(t, err)
}
