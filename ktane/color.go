package ktane

import (
	"github.com/Xuanwo/go-locale"
	"golang.org/x/text/language"
)

//go:generate go-enum -f=$GOFILE --marshal

// Color is an enumeration of color values
/*
ENUM(
white
black
blue
red
yellow
green
noMore
)
*/
type Color int

func (x Color) BySysLocale() string {
	languageTag, err := locale.Detect()
	if err != nil {
		return x.String()
	}
	return x.Localized(languageTag)
}

func (x Color) Localized(languageTag language.Tag) string {
	matcher := language.NewMatcher([]language.Tag{
		language.German,
	})

	bestTag, _, _ := matcher.Match(
		languageTag,
	)
	l, _ := bestTag.Base()

	german, _ := language.German.Base()
	colorNameMap := map[language.Base]map[Color]string{
		german: {
			ColorWhite:  "weiß",
			ColorBlack:  "schwarz",
			ColorBlue:   "blau",
			ColorRed:    "rot",
			ColorYellow: "gelb",
			ColorGreen:  "grün",
			ColorNoMore: "Keine weitere Farbe",
		},
	}

	name, ok := colorNameMap[l][x]
	if !ok {
		return x.String()
	}

	return name
}
