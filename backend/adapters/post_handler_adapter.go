package adapters

import (
	"show-my-performance/backend/core"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService core.PostService
}

func NewPostHandler(postService core.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}
	post, err := h.postService.GetPostByID(uint(parsedID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var post core.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	createdPost, err := h.postService.CreatePost(&post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	var post core.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	updatedPost, err := h.postService.UpdatePost(&post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(updatedPost)
}

func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}
	if err := h.postService.DeletePost(uint(parsedID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *PostHandler) GetPostsByUserID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	parsedUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID format"})
	}
	posts, err := h.postService.GetPostsByUserID(uint(parsedUserID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}
