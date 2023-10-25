package data

import (
	"database/sql"
	"logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseHost     = "localhost"
	databasePort     = "3306"
	databaseUser     = "race"
	databasePassword = "race"
	databaseName     = "RaceAPI"
)

var db *sql.DB

func InitDatabase() {
	var err error
	// if env var STARTED_BY_DOCKER is set to true, then we are running in a docker container so we need to use the docker host
	if os.Getenv("STARTED_BY_DOCKER") == "true" {
		db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp(db)/"+databaseName)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
	} else {
		db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
	}
}

func GetDatabase() *sql.DB {
	return db
}

func CloseDatabase() {
	db.Close()
}

func InitTables() {
	InitDatabase()
	defer CloseDatabase()
	var err error

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Race (Id INT NOT NULL AUTO_INCREMENT, Code VARCHAR(255) NOT NULL, Name VARCHAR(255) NOT NULL, PRIMARY KEY (Id))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS RaceLiveLeaderboard (Id INT NOT NULL AUTO_INCREMENT, RaceId INT NOT NULL, Position INT NOT NULL, DriverId VARCHAR(255) NOT NULL, PRIMARY KEY (Id))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS RaceDrivers (Id INT NOT NULL AUTO_INCREMENT, DriverFName VARCHAR(255) NOT NULL, DriverLName VARCHAR(255) NOT NULL, DriverNumber VARCHAR(255) NOT NULL, PRIMARY KEY (Id))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS RaceDriverResults (Id INT NOT NULL AUTO_INCREMENT, RaceId INT NOT NULL, DriversIdInOrder VARCHAR(255) NOT NULL, PRIMARY KEY (Id))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err.Error())
	}
}
