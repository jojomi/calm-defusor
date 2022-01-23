package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane_color"
)

type SimpleWiresModule struct {
	allColors []ktane_color.Color
	colors    []ktane_color.Color
}

func (s *SimpleWiresModule) Name() string {
	return "Einfache Drähte"
}

func NewSimpleWiresModule() *SimpleWiresModule {
	return &SimpleWiresModule{
		allColors: []ktane_color.Color{
			ktane_color.ColorRed,
			ktane_color.ColorBlue,
			ktane_color.ColorWhite,
			ktane_color.ColorYellow,
			ktane_color.ColorBlack,
		},
	}
}

func (s *SimpleWiresModule) Solve() error {
	// ask for colors
	index := 0
	for {
		choices := s.allColors
		if index > 2 {
			choices = append(choices, ktane_color.ColorNoMore)
		}
		color, err := communication.ChooseOneColor(fmt.Sprintf("Farbe %d. Draht", index+1), choices)
		if err != nil {
			return err
		}

		if color == ktane_color.ColorNoMore {
			break
		}

		s.colors = append(s.colors, color)
		index++
		if index == 6 {
			break
		}
	}
	fmt.Println(s.colors)

	// evaluate and ask more
	switch len(s.colors) {
	case 3:
		s.handleThree()
	case 4:
		s.handleFour()
	case 5:
		s.handleFive()
	case 6:
		s.handleSix()
	}
	return nil
}

func (s SimpleWiresModule) handleThree() {
	if s.countColor(ktane_color.ColorRed) == 0 {
		s.cut(2)
		return
	}
	if s.colors[2] == ktane_color.ColorWhite {
		s.cut(3)
	}
	if s.countColor(ktane_color.ColorBlue) > 1 {
		s.cut(s.getLastIndex(ktane_color.ColorBlue) + 1)
		return
	}
	s.cut(3)
}

func (s SimpleWiresModule) handleFour() {
	if s.countColor(ktane_color.ColorRed) > 1 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(s.getLastIndex(ktane_color.ColorRed) + 1)
			return
		}
	}

	if s.colors[3] == ktane_color.ColorYellow && s.countColor(ktane_color.ColorRed) == 0 {
		s.cut(1)
		return
	}

	if s.countColor(ktane_color.ColorBlue) == 1 {
		s.cut(1)
		return
	}

	if s.countColor(ktane_color.ColorYellow) > 1 {
		s.cut(4)
		return
	}

	s.cut(2)
	return
}

func (s SimpleWiresModule) handleFive() {
	if s.colors[4] == ktane_color.ColorBlack {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(4)
			return
		}
	}

	if s.countColor(ktane_color.ColorRed) == 1 && s.countColor(ktane_color.ColorYellow) > 1 {
		s.cut(1)
		return
	}

	if s.countColor(ktane_color.ColorBlack) == 0 {
		s.cut(2)
		return
	}

	s.cut(1)
	return
}

func (s SimpleWiresModule) handleSix() {
	if s.countColor(ktane_color.ColorYellow) == 0 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(3)
			return
		}
	}

	if s.countColor(ktane_color.ColorYellow) == 1 && s.countColor(ktane_color.ColorWhite) > 1 {
		s.cut(4)
		return
	}

	if s.countColor(ktane_color.ColorRed) == 0 {
		s.cut(6)
		return
	}

	s.cut(4)
	return
}

// one-based!
func (s SimpleWiresModule) cut(index int) {
	communication.Tellf("%d. Draht trennen (dieser hat die Farbe %s)", index, s.colors[index-1].BySysLocale())
}

func (s SimpleWiresModule) countColor(searchColor ktane_color.Color) int {
	result := 0
	for _, color := range s.colors {
		if color != searchColor {
			continue
		}
		result++
	}
	return result
}

func (s SimpleWiresModule) getLastIndex(searchColor ktane_color.Color) int {
	result := 0
	for i, color := range s.colors {
		if color != searchColor {
			continue
		}
		result = i
	}
	return result
}
