package data

import (
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func InsertRacerInLeaderboard(RaceId int, Position int, DriverId int) {
	InitDatabase()
	defer CloseDatabase()

	_, err := db.Exec("INSERT INTO RaceLiveLeaderboard (RaceId, Position, DriverId) VALUES (?, ?, ?)", RaceId, Position, DriverId)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func DeleteRacerInLeaderboard(RaceId int, DriverId int) {
	InitDatabase()
	defer CloseDatabase()

	_, err := db.Exec("DELETE FROM RaceLiveLeaderboard WHERE RaceId = ? AND DriverId = ?", RaceId, DriverId)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func GetRaceLeaderboard(RaceId int) []int {
	InitDatabase()
	defer CloseDatabase()

	var leaderboard []int

	rows, err := db.Query("SELECT DriverId FROM RaceLiveLeaderboard WHERE RaceId = ? ORDER BY Position ASC", RaceId)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	for rows.Next() {
		var DriverId int
		err = rows.Scan(&DriverId)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		leaderboard = append(leaderboard, DriverId)
	}

	return leaderboard
}

func UpdateRacerPlace(RaceId int, DriverId int, Action bool) {
	// If Action is true, then the driver has moved up one place
	// If Action is false, then the driver has moved down one place
	// Only one can be at one place at a time, so we can just swap the two drivers
	InitDatabase()
	defer CloseDatabase()

	var DriverIdToSwapWith int

	if Action {
		// The driver has moved up one place, so we need to find the driver that is one place below
		DriverIdToSwapWith = GetRaceLeaderboard(RaceId)[GetRacerPlace(RaceId, DriverId)-2]
	} else {
		// The driver has moved down one place, so we need to find the driver that is one place above
		DriverIdToSwapWith = GetRaceLeaderboard(RaceId)[GetRacerPlace(RaceId, DriverId)]
	}

	// Swap the two drivers
	_, err := db.Exec("UPDATE RaceLiveLeaderboard SET DriverId = ? WHERE RaceId = ? AND DriverId = ?", DriverId, RaceId, DriverIdToSwapWith)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
	_, err = db.Exec("UPDATE RaceLiveLeaderboard SET DriverId = ? WHERE RaceId = ? AND DriverId = ?", DriverIdToSwapWith, RaceId, DriverId)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func GetRacerPlace(RaceId int, DriverId int) int {
	InitDatabase()
	defer CloseDatabase()

	var place int

	err := db.QueryRow("SELECT Position FROM RaceLiveLeaderboard WHERE RaceId = ? AND DriverId = ?", RaceId, DriverId).Scan(&place)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return place
}

func GetRacerId(RaceId int, Position int) int {
	InitDatabase()
	defer CloseDatabase()

	var DriverId int

	err := db.QueryRow("SELECT DriverId FROM RaceLiveLeaderboard WHERE RaceId = ? AND Position = ?", RaceId, Position).Scan(&DriverId)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return DriverId
}
