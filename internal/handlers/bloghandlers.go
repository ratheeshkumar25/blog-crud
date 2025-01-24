package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/blog-crud-api/internal/model"
	inter "github.com/ratheeshkumar25/blog-crud-api/internal/services/interfaces"
)

type BlogHandler struct {
	SVC inter.BlogServiceInter
}

func NewBlogHandler(router fiber.Router, svc inter.BlogServiceInter) {
	handler := &BlogHandler{SVC: svc}
	router.Post("blog-post", handler.Create)
	router.Get("blog-post", handler.GetAll)
	router.Get("blog-post/:id", handler.GetByID)
	router.Patch("blog-post/:id", handler.Update)
	router.Delete("blog-post/:id", handler.Delete)
}

// Create creates a new blog post.
// @Summary Create a blog post
// @Description Create a new blog post with title and content
// @Tags blog
// @Accept json
// @Produce json
// @Param blog body model.BlogPost true "Blog post data"
// @Success 201 {object} model.BlogPost
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/blog-post [post]
func (h *BlogHandler) Create(c *fiber.Ctx) error {
	blog := new(model.BlogPost)
	if err := c.BodyParser(blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := h.SVC.Create(blog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(blog)
}

// GetAll retrieves all blog posts.
// @Summary Get all blog posts
// @Description Retrieve all blog posts in the database
// @Tags blog
// @Accept json
// @Produce json
// @Success 200 {array} model.BlogPost
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/blog-post [get]
func (h *BlogHandler) GetAll(c *fiber.Ctx) error {
	blogs, err := h.SVC.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(blogs)
}

// GetByID retrieves a blog post by ID.
// @Summary Get a blog post by ID
// @Description Retrieve a single blog post by its unique ID
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "Blog post ID"
// @Success 200 {object} model.BlogPost
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/blog-post/{id} [get]
func (h *BlogHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := h.SVC.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.JSON(blog)
}

// Update updates a blog post by ID.
// @Summary Update a blog post
// @Description Update a blog post's title and content by its ID
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "Blog post ID"
// @Param blog body model.BlogPost true "Updated blog post data"
// @Success 200 {object} model.BlogPost
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/blog-post/{id} [patch]
func (h *BlogHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := new(model.BlogPost)
	if err := c.BodyParser(blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	blog.ID = uint(id)
	if err := h.SVC.Update(blog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(blog)
}

// Delete deletes a blog post by ID.
// @Summary Delete a blog post
// @Description Delete a blog post by its unique ID
// @Tags blog
// @Accept json
// @Produce json
// @Param id path int true "Blog post ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/blog-post/{id} [delete]
func (h *BlogHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.SVC.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
