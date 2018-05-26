//Package models package
package models

// Game model
type Game struct {
	Title   string `json:"Title"`
	Year    int32  `json:"Year"`
	Console string `json:"Console"`
}

//GetGames gets a list of games
func GetGames() []Game {
	var list []Game

	list = append(list, Game{Title: "Kirby's Adventure", Year: 1991, Console: "NES"})
	list = append(list, Game{Title: "Metroid", Year: 1989, Console: "NES"})

	return list
}
