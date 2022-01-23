package modules

import (
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/go-script/v2/interview"
	"strings"
)

type SimonSaysModule struct {
	allColors []string
}

func (s *SimonSaysModule) Name() string {
	return "Simon sagt"
}

func NewSimonSaysModule() *SimonSaysModule {
	allColors := []string{
		"blue",
		"yellow",
		"green",
		"red",
	}
	return &SimonSaysModule{
		allColors: allColors,
	}
}

func (s *SimonSaysModule) Solve() error {
	strikeCount := 0
	strikes, err := interview.ChooseOneString("Bisherige Fehler (Strikes)?", []string{
		"Keine Fehler",
		"1 Fehler",
		"2 Fehler",
	})
	if err != nil {
		return err
	}
	switch strikes {
	case "1 Fehler":
		strikeCount = 1
	case "2 Fehler":
		strikeCount = 2
	}

	hasVocal, err := interview.ConfirmNoDefault("Enthält die Seriennummer einen Vokal?")
	if err != nil {
		return err
	}

	colorChoices := append(s.allColors, "ENDE")
	var (
		colors       []string
		mappedColors []string
		mappedColor  string
	)
	for {
		newColor, err := interview.ChooseOneString("Neue/letzte Farbe?", colorChoices)
		if err != nil {
			return err
		}
		if newColor == "ENDE" {
			break
		}
		colors = append(colors, newColor)

		mappedColor = s.getMappedColor(strikeCount, hasVocal, newColor)
		mappedColors = append(mappedColors, mappedColor)

		communication.Tellf("Folgende Farben nacheinander drücken: %s\n", strings.Join(mappedColors, ", "))
	}

	return nil
}

func (s *SimonSaysModule) getMappedColor(strikeCount int, hasVocal bool, color string) string {
	if hasVocal {
		return s.getVocalMappedColor(strikeCount, color)
	}
	return s.getNonVocalMappedColor(strikeCount, color)
}

func (s *SimonSaysModule) getVocalMappedColor(strikeCount int, color string) string {
	colorMap := map[string]string{}
	switch strikeCount {
	case 0:
		colorMap = map[string]string{
			"red":    "blue",
			"blue":   "red",
			"green":  "yellow",
			"yellow": "green",
		}
	case 1:
		colorMap = map[string]string{
			"red":    "yellow",
			"blue":   "green",
			"green":  "blue",
			"yellow": "red",
		}
	case 2:
		colorMap = map[string]string{
			"red":    "green",
			"blue":   "red",
			"green":  "yellow",
			"yellow": "blue",
		}
	}
	return colorMap[color]
}

func (s *SimonSaysModule) getNonVocalMappedColor(strikeCount int, color string) string {
	colorMap := map[string]string{}
	switch strikeCount {
	case 0:
		colorMap = map[string]string{
			"red":    "blue",
			"blue":   "yellow",
			"green":  "green",
			"yellow": "red",
		}
	case 1:
		colorMap = map[string]string{
			"red":    "red",
			"blue":   "blue",
			"green":  "yellow",
			"yellow": "green",
		}
	case 2:
		colorMap = map[string]string{
			"red":    "yellow",
			"blue":   "green",
			"green":  "blue",
			"yellow": "red",
		}
	}
	return colorMap[color]
}
