package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane_color"
	"github.com/jojomi/go-script/v2/interview"
	"strings"
)

type SimonSaysModule struct {
	allColors []ktane_color.Color
}

func (s *SimonSaysModule) Name() string {
	return "Simon sagt"
}

func NewSimonSaysModule() *SimonSaysModule {
	allColors := []ktane_color.Color{
		ktane_color.ColorBlue,
		ktane_color.ColorYellow,
		ktane_color.ColorGreen,
		ktane_color.ColorRed,
	}
	return &SimonSaysModule{
		allColors: allColors,
	}
}

func (s *SimonSaysModule) Solve() error {
	strikeCount, err := communication.ChooseOneWithMapper[int]("Bisherige Fehler (Strikes)?", []int{0, 1, 2}, func(t int) string {
		if t == 0 {
			return "Keine Fehler"
		}
		return fmt.Sprintf("%d Fehler", t)
	})
	fmt.Println(strikeCount)
	if err != nil {
		return err
	}

	hasVocal, err := interview.ConfirmNoDefault("Enthält die Seriennummer einen Vokal?")
	if err != nil {
		return err
	}

	colorChoices := append(s.allColors, ktane_color.ColorNoMore)
	var (
		colors       []ktane_color.Color
		mappedColors []ktane_color.Color
		mappedColor  ktane_color.Color
	)
	for {
		newColor, err := communication.ChooseOneColor("Neue/letzte Farbe?", colorChoices)
		if err != nil {
			return err
		}
		if newColor == ktane_color.ColorNoMore {
			break
		}
		colors = append(colors, newColor)

		mappedColor = s.getMappedColor(strikeCount, hasVocal, newColor)
		mappedColors = append(mappedColors, mappedColor)

		communication.Tellf("Folgende Farben nacheinander drücken: %s\n", strings.Join(
			mapList[ktane_color.Color, string](
				mappedColors,
				func(c ktane_color.Color) string {
					return c.BySysLocale()
				},
			), ", "))
	}

	return nil
}

func (s *SimonSaysModule) getMappedColor(strikeCount int, hasVocal bool, color ktane_color.Color) ktane_color.Color {
	if hasVocal {
		return s.getVocalMappedColor(strikeCount, color)
	}
	return s.getNonVocalMappedColor(strikeCount, color)
}

func (s *SimonSaysModule) getVocalMappedColor(strikeCount int, color ktane_color.Color) ktane_color.Color {
	colorMap := map[ktane_color.Color]ktane_color.Color{}
	switch strikeCount {
	case 0:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorBlue,
			ktane_color.ColorBlue:   ktane_color.ColorRed,
			ktane_color.ColorGreen:  ktane_color.ColorYellow,
			ktane_color.ColorYellow: ktane_color.ColorGreen,
		}
	case 1:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorYellow,
			ktane_color.ColorBlue:   ktane_color.ColorGreen,
			ktane_color.ColorGreen:  ktane_color.ColorBlue,
			ktane_color.ColorYellow: ktane_color.ColorRed,
		}
	case 2:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorGreen,
			ktane_color.ColorBlue:   ktane_color.ColorRed,
			ktane_color.ColorGreen:  ktane_color.ColorYellow,
			ktane_color.ColorYellow: ktane_color.ColorBlue,
		}
	}
	return colorMap[color]
}

func (s *SimonSaysModule) getNonVocalMappedColor(strikeCount int, color ktane_color.Color) ktane_color.Color {
	colorMap := map[ktane_color.Color]ktane_color.Color{}
	switch strikeCount {
	case 0:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorBlue,
			ktane_color.ColorBlue:   ktane_color.ColorYellow,
			ktane_color.ColorGreen:  ktane_color.ColorGreen,
			ktane_color.ColorYellow: ktane_color.ColorRed,
		}
	case 1:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorRed,
			ktane_color.ColorBlue:   ktane_color.ColorBlue,
			ktane_color.ColorGreen:  ktane_color.ColorYellow,
			ktane_color.ColorYellow: ktane_color.ColorGreen,
		}
	case 2:
		colorMap = map[ktane_color.Color]ktane_color.Color{
			ktane_color.ColorRed:    ktane_color.ColorYellow,
			ktane_color.ColorBlue:   ktane_color.ColorGreen,
			ktane_color.ColorGreen:  ktane_color.ColorBlue,
			ktane_color.ColorYellow: ktane_color.ColorRed,
		}
	}
	return colorMap[color]
}
