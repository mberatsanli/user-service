package Controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"user/app/Http/Request"
	"user/app/Models"
	"user/app/Repositories"
	"user/database"
)

func UserIndex(c *fiber.Ctx) error {
	users := []Models.User{}
	database.DBConn.Find(&users)

	res := c.JSON(users)
	return res
}

func UserDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	user := Models.User{}
	database.DBConn.First(&user, id)

	return c.JSON(user)
}

func UserPermissions(c *fiber.Ctx) error {
	id := c.Params("id")
	user := Models.User{}
	database.DBConn.First(&user, id)

	permissions := []Models.Permission{}
	database.DBConn.Model(&user).Association("Permissions").Find(&permissions)

	// Extract IDs of updated permissions
	var permissionIDs []uint
	for _, perm := range permissions {
		permissionIDs = append(permissionIDs, perm.ID)
	}

	res := c.JSON(permissionIDs)
	return res
}

func UserSavePermissions(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "USER-T0",
			"message": "Invalid user ID",
		})
	}

	request := new(Request.UserSavePermissionsRequest)
	_ = c.BodyParser(&request)
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

	user, err := Repositories.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    "USER-T1",
			"message": "User not found",
		})
	}

	// Retrieve current permissions
	var currentPermissions []Models.Permission
	userPermissionError := database.DBConn.Model(&user).Association("Permissions").Find(&currentPermissions)
	if userPermissionError != nil {
		return err
	}

	// Create a map of current permissions for easy lookup
	currentPermissionMap := make(map[uint]bool)
	for _, perm := range currentPermissions {
		currentPermissionMap[perm.ID] = true
	}

	// Create a map of incoming permissions for easy lookup
	incomingPermissionMap := make(map[uint]bool)
	for _, perm := range request.Permissions {
		incomingPermissionMap[uint(perm)] = true
	}

	// Determine permissions to add and remove
	var permissionsToAdd []Models.Permission
	var permissionsToRemove []Models.Permission

	for _, perm := range request.Permissions {
		if !currentPermissionMap[uint(perm)] {
			permissionsToAdd = append(permissionsToAdd, Models.Permission{ID: uint(perm)})
		}
	}

	for _, perm := range currentPermissions {
		if !incomingPermissionMap[perm.ID] {
			permissionsToRemove = append(permissionsToRemove, perm)
		}
	}

	// Add new permissions
	if len(permissionsToAdd) > 0 {
		database.DBConn.Model(&user).Association("Permissions").Append(&permissionsToAdd)
	}

	// Remove old permissions
	if len(permissionsToRemove) > 0 {
		database.DBConn.Model(&user).Association("Permissions").Delete(&permissionsToRemove)
	}

	// Retrieve updated permissions
	var updatedPermissions []Models.Permission
	database.DBConn.Model(&user).Association("Permissions").Find(&updatedPermissions)

	// Extract IDs of updated permissions
	var updatedPermissionIDs []uint
	for _, perm := range updatedPermissions {
		updatedPermissionIDs = append(updatedPermissionIDs, perm.ID)
	}

	return c.JSON(fiber.Map{
		"message":     "Permissions updated successfully",
		"permissions": updatedPermissionIDs,
	})
}
