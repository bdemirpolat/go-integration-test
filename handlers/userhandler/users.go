package userhandler

import (
	"github.com/bdemirpolat/integration-test/models"
	"github.com/bdemirpolat/integration-test/repository"
	"github.com/gofiber/fiber/v2"
)

type UHandler interface {
	Create(c *fiber.Ctx) error
}

type UserHandler struct {
	Repo repository.UserRepository
}

// Create creates new user
func (u *UserHandler) Create(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	err = u.Repo.Create(user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
