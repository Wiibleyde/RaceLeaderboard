package api

import (
	"data"

	"github.com/gofiber/fiber/v2"
)

type createRaceApiStruct struct {
	RaceCode string
	RaceName string
}

func createRaceApi(c *fiber.Ctx) error {
	var createRaceApi createRaceApiStruct
	if err := c.BodyParser(&createRaceApi); err != nil {
		return err
	}

	data.InsertRace(createRaceApi.RaceCode, createRaceApi.RaceName)

	return c.SendString("OK")
}

type deleteRaceApiStruct struct {
	RaceId int
}

func deleteRaceApi(c *fiber.Ctx) error {
	var deleteRaceApi deleteRaceApiStruct
	if err := c.BodyParser(&deleteRaceApi); err != nil {
		return err
	}

	data.DeleteRace(deleteRaceApi.RaceId)

	return c.SendString("OK")
}

func listRaceApi(c *fiber.Ctx) error {
	return c.JSON(data.GetRaces())
}