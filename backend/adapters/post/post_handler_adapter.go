package postAdapter

import (
	postCore "show-my-performance/backend/core/post"
	"show-my-performance/backend/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService postCore.PostService
}

func NewPostHandler(postService postCore.PostService) *PostHandler {
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
	post := model.Post{}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to parse form data"})
	}
	files := form.File["image"]
	createdPost, err := h.postService.CreatePost(&post, files)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdPost)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var post model.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	post.ID = uint(id)
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
