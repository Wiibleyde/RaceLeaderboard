package api

import (
	"data"

	"github.com/gofiber/fiber/v2"
)

type addDriverApiStruct struct {
	DriverFName  string
	DriverLName  string
	DriverNumber string
}

func addDriverApi(c *fiber.Ctx) error {
	var addDriverApi addDriverApiStruct
	if err := c.BodyParser(&addDriverApi); err != nil {
		return err
	}

	data.InsertRaceDriver(addDriverApi.DriverFName, addDriverApi.DriverLName, addDriverApi.DriverNumber)

	racerAdded := data.GetRaceDriverByNumber(addDriverApi.DriverNumber)
	if racerAdded.Id == 0 {
		return c.SendString("ERROR")
	}
	
	return c.SendString("OK")
}

type deleteDriverApiStruct struct {
	DriverId int
}

func deleteDriverApi(c *fiber.Ctx) error {
	var deleteDriverApi deleteDriverApiStruct
	if err := c.BodyParser(&deleteDriverApi); err != nil {
		return err
	}

	data.DeleteRaceDriver(deleteDriverApi.DriverId)

	return c.SendString("OK")
}

func listDriverApi(c *fiber.Ctx) error {
	return c.JSON(data.GetRaceDrivers())
}

