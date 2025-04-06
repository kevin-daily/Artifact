package main

type Coordinates struct {
	X string
	Y string
}

func getCoordinates(first string, second string) Coordinates {
	var x, y string

	switch first {
	case "mining":
		x = "1"
		y = "5"
	case "woodcutting":
		x = "-2"
		y = "-3"
	case "cooking":
		x = "1"
		y = "1"
	case "weaponcrafting":
		x = "2"
		y = "1"
	case "gearcrafting":
		x = "3"
		y = "1"
	case "jewelrycrafting":
		x = "1"
		y = "3"
	case "alchemy":
		x = "2"
		y = "3"
	case "copper_rocks":
		x = "2"
		y = "0"
	case "ash_tree":
		x = "-1"
		y = "0"
	case "salmon":
		x = "-3"
		y = "-4"
	default:
		x = first
		y = second
	}

	return Coordinates{x, y}
}
