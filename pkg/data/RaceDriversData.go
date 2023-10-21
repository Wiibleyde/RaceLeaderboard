package data

import (
	"logger"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func minimizePhoneNumber(stringToTest string) int {
	var toReturn int
	var err error
	regExpFull := regexp.MustCompile(`^555-[0-9]{4}$`)
	regExpMidFull := regexp.MustCompile(`^555[0-9]{4}$`)
	regExp := regexp.MustCompile(`^[0-9]{4}$`)
	if regExp.MatchString(stringToTest) {
		toReturn, err = strconv.Atoi(stringToTest)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		return toReturn
	} else if regExpMidFull.MatchString(stringToTest) {
		toReturn, err = strconv.Atoi(stringToTest[3:])
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		return toReturn
	} else if regExpFull.MatchString(stringToTest) {
		toReturn, err = strconv.Atoi(stringToTest[4:])
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		return toReturn
	} else {
		return 0
	}
}

func InsertRaceDriver(DriverFName string, DriverLName string, DriverNumber string) {
	InitDatabase()
	defer CloseDatabase()

	minizedDriverNumber := minimizePhoneNumber(DriverNumber)

	_, err := db.Exec("INSERT INTO RaceDrivers (DriverFName, DriverLName, DriverNumber) VALUES (?, ?, ?)", DriverFName, DriverLName, minizedDriverNumber)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func GetRaceDriver(Id int) RaceDrivers {
	InitDatabase()
	defer CloseDatabase()

	var raceDriver RaceDrivers

	err := db.QueryRow("SELECT * FROM RaceDrivers WHERE Id = ?", Id).Scan(&raceDriver.Id, &raceDriver.DriverFName, &raceDriver.DriverLName, &raceDriver.DriverNumber)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return raceDriver
}
