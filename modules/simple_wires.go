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

func (s *SimpleWiresModule) String() string {
	return s.Name()
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

func (s *SimpleWiresModule) Reset() error {
	s.colors = []ktane.Color{}
	return nil
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

		if color.IsNoMore() {
			break
		}

		s.colors = append(s.colors, color)
		index++
		if index == 6 {
			break
		}
	}

	// evaluate and ask more
	switch len(s.colors) {
	case 3:
		return s.handleThree()
	case 4:
		return s.handleFour()
	case 5:
		return s.handleFive()
	case 6:
		return s.handleSix()
	}
	return nil
}

func (s SimpleWiresModule) handleThree() error {
	if s.countColor(ktane.ColorRed) == 0 {
		s.cut(2)
		return nil
	}
	if s.colors[2].IsWhite() {
		s.cut(3)
	}
	if s.countColor(ktane.ColorBlue) > 1 {
		s.cut(s.getLastIndex(ktane.ColorBlue) + 1)
		return nil
	}
	s.cut(3)
	return nil
}

func (s SimpleWiresModule) handleFour() error {
	if s.countColor(ktane.ColorRed) > 1 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			return err
		}
		if serial%2 == 1 {
			s.cut(s.getLastIndex(ktane.ColorRed) + 1)
			return nil
		}
	}

	if s.colors[3].IsYellow() && s.countColor(ktane.ColorRed) == 0 {
		s.cut(1)
		return nil
	}

	if s.countColor(ktane.ColorBlue) == 1 {
		s.cut(1)
		return nil
	}

	if s.countColor(ktane.ColorYellow) > 1 {
		s.cut(4)
		return nil
	}

	s.cut(2)
	return nil
}

func (s SimpleWiresModule) handleFive() error {
	if s.colors[4].IsBlack() {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			return err
		}
		if serial%2 == 1 {
			s.cut(4)
			return nil
		}
	}

	if s.countColor(ktane.ColorRed) == 1 && s.countColor(ktane.ColorYellow) > 1 {
		s.cut(1)
		return nil
	}

	if s.countColor(ktane.ColorBlack) == 0 {
		s.cut(2)
		return nil
	}

	s.cut(1)
	return nil
}

func (s SimpleWiresModule) handleSix() error {
	if s.countColor(ktane.ColorYellow) == 0 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			return err
		}
		if serial%2 == 1 {
			s.cut(3)
			return nil
		}
	}

	if s.countColor(ktane.ColorYellow) == 1 && s.countColor(ktane.ColorWhite) > 1 {
		s.cut(4)
		return nil
	}

	if s.countColor(ktane.ColorRed) == 0 {
		s.cut(6)
		return nil
	}

	s.cut(4)
	return nil
}

// one-based!
func (s SimpleWiresModule) cut(index int) {
	communication.Tellf("%d. Draht trennen (dieser hat die Farbe %s", index, s.colors[index-1].BySysLocaleForTerminal())
	communication.Tellf(")")
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
