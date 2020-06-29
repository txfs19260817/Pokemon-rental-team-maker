package utils

import (
	"errors"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"io/ioutil"
)

var (
	// Path
	FontPath = SpritePath + "Lato-Bold.ttf"

	// Size
	FontSize = 20.0

	// Point sets
	// 3. Info
	PointsInfoX = 88
	PointsInfoY = 114
	PointsInfo = []image.Point{
		{PointsInfoX, PointsInfoY}, {PointsInfoX + OffsetX, PointsInfoY},
		{PointsInfoX, PointsInfoY + OffsetY}, {PointsInfoX + OffsetX, PointsInfoY + OffsetY},
		{PointsInfoX, PointsInfoY + OffsetY*2}, {PointsInfoX + OffsetX, PointsInfoY + OffsetY*2},
	}
)

// Append each pokemon name, ability and item text on the left part of each slot.
func AppendInfo(canvas image.Image, pokemonList *[]Pokemon) (image.Image, error) {
	// Check length of pokemon list
	if len(*pokemonList) <= 0 || len(*pokemonList) > 6 {
		return nil, errors.New("Length of `pokemonList` must between 1 and 6. ")
	}

	// Info line space
	var lineSpace = 34
	
	// Load font file
	fontBytes, err := ioutil.ReadFile(FontPath)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	// Generate output image
	b := canvas.Bounds()
	out := image.NewRGBA(b)
	draw.Draw(out, out.Bounds(), canvas, b.Min, draw.Src)

	// Set FreeType context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(FontSize)
	c.SetClip(canvas.Bounds())
	c.SetDst(out)
	c.SetSrc(image.White)

	for i := range *pokemonList {
		pt := freetype.Pt(PointsInfo[i].X, PointsInfo[i].Y)
		if _, err = c.DrawString((*pokemonList)[i].Name, pt); err != nil {
			return nil, err
		}
		pt = freetype.Pt(PointsInfo[i].X, PointsInfo[i].Y + lineSpace)
		if _, err = c.DrawString((*pokemonList)[i].Ability, pt); err != nil {
			return nil, err
		}
		pt = freetype.Pt(PointsInfo[i].X, PointsInfo[i].Y + lineSpace * 2)
		if _, err = c.DrawString((*pokemonList)[i].Item, pt); err != nil {
			return nil, err
		}
	}

	return out, nil
}
