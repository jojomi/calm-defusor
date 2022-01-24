package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane_color"
	"github.com/jojomi/go-script/v2/interview"
)

type WireScreensModule struct {
	allColors []ktane_color.Color
}

func (x *WireScreensModule) Name() string {
	return "Drahtfolgen"
}

func NewWireScreensModule() *WireScreensModule {
	return &WireScreensModule{
		allColors: []ktane_color.Color{
			ktane_color.ColorRed,
			ktane_color.ColorBlue,
			ktane_color.ColorBlack,
		},
	}
}

func (x *WireScreensModule) Solve() error {
	connectionMap := map[ktane_color.Color][][]string{
		ktane_color.ColorRed: {
			{"C"},
			{"B"},
			{"A"},
			{"A", "C"},
			{"B"},
			{"A", "C"},
			{"A", "B", "C"},
			{"A", "B"},
			{"B"},
		},
		ktane_color.ColorBlue: {
			{"B"},
			{"A", "C"},
			{"B"},
			{"A"},
			{"B"},
			{"B", "C"},
			{"C"},
			{"A", "C"},
			{"A"},
		},
		ktane_color.ColorBlack: {
			{"A", "B", "C"},
			{"A", "C"},
			{"B"},
			{"A", "C"},
			{"B"},
			{"B", "C"},
			{"A", "B"},
			{"C"},
			{"C"},
		},
	}

	communication.Tell("Gehe dir Drähte von oben nach unten durch. Wenn alle dran waren, auf die nächste Tafel schalten und auf die gleiche Art weitermachen.")

	colors := append(x.allColors, ktane_color.ColorNoMore)
	var (
		col    ktane_color.Color
		target string
		err    error
	)

	colorIndex := make(map[ktane_color.Color]int)
	for _, col := range x.allColors {
		colorIndex[col] = 0
	}

	for {
		col, err = communication.ChooseOneColor("Farbe Draht?", colors)
		if err != nil {
			return err
		}

		if col == ktane_color.ColorNoMore {
			break
		}

		target, err = interview.ChooseOneString("Geht zu?", []string{"A", "B", "C"})
		if err != nil {
			return err
		}

		doCut := inArray[string](connectionMap[col][colorIndex[col]], target)

		colorIndex[col]++
		if doCut {
			communication.Tell("Draht durchschneiden")
		} else {
			communication.Tell("Draht NICHT durchschneiden")
		}
		fmt.Println()
	}

	return nil
}
