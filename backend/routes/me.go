package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxkruse/magnusopus/backend/util"
)

func Me(c *fiber.Ctx) error {
	self, err := util.GetSelf(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(self)
}