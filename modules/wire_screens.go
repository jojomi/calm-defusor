package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane"
	"github.com/jojomi/go-script/v2/interview"
)

type WireScreensModule struct {
	allColors []ktane.Color
}

func (x *WireScreensModule) Name() string {
	return "Drahtfolgen"
}

func (x *WireScreensModule) String() string {
	return x.Name()
}

func NewWireScreensModule() *WireScreensModule {
	return &WireScreensModule{
		allColors: []ktane.Color{
			ktane.ColorRed,
			ktane.ColorBlue,
			ktane.ColorBlack,
		},
	}
}

func (x *WireScreensModule) Reset() error {
	return nil
}

func (x *WireScreensModule) Solve() error {
	connectionMap := map[ktane.Color][][]string{
		ktane.ColorRed: {
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
		ktane.ColorBlue: {
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
		ktane.ColorBlack: {
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

	colors := append(x.allColors, ktane.ColorNoMore)
	var (
		col    ktane.Color
		target string
		err    error
	)

	colorIndex := make(map[ktane.Color]int)
	for _, col := range x.allColors {
		colorIndex[col] = 0
	}

	for {
		col, err = communication.ChooseOneColor("Farbe Draht?", colors)
		if err != nil {
			return err
		}

		if col == ktane.ColorNoMore {
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
