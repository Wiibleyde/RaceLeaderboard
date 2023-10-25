package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitApi() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// API status
	app.Get("/", homeApi)
	app.Get("/status", statusApi)

	app.Post("/createRace", createRaceApi)
	app.Post("/deleteRace", deleteRaceApi)
	app.Get("/listRaces", listRaceApi)

	app.Get("/getDriver", getDriverApi)
	app.Post("/addDriver", addDriverApi)
	app.Post("/deleteDriver", deleteDriverApi)
	app.Post("/updateDriver", updateDriverApi)
	app.Get("/listDriver", listDriverApi)
	
	app.Listen(":3001")
}
