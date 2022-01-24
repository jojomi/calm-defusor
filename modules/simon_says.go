package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane"
	"github.com/jojomi/go-script/v2/interview"
	"strings"
)

type SimonSaysModule struct {
	allColors []ktane.Color
}

func (s *SimonSaysModule) Name() string {
	return "Simon sagt"
}

func NewSimonSaysModule() *SimonSaysModule {
	allColors := []ktane.Color{
		ktane.ColorBlue,
		ktane.ColorYellow,
		ktane.ColorGreen,
		ktane.ColorRed,
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

	colorChoices := append(s.allColors, ktane.ColorNoMore)
	var (
		colors       []ktane.Color
		mappedColors []ktane.Color
		mappedColor  ktane.Color
	)
	for {
		newColor, err := communication.ChooseOneColor("Neue/letzte Farbe?", colorChoices)
		if err != nil {
			return err
		}
		if newColor == ktane.ColorNoMore {
			break
		}
		colors = append(colors, newColor)

		mappedColor = s.getMappedColor(strikeCount, hasVocal, newColor)
		mappedColors = append(mappedColors, mappedColor)

		communication.Tellf("Folgende Farben nacheinander drücken: %s\n", strings.Join(
			mapList[ktane.Color, string](
				mappedColors,
				func(c ktane.Color) string {
					return c.BySysLocaleForTerminal()
				},
			), ", "))
	}

	return nil
}

func (s *SimonSaysModule) getMappedColor(strikeCount int, hasVocal bool, color ktane.Color) ktane.Color {
	if hasVocal {
		return s.getVocalMappedColor(strikeCount, color)
	}
	return s.getNonVocalMappedColor(strikeCount, color)
}

func (s *SimonSaysModule) getVocalMappedColor(strikeCount int, color ktane.Color) ktane.Color {
	colorMap := map[ktane.Color]ktane.Color{}
	switch strikeCount {
	case 0:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorBlue,
			ktane.ColorBlue:   ktane.ColorRed,
			ktane.ColorGreen:  ktane.ColorYellow,
			ktane.ColorYellow: ktane.ColorGreen,
		}
	case 1:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorYellow,
			ktane.ColorBlue:   ktane.ColorGreen,
			ktane.ColorGreen:  ktane.ColorBlue,
			ktane.ColorYellow: ktane.ColorRed,
		}
	case 2:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorGreen,
			ktane.ColorBlue:   ktane.ColorRed,
			ktane.ColorGreen:  ktane.ColorYellow,
			ktane.ColorYellow: ktane.ColorBlue,
		}
	}
	return colorMap[color]
}

func (s *SimonSaysModule) getNonVocalMappedColor(strikeCount int, color ktane.Color) ktane.Color {
	colorMap := map[ktane.Color]ktane.Color{}
	switch strikeCount {
	case 0:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorBlue,
			ktane.ColorBlue:   ktane.ColorYellow,
			ktane.ColorGreen:  ktane.ColorGreen,
			ktane.ColorYellow: ktane.ColorRed,
		}
	case 1:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorRed,
			ktane.ColorBlue:   ktane.ColorBlue,
			ktane.ColorGreen:  ktane.ColorYellow,
			ktane.ColorYellow: ktane.ColorGreen,
		}
	case 2:
		colorMap = map[ktane.Color]ktane.Color{
			ktane.ColorRed:    ktane.ColorYellow,
			ktane.ColorBlue:   ktane.ColorGreen,
			ktane.ColorGreen:  ktane.ColorBlue,
			ktane.ColorYellow: ktane.ColorRed,
		}
	}
	return colorMap[color]
}
