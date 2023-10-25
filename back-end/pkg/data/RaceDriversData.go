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

	println(testIfRaceDriverExist(DriverFName, DriverLName, minizedDriverNumber))

	if testIfRaceDriverExist(DriverFName, DriverLName, minizedDriverNumber) {
		return
	}

	_, err := db.Exec("INSERT INTO RaceDrivers (DriverFName, DriverLName, DriverNumber) VALUES (?, ?, ?)", DriverFName, DriverLName, minizedDriverNumber)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func testIfRaceDriverExist(DriverFName string, DriverLName string, DriverNumber int) bool {
	var raceDriver RaceDrivers

	err := db.QueryRow("SELECT * FROM RaceDrivers WHERE DriverFName = ? AND DriverLName = ? AND DriverNumber = ?", DriverFName, DriverLName, DriverNumber).Scan(&raceDriver.Id, &raceDriver.DriverFName, &raceDriver.DriverLName, &raceDriver.DriverNumber)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	if raceDriver.Id == 0 {
		return false
	} else {
		return true
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

func DeleteRaceDriver(Id int) {
	InitDatabase()
	defer CloseDatabase()

	_, err := db.Exec("DELETE FROM RaceDrivers WHERE Id = ?", Id)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func GetRaceDrivers() []RaceDrivers {
	InitDatabase()
	defer CloseDatabase()

	var raceDrivers []RaceDrivers

	rows, err := db.Query("SELECT * FROM RaceDrivers")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	for rows.Next() {
		var raceDriver RaceDrivers
		err := rows.Scan(&raceDriver.Id, &raceDriver.DriverFName, &raceDriver.DriverLName, &raceDriver.DriverNumber)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		raceDrivers = append(raceDrivers, raceDriver)
	}

	return raceDrivers
}

func GetRaceDriverByNumber(DriverNumber string) RaceDrivers {
	InitDatabase()
	defer CloseDatabase()

	var raceDriver RaceDrivers

	minizedDriverNumber := minimizePhoneNumber(DriverNumber)

	err := db.QueryRow("SELECT * FROM RaceDrivers WHERE DriverNumber = ?", minizedDriverNumber).Scan(&raceDriver.Id, &raceDriver.DriverFName, &raceDriver.DriverLName, &raceDriver.DriverNumber)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return raceDriver
}

func UpdateRaceDriver(Id int, DriverFName string, DriverLName string, DriverNumber string) {
	InitDatabase()
	defer CloseDatabase()

	minizedDriverNumber := minimizePhoneNumber(DriverNumber)

	_, err := db.Exec("UPDATE RaceDrivers SET DriverFName = ?, DriverLName = ?, DriverNumber = ? WHERE Id = ?", DriverFName, DriverLName, minizedDriverNumber, Id)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}
