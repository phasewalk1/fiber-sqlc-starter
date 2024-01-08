package handlers

import (
	"net/http"

	"github.com/phasewalk1/fiber-sqlc-starter/database"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user database.User
	req := &createUserRequest{}

	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	hash, err := user.HashPassword(req.User.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser, err := h.queries.CreateUser(c.Context(), database.CreateUserParams{
		Fname: req.User.Fname,
		Lname: req.User.Lname,
		Email: req.User.Email,
		Hash:  hash,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser.Hash = ""

	return c.Status(http.StatusCreated).JSON(newUser)

}
