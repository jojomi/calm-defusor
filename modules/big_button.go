package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane_color"
	"github.com/jojomi/go-script/v2/interview"
	"github.com/rs/zerolog/log"
)

type BigButtonModule struct {
	allColors      []ktane_color.Color
	allStripColors []string
	allTexts       []string

	textCache string
}

func (b *BigButtonModule) Name() string {
	return "Großer Knopf"
}

func NewBigButtonModule() *BigButtonModule {
	return &BigButtonModule{
		allColors: []ktane_color.Color{
			ktane_color.ColorRed,
			ktane_color.ColorBlue,
			ktane_color.ColorWhite,
			ktane_color.ColorYellow,
			ktane_color.ColorBlack,
		},
		allStripColors: []string{"blue", "white", "yellow", "OTHER"},
		allTexts:       []string{"Abbrechen", "Gedrückt halten", "Sprengen", "ANDERE"},
	}
}

func (b *BigButtonModule) Solve() error {
	color, err := communication.ChooseOneStringable("Farbe Knopf?", b.allColors)
	if err != nil {
		return err
	}

	// 1.
	if color == ktane_color.ColorBlue {
		text, err := b.getText()
		if err != nil {
			return err
		}
		if text == "Abbrechen" {
			log.Info().Msg("Regel 1 triggered")
			return b.timedRelease()
		}
	}

	// 2.
	numBatteries, err := communication.AskInt("Anzahl Batterien an der Bombe?")
	if err != nil {
		return err
	}
	if numBatteries > 1 {
		text, err := b.getText()
		if err != nil {
			return err
		}
		if text == "Sprengen" {
			log.Info().Msg("Regel 2 triggered")
			b.tap()
			return nil
		}
	}

	// 3.
	if color == ktane_color.ColorWhite {
		carIndicator, err := interview.ConfirmNoDefault("Hat die Bombe einen CAR Indikator?")
		if err != nil {
			return err
		}
		if carIndicator {
			log.Info().Msg("Regel 3 triggered")
			return b.timedRelease()
		}
	}

	// 4.
	if numBatteries > 2 {
		frkIndicator, err := interview.ConfirmNoDefault("Hat die Bombe einen FRK Indikator?")
		if err != nil {
			return err
		}
		if frkIndicator {
			log.Info().Msg("Regel 4 triggered")
			b.tap()
			return nil
		}
	}

	// 5.
	if color == ktane_color.ColorYellow {
		log.Info().Msg("Regel 5 triggered")
		return b.timedRelease()
	}

	// 6.
	if color == ktane_color.ColorRed {
		text, err := b.getText()
		if err != nil {
			return err
		}
		if text == "Gedrückt halten" {
			log.Info().Msg("Regel 6 triggered")
			b.tap()
			return nil
		}
	}

	// 7.
	log.Info().Msg("Regel 7 triggered")
	return b.timedRelease()
}

func (b *BigButtonModule) tap() {
	communication.Tell("Knopf kurz antippen. Fertig.")
}

func (b *BigButtonModule) timedRelease() error {
	communication.Tell("Knopf drücken und gedrückt halten.")
	color, err := interview.ChooseOneString("Farbe Streifen?", b.allStripColors)
	if err != nil {
		return err
	}
	switch color {
	case "blue":
		b.releaseAt(4)
	case "yellow":
		b.releaseAt(5)
	default:
		b.releaseAt(1)
	}
	return nil
}

func (b *BigButtonModule) releaseAt(value int) {
	communication.Tell(fmt.Sprintf("Knopf loslassen, wenn der Timer an einer beliebigen Stelle eine %d enthält.", value))
}

func (b *BigButtonModule) getText() (string, error) {
	if b.textCache != "" {
		return b.textCache, nil
	}
	text, err := interview.ChooseOneString("Text auf dem Knopf?", b.allTexts)
	b.textCache = text
	return text, err
}
