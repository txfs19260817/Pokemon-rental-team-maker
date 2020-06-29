package utils

import (
	"regexp"
	"strings"
)

type Pokemon struct {
	Name string
	Type []string
	Item string
	Ability string
	Moves []string
}

var (
	// Paths
	SpritePath = "sprites/"

	// Offsets
	OffsetX = 588
	OffsetY = 186
)

func String2Filename(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return strings.ToLower(reg.ReplaceAllString(s, "${1}"))
}