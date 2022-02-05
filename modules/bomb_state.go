package modules

import (
	"github.com/jojomi/calm-defusor/state"
)

type BombStateModule struct {
}

func (s *BombStateModule) Name() string {
	return "Bomb State verwalten"
}

func (s *BombStateModule) String() string {
	return s.Name()
}

func NewBombStateModule() *BombStateModule {
	return &BombStateModule{}
}

func (s *BombStateModule) Reset(_ *state.BombState) error {
	// irrelevant
	return nil
}

func (s *BombStateModule) Solve(bombState *state.BombState) error {
	// battery
	if !bombState.Batteries.HasCount() {
		err := bombState.Batteries.AskCount()
		if err != nil {
			return err
		}
	}

	// serial
	if !bombState.Serial.HasVowelSet() {
		err := bombState.Serial.AskHasVowel()
		if err != nil {
			return err
		}
	}
	if !bombState.Serial.HasLastDigit() {
		err := bombState.Serial.AskLastDigit()
		if err != nil {
			return err
		}
	}

	return nil
}
