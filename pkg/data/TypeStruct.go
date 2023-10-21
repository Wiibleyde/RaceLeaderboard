package data

type Race struct {
	Id   int
	Code string
	Name string
}

type RaceLiveLeaderboard struct {
	Id       int
	RaceId   int
	Position int
	DriverId string
}

type RaceDrivers struct {
	Id           int
	DriverFName  string
	DriverLName  string
	DriverNumber string
}

type RaceDriverResults struct {
	Id               int
	RaceId           int
	DriversIdInOrder string
}
