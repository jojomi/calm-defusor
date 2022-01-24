package ktane

import (
	"github.com/Xuanwo/go-locale"
	"github.com/gookit/color"
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

func (x Color) BySysLocaleForTerminal() string {
	value := x.BySysLocale()

	colorPrinter := color.New(color.FgBlack, color.BgWhite, color.Bold)
	switch x {
	case ColorRed:
		colorPrinter = color.New(color.FgRed, color.BgWhite, color.Bold)
	case ColorGreen:
		colorPrinter = color.New(color.FgGreen, color.BgWhite, color.Bold)
	case ColorYellow:
		colorPrinter = color.New(color.FgYellow, color.BgWhite, color.Bold)
	case ColorBlue:
		colorPrinter = color.New(color.FgBlue, color.BgWhite, color.Bold)
	case ColorBlack:
		colorPrinter = color.New(color.FgBlack, color.BgWhite, color.Bold)
	case ColorWhite:
		colorPrinter = color.New(color.FgWhite, color.BgBlack, color.Bold)
	}

	return colorPrinter.Render(" " + value + " ")
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
