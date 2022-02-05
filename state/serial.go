package state

import (
	"github.com/jojomi/calm-defusor/communication"
)

type Serial struct {
	hasVowel  *bool
	lastDigit *int
}

func (x *Serial) SetHasVowel(value bool) {
	x.hasVowel = &value
}

func (x Serial) HasVowelSet() bool {
	return x.hasVowel != nil
}

func (x *Serial) AskHasVowel() error {
	value, err := communication.ConfirmNoDefault("Beinhaltet die Seriennummer einen Vokal?")
	if err == nil {
		x.SetHasVowel(value)
	}
	return err
}

func (x *Serial) HasVowel() (bool, error) {
	if !x.HasVowelSet() {
		err := x.AskHasVowel()
		if err != nil {
			return false, err
		}
	}

	return *x.hasVowel, nil
}

func (x *Serial) SetLastDigit(value int) {
	x.lastDigit = &value
}

func (x Serial) HasLastDigit() bool {
	return x.lastDigit != nil
}

func (x *Serial) AskLastDigit() error {
	lastDigit, err := communication.AskInt("Letzte Ziffer in der Seriennummer?")
	if err == nil {
		x.SetLastDigit(lastDigit)
	}
	return err
}

func (x *Serial) GetLastDigit() (int, error) {
	if !x.HasLastDigit() {
		err := x.AskLastDigit()
		if err != nil {
			return 0, err
		}
	}

	return *x.lastDigit, nil
}
