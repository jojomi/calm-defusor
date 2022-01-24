package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane"
)

type SimpleWiresModule struct {
	allColors []ktane.Color
	colors    []ktane.Color
}

func (s *SimpleWiresModule) Name() string {
	return "Einfache Drähte"
}

func NewSimpleWiresModule() *SimpleWiresModule {
	return &SimpleWiresModule{
		allColors: []ktane.Color{
			ktane.ColorRed,
			ktane.ColorBlue,
			ktane.ColorWhite,
			ktane.ColorYellow,
			ktane.ColorBlack,
		},
	}
}

func (s *SimpleWiresModule) Solve() error {
	// ask for colors
	index := 0
	for {
		choices := s.allColors
		if index > 2 {
			choices = append(choices, ktane.ColorNoMore)
		}
		color, err := communication.ChooseOneColor(fmt.Sprintf("Farbe %d. Draht", index+1), choices)
		if err != nil {
			return err
		}

		if color == ktane.ColorNoMore {
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
	if s.countColor(ktane.ColorRed) == 0 {
		s.cut(2)
		return
	}
	if s.colors[2] == ktane.ColorWhite {
		s.cut(3)
	}
	if s.countColor(ktane.ColorBlue) > 1 {
		s.cut(s.getLastIndex(ktane.ColorBlue) + 1)
		return
	}
	s.cut(3)
}

func (s SimpleWiresModule) handleFour() {
	if s.countColor(ktane.ColorRed) > 1 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(s.getLastIndex(ktane.ColorRed) + 1)
			return
		}
	}

	if s.colors[3] == ktane.ColorYellow && s.countColor(ktane.ColorRed) == 0 {
		s.cut(1)
		return
	}

	if s.countColor(ktane.ColorBlue) == 1 {
		s.cut(1)
		return
	}

	if s.countColor(ktane.ColorYellow) > 1 {
		s.cut(4)
		return
	}

	s.cut(2)
	return
}

func (s SimpleWiresModule) handleFive() {
	if s.colors[4] == ktane.ColorBlack {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(4)
			return
		}
	}

	if s.countColor(ktane.ColorRed) == 1 && s.countColor(ktane.ColorYellow) > 1 {
		s.cut(1)
		return
	}

	if s.countColor(ktane.ColorBlack) == 0 {
		s.cut(2)
		return
	}

	s.cut(1)
	return
}

func (s SimpleWiresModule) handleSix() {
	if s.countColor(ktane.ColorYellow) == 0 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(3)
			return
		}
	}

	if s.countColor(ktane.ColorYellow) == 1 && s.countColor(ktane.ColorWhite) > 1 {
		s.cut(4)
		return
	}

	if s.countColor(ktane.ColorRed) == 0 {
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

func (s SimpleWiresModule) countColor(searchColor ktane.Color) int {
	result := 0
	for _, color := range s.colors {
		if color != searchColor {
			continue
		}
		result++
	}
	return result
}

func (s SimpleWiresModule) getLastIndex(searchColor ktane.Color) int {
	result := 0
	for i, color := range s.colors {
		if color != searchColor {
			continue
		}
		result = i
	}
	return result
}
