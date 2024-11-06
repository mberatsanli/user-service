package Controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"user/app/Http/Request"
	"user/app/Http/Response"
	"user/app/Repositories"
)

var validate = validator.New()

func Login(c *fiber.Ctx) error {
	request := new(Request.LoginRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if err := validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)

		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = fieldErr.Tag()
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user, _ := Repositories.GetUserByEmail(request.Email)
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "AUTH-T0",
			"message": "Invalid credentials",
		})
	}

	res, err := Response.LoginResponse(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "AUTH-T1",
			"message": "Invalid credentials",
		})
	}

	return c.JSON(res)
}
