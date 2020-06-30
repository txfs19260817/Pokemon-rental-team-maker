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
			Moves:   []string{"Dragon Darts", "Phantom Force", "Fly", "Draco Meteor"},
		},
		{
			Name:    "Arcanine",
			Type:    []string{"Fire"},
			Item:    "Assault Vest",
			Ability: "Intimidate",
			Moves:   []string{"Dragon Darts", "Phantom Force", "Fly", "Protect"},
		},
		{
			Name:    "Rotom",
			Type:    []string{"Electric", "Water"},
			Item:    "Sitrus Berry",
			Ability: "Levitate",
			Moves:   []string{"Thunderbolt", "Phantom Force", "Fly", "Protect"},
		},
		{
			Name:    "Tyranitar",
			Type:    []string{"Rock", "Dark"},
			Item:    "Focus Sash",
			Ability: "Sand Stream",
			Moves:   []string{"Crunch", "Phantom Force", "Fly", "Protect"},
		},
		{
			Name:    "foo",
			Type:    []string{"Dragon", "Ghost"},
			Item:    "bar",
			Ability: "foobar",
			Moves:   []string{"foo", "Phantom Force", "Fly", "Protect"},
		},
		{
			Name:    "Togekiss",
			Type:    []string{"Fairy", "Flying"},
			Item:    "Scope Lens",
			Ability: "Super Luck",
			Moves:   []string{"Heat Wave", "Phantom Force", "Fly", "Protect"},
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
	if err != nil {
		panic(err)
	}

	bg, err = utils.AppendTypes(bg, &pms)
	if err != nil {
		panic(err)
	}

	bg, err = utils.AppendMoves(bg, &pms)
	if err != nil {
		panic(err)
	}

	// Save an image
	err = utils.SaveImage(bg, "outimage.png")
	if err != nil {
		panic(err)
	}
}
