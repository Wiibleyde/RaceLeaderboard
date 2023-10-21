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
