package main

import (
	"image/png"
	"os"
	"rental-team-maker/utils"
)

func main() {
	// Test data
	pms := []utils.Pokemon{
		{
			Name:    "Dragapult",
			Type:    []string{"Dragon", "Ghost"},
			Item:    "Weakness Policy",
			Ability: "Clear Body",
			Moves:   nil,
		},
		{
			Name:    "Arcanine",
			Type:    []string{"Fire"},
			Item:    "Assault Vest",
			Ability: "Intimidate",
			Moves:   nil,
		},
		{
			Name:    "Rotom",
			Type:    []string{"Electric", "Water"},
			Item:    "Sitrus Berry",
			Ability: "Levitate",
			Moves:   nil,
		},
		{
			Name:    "Tyranitar",
			Type:    []string{"Rock", "Dark"},
			Item:    "Focus Sash",
			Ability: "Sand Stream",
			Moves:   nil,
		},
		{
			Name:    "foo",
			Type:    []string{"Dragon", "Ghost"},
			Item:    "bar",
			Ability: "foobar",
			Moves:   nil,
		},
		{
			Name:    "Togekiss",
			Type:    []string{"Fairy", "Flying"},
			Item:    "Scope Lens",
			Ability: "Super Luck",
			Moves:   nil,
		},
	}

	// Read image from file that already exists
	bgImageFile, err := os.Open("sprites/bg.png")
	if err != nil {
		panic(err)
	}
	defer bgImageFile.Close()

	// Since we know it is a png already, call png.Decode()
	bg, err := png.Decode(bgImageFile)
	if err != nil {
		panic(err)
	}

	bg, err = utils.AppendPokemon(bg, &pms)
	if err != nil {
		panic(err)
	}

	bg, err = utils.AppendItems(bg, &pms)
	if err != nil {
		panic(err)
	}

	bg, err = utils.AppendInfo(bg, &pms)

	// Save an image
	err = utils.SaveImage(bg, "outimage.png")
	if err != nil {
		panic(err)
	}
}
