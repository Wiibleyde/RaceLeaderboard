package data

import (
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func InsertRace(Code string, Name string) {
	InitDatabase()
	defer CloseDatabase()

	_, err := db.Exec("INSERT INTO Race (Code, Name) VALUES (?, ?)", Code, Name)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func DeleteRace(Id int) {
	InitDatabase()
	defer CloseDatabase()

	_, err := db.Exec("DELETE FROM Race WHERE Id = ?", Id)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
}

func GetRace(Id int) Race {
	InitDatabase()
	defer CloseDatabase()

	var race Race

	err := db.QueryRow("SELECT * FROM Race WHERE Id = ?", Id).Scan(&race.Id, &race.Code, &race.Name)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return race
}

func GetRaces() []Race {
	InitDatabase()
	defer CloseDatabase()

	var races []Race

	rows, err := db.Query("SELECT * FROM Race")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	for rows.Next() {
		var race Race
		err := rows.Scan(&race.Id, &race.Code, &race.Name)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		races = append(races, race)
	}

	return races
}
