package modules

import (
	"fmt"

	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/ktane"
	"github.com/jojomi/calm-defusor/state"
	"github.com/rs/zerolog/log"
)

type BigButtonModule struct {
	allColors      []ktane.Color
	allStripColors []string
	allTexts       []string

	textCache *string
}

func (b BigButtonModule) Name() string {
	return "Großer Knopf"
}

func (b BigButtonModule) String() string {
	return b.Name()
}

func NewBigButtonModule() *BigButtonModule {
	return &BigButtonModule{
		allColors: []ktane.Color{
			ktane.ColorRed,
			ktane.ColorBlue,
			ktane.ColorWhite,
			ktane.ColorYellow,
			ktane.ColorBlack,
		},
		allStripColors: []string{"blue", "white", "yellow", "OTHER"},
		allTexts:       []string{"Abbrechen", "Gedrückt halten", "Sprengen", "ANDERE"},
	}
}

func (b *BigButtonModule) Reset(_ *state.BombState) error {
	b.textCache = nil
	return nil
}

func (b *BigButtonModule) Solve(bombState *state.BombState) error {
	color, err := communication.ChooseOneColor("Farbe Knopf?", b.allColors)
	if err != nil {
		return err
	}

	// 1.
	if color.IsBlue() {
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
	numBatteries, err := bombState.Batteries.GetCount()
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
	if color.IsWhite() {
		carIndicator, err := communication.ConfirmNoDefault("Hat die Bombe einen CAR Indikator?")
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
		frkIndicator, err := communication.ConfirmNoDefault("Hat die Bombe einen FRK Indikator?")
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
	if color.IsYellow() {
		log.Info().Msg("Regel 5 triggered")
		return b.timedRelease()
	}

	// 6.
	if color.IsRed() {
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
	color, err := communication.ChooseOneString("Farbe Streifen?", b.allStripColors)
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
	communication.Tell(fmt.Sprintf("Knopf loslassen, wenn der Timer irgendwo eine %d enthält.", value))
}

func (b *BigButtonModule) getText() (string, error) {
	if b.textCache != nil {
		return *b.textCache, nil
	}
	text, err := communication.ChooseOneString("Text auf dem Knopf?", b.allTexts)
	b.textCache = &text
	return text, err
}
