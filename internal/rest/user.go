package rest

import (
	"github.com/seefmitrais/go-rest-api-practice/internal/repository"
	"github.com/seefmitrais/go-rest-api-practice/internal/service"

	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
)

// UserHandler ...
type UserHandler struct {
	service service.UserService
}

// NewUserHandler ...
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{service: *userService}
}

/**
	Register User Handler Routes
**/
func (handler *UserHandler) RegisterRoutes(r fiber.Router) fiber.Router {
	group := r.Group("/users")
	group.Post("/register", handler.RegisterNewUser)
	return r
}

/**
	Register New User Validation
**/
type InputRegisterNewUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (input InputRegisterNewUser) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Email, validation.Required, is.Email, validation.Length(5, 20)),
		validation.Field(&input.Password, validation.Required, validation.Length(8, 50)),
	)
}

/**
	Register New User Handler
**/
func (u UserHandler) RegisterNewUser(c *fiber.Ctx) error {

	input := new(InputRegisterNewUser)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := repository.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err := u.service.CreateNewUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}
