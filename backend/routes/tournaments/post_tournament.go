package tournaments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxkruse/magnusopus/backend/globals"
	"github.com/maxkruse/magnusopus/backend/structs"
	"github.com/maxkruse/magnusopus/backend/utils"
)

func PostTournament(c *fiber.Ctx) error {
	if err := utils.IsSuperadmin(c); err != nil {
		return utils.DefaultErrorMessage(c, err, fiber.StatusUnauthorized)
	}
	c.Accepts("application/json")
	t := structs.Tournament{}

	// Decode body
	err := c.BodyParser(&t)

	if err != nil {
		return utils.DefaultErrorMessage(c, err, fiber.StatusBadRequest)
	}

	localDB := globals.DBConn

	// TODO: Validate tournament
	err = t.ValidTournament(localDB)
	if err != nil {
		return utils.DefaultErrorMessage(c, err, fiber.StatusUnprocessableEntity)
	}

	// Dont Accept: Rounds, Staff. Instead, set them to nil
	t.Rounds = nil
	t.Staffs = nil

	self, err := utils.GetSelfFromDB(c)
	if err != nil {
		return utils.DefaultErrorMessage(c, err, fiber.StatusInternalServerError)
	}

	staffMember := structs.Staff{
		User: &self,
		Role: "owner",
	}
	t.Staffs = append(t.Staffs, staffMember)

	err = localDB.Save(&t).Error

	if err != nil {
		return utils.DefaultErrorMessage(c, err, fiber.StatusInternalServerError)
	}

	// Remove session from staff user to not display
	for _, staff := range t.Staffs {
		staff.User.Sessions = nil
	}

	return c.Status(fiber.StatusOK).JSON(t)
}
