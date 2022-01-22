package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/go-script/interview"
)

type SimpleWiresModule struct {
	allColors []string
	colors    []string
}

func (s *SimpleWiresModule) Name() string {
	return "Einfache Drähte"
}

func NewSimpleWiresModule() *SimpleWiresModule {
	return &SimpleWiresModule{
		allColors: []string{"red", "blue", "white", "yellow", "black"},
	}
}

func (s *SimpleWiresModule) Solve() error {
	// ask for colors
	index := 0
	for {
		choices := s.allColors
		if index > 2 {
			choices = append(choices, "ENDE")
		}
		color, err := interview.ChooseOneString(fmt.Sprintf("Farbe %d. Draht", index+1), choices)
		if err != nil {
			return err
		}

		if color == "ENDE" {
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
	if s.countColor("red") == 0 {
		s.cut(2)
		return
	}
	if s.colors[2] == "white" {
		s.cut(3)
	}
	if s.countColor("blue") > 1 {
		s.cut(s.getLastIndex("blue") + 1)
		return
	}
	s.cut(3)
}

func (s SimpleWiresModule) handleFour() {
	if s.countColor("red") > 1 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(s.getLastIndex("red") + 1)
			return
		}
	}

	if s.colors[3] == "yellow" && s.countColor("red") == 0 {
		s.cut(1)
		return
	}

	if s.countColor("blue") == 1 {
		s.cut(1)
		return
	}

	if s.countColor("yellow") > 1 {
		s.cut(4)
		return
	}

	s.cut(2)
	return
}

func (s SimpleWiresModule) handleFive() {
	if s.colors[4] == "black" {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(4)
			return
		}
	}

	if s.countColor("red") == 1 && s.countColor("yellow") > 1 {
		s.cut(1)
		return
	}

	if s.countColor("black") == 0 {
		s.cut(2)
		return
	}

	s.cut(1)
	return
}

func (s SimpleWiresModule) handleSix() {
	if s.countColor("yellow") == 0 {
		serial, err := communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			panic(err)
		}
		if serial%2 == 1 {
			s.cut(3)
			return
		}
	}

	if s.countColor("yellow") == 1 && s.countColor("white") > 1 {
		s.cut(4)
		return
	}

	if s.countColor("red") == 0 {
		s.cut(6)
		return
	}

	s.cut(4)
	return
}

// one-based!
func (s SimpleWiresModule) cut(index int) {
	communication.Tellf("%d. Draht trennen (dieser hat die Farbe %s)", index, s.colors[index-1])
}

func (s SimpleWiresModule) countColor(searchColor string) int {
	result := 0
	for _, color := range s.colors {
		if color != searchColor {
			continue
		}
		result++
	}
	return result
}

func (s SimpleWiresModule) getLastIndex(searchColor string) int {
	result := 0
	for i, color := range s.colors {
		if color != searchColor {
			continue
		}
		result = i
	}
	return result
}
