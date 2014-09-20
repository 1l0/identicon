package main

import (
	"github.com/1l0/identicon"
)

func main() {

	id := identicon.New()

	// Default
	id.GeneratePNGToFile("identicons/default")

	// Type: Vertical, Theme: Black
	id.Type = identicon.Mirrorv
	id.Theme = identicon.Black
	id.GeneratePNGToFile("identicons/vertical_black")

	// Divisions: 7, Theme: Gray
	id.Type = identicon.Mirrorh
	id.Theme = identicon.Gray
	id.Q = 50
	id.Div = 7
	id.GeneratePNGToFile("identicons/div7_gray")

	// Margin: 140, Theme: Free
	id.Theme = identicon.Free
	id.Q = 70
	id.Div = 5
	id.Margin = 140
	id.GeneratePNGToFile("identicons/margin140_free")

	// Type: Normal, Theme: White
	id.Type = identicon.Normal
	id.Theme = identicon.White
	id.Margin = 35
	id.GeneratePNGToFile("identicons/normal_white")

	id = identicon.New()

	// Random batch
	id.GenerateRandomThemes("identicons/rand", 4)

	// Sequential batch
	id.GenerateSequentialThemes("identicons/seq", 1)
}
