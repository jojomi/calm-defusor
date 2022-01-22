package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/go-script/v2/interview"
	"github.com/rs/zerolog/log"
)

type BigButtonModule struct {
	allColors      []string
	allStripColors []string
	allTexts       []string
}

func (b *BigButtonModule) Name() string {
	return "Großer Knopf"
}

func NewBigButtonModule() *BigButtonModule {
	return &BigButtonModule{
		allColors:      []string{"red", "blue", "white", "yellow", "black"},
		allStripColors: []string{"blue", "white", "yellow", "OTHER"},
		allTexts:       []string{"Abbrechen", "Gedrückt halten", "Sprengen", "ANDERE"},
	}
}

func (b *BigButtonModule) Solve() error {
	color, err := interview.ChooseOneString("Farbe Knopf?", b.allColors)
	if err != nil {
		return err
	}

	// 1.
	if color == "blue" {
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
	if color == "white" {
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
	if color == "yellow" {
		log.Info().Msg("Regel 5 triggered")
		return b.timedRelease()
	}

	// 6.
	if color == "red" {
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
	communication.Tell(fmt.Sprintf("Knopf drücken und gedrückt halten. Loslassen, wenn der Timer an einer beliebigen Stelle eine %d enthält.", value))
}

func (b *BigButtonModule) getText() (string, error) {
	return interview.ChooseOneString("Text auf dem Knopf?", b.allTexts)
}
