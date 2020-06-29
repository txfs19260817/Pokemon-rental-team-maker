package utils

import (
	"errors"
	"image"
	"image/draw"
	"image/png"
	"os"
)

var (
	// Point sets
	// 1. Pokemon
	PointsPmX = 95
	PointsPmY = 30
	PointsPm  = []image.Point{
		{PointsPmX, PointsPmY}, {PointsPmX + OffsetX, PointsPmY},
		{PointsPmX, PointsPmY + OffsetY}, {PointsPmX + OffsetX, PointsPmY + OffsetY},
		{PointsPmX, PointsPmY + OffsetY*2}, {PointsPmX + OffsetX, PointsPmY + OffsetY*2},
	}
	// 2. Items
	PointsItemX = 150
	PointsItemY = 50
	PointsItem = []image.Point{
		{PointsItemX, PointsItemY}, {PointsItemX + OffsetX, PointsItemY},
		{PointsItemX, PointsItemY + OffsetY}, {PointsItemX + OffsetX, PointsItemY + OffsetY},
		{PointsItemX, PointsItemY + OffsetY*2}, {PointsItemX + OffsetX, PointsItemY + OffsetY*2},
	}
)

// Save image.Image object to path `dst`.
func SaveImage(loadedImage image.Image, dst string) error {
	// Save an image
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode to `PNG` with `BestSpeed` level, then save to file
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	err = enc.Encode(f, loadedImage)
	if err != nil {
		return err
	}
	return nil
}

// Core append method. Call to return a concatenated image.
func AppendImage(canvas image.Image, targetPath string, offset image.Point) (*image.RGBA, error) {
	// Open target image
	targetFile, err := os.Open(targetPath)
	if err != nil {
		return nil, err
	}
	defer targetFile.Close()
	target, err := png.Decode(targetFile)
	if err != nil {
		return nil, err
	}

	// Append target image on canvas
	b := canvas.Bounds()
	out := image.NewRGBA(b)
	draw.Draw(out, b, canvas, image.Point{}, draw.Src)
	draw.Draw(out, target.Bounds().Add(offset), target, image.Point{}, draw.Over)

	return out, nil
}

// Append all pokemon sprites.
func AppendPokemon(canvas image.Image, pokemonList *[]Pokemon) (image.Image, error) {
	// Check length of pokemon list
	if len(*pokemonList) <= 0 || len(*pokemonList) > 6 {
		return nil, errors.New("Length of `pokemonList` must between 1 and 6. ")
	}

	// Append loop
	for i := range *pokemonList {
		var err error
		targetPath := SpritePath + "2d/" + String2Filename((*pokemonList)[i].Name) + ".png"

		// When file not found error occurred, use unknown.png instead.
		if _, err := os.Stat(targetPath); err != nil {
			targetPath = SpritePath + "2d/unknown.png"
		}

		// Append pokemon sprites at specified points
		canvas, err = AppendImage(canvas, targetPath, PointsPm[i])
		if err != nil {
			return nil, err
		}
	}
	return canvas, nil
}

// Append all items sprites.
func AppendItems(canvas image.Image, pokemonList *[]Pokemon) (image.Image, error) {
	// Check length of pokemon list
	if len(*pokemonList) <= 0 || len(*pokemonList) > 6 {
		return nil, errors.New("Length of `pokemonList` must between 1 and 6. ")
	}

	// Append loop
	for i := range *pokemonList {
		var err error
		targetPath := SpritePath + "items/" + String2Filename((*pokemonList)[i].Item) + ".png"

		// When file not found error occurred, use unknown.png instead.
		if _, err := os.Stat(targetPath); err != nil {
			targetPath = SpritePath + "items/unknown.png"
		}

		// Append pokemon sprites at specified points
		canvas, err = AppendImage(canvas, targetPath, PointsItem[i])
		if err != nil {
			return nil, err
		}
	}
	return canvas, nil
}