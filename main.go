package main

import (
	"bufio"
	"github.com/sugoiuguu/showgone"
	"image/png"
	"os"
	"rental-team-maker/utils"
	"strings"
)

func main() {
	pms := []utils.Pokemon{}
	r := bufio.NewReader(strings.NewReader("Rillaboom @ Leftovers  \nAbility: Grassy Surge  \nLevel: 50  \nEVs: 252 HP / 60 Atk / 92 Def / 92 SpD / 12 Spe  \nAdamant Nature  \n- Fake Out  \n- Knock Off  \n- Protect  \n- Grassy Glide  \n\nComfey @ Babiri Berry  \nAbility: Triage  \nLevel: 50  \nEVs: 252 HP / 252 Def / 4 SpD  \nBold Nature  \nIVs: 0 Atk  \n- Ally Switch  \n- Floral Healing  \n- Protect  \n- Giga Drain  \n\nTogekiss @ Safety Goggles  \nAbility: Serene Grace  \nLevel: 50  \nEVs: 188 HP / 4 Def / 76 SpA / 148 SpD / 92 Spe  \nModest Nature  \nIVs: 0 Atk  \n- Dazzling Gleam  \n- Protect  \n- Air Slash  \n- Follow Me  \n\nIncineroar @ Aguav Berry  \nAbility: Intimidate  \nEVs: 252 HP / 4 Atk / 100 Def / 124 SpD / 28 Spe  \nCareful Nature  \n- Flare Blitz  \n- Fake Out  \n- Parting Shot  \n- Darkest Lariat  \n\nDracozolt @ Life Orb  \nAbility: Hustle  \nLevel: 50  \nEVs: 252 Atk / 4 SpD / 252 Spe  \nJolly Nature  \n- Bolt Beak  \n- Dragon Claw  \n- Protect  \n- Aerial Ace  \n\nLapras-Gmax @ Weakness Policy  \nAbility: Hydration  \nLevel: 50  \nEVs: 20 HP / 236 Def / 196 SpA / 4 SpD / 52 Spe  \nModest Nature  \nIVs: 0 Atk  \n- Hydro Pump  \n- Freeze-Dry  \n- Protect  \n- Perish Song  \n\n"))
	for {
		poke, err := showgone.Parse(r)
		if err != nil {
			break
		}
		var tp []string
		if t, ok := utils.Poke2Types[utils.String2Filename(string(poke.Species))]; ok {
			tp = t
		}
		pm := utils.Pokemon{
			Name:    string(poke.Species),
			Type:    tp,
			Item:    string(poke.Item),
			Ability: string(poke.Ability),
			Moves:   []string{string(poke.Moves[0]), string(poke.Moves[1]), string(poke.Moves[2]), string(poke.Moves[3])},
		}
		pms = append(pms, pm)
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
