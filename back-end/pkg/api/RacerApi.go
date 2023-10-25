package api

import (
	"data"
	"logger"
	"strconv"

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
		logger.ErrorLogger.Println(err)
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
		logger.ErrorLogger.Println(err)
		return err
	}

	data.DeleteRaceDriver(deleteDriverApi.DriverId)

	return c.SendString("OK")
}

type updateDriverApiStruct struct {
	DriverId     string
	DriverFName  string
	DriverLName  string
	DriverNumber string
}

func updateDriverApi(c *fiber.Ctx) error {
	var updateDriverApi updateDriverApiStruct
	if err := c.BodyParser(&updateDriverApi); err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	id, err := strconv.Atoi(updateDriverApi.DriverId)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	data.UpdateRaceDriver(id, updateDriverApi.DriverFName, updateDriverApi.DriverLName, updateDriverApi.DriverNumber)

	return c.SendString("OK")
}

func listDriverApi(c *fiber.Ctx) error {
	return c.JSON(data.GetRaceDrivers())
}

func getDriverApi(c *fiber.Ctx) error {
	var id int
	var err error

	id, err = strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	return c.JSON(data.GetRaceDriver(id))
}
