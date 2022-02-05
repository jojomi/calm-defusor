package modules

import (
	"fmt"

	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/calm-defusor/state"
)

// Optimization: Ask for number of wires upfront and tell to cut last before talking about it if not yet solved

const (
	complexWiresHasRed = 1 << iota
	complexWiresHasBlue
	complexWiresHasLED
	complexWiresHasStar
)

type ComplexWiresModule struct {
	hasParallelPortCache *bool
}

func (c ComplexWiresModule) Name() string {
	return "Komplexe Drähte"
}

func (c ComplexWiresModule) String() string {
	return c.Name()
}

func NewComplexWiresModule() *ComplexWiresModule {
	return &ComplexWiresModule{}
}

func (c *ComplexWiresModule) Reset(_ *state.BombState) error {
	c.hasParallelPortCache = nil
	return nil
}

func (c *ComplexWiresModule) Solve(bombState *state.BombState) error {
	index := 1

	var (
		val int
		err error
	)
	for {
		fmt.Println()
		communication.Tellf("Nenne mir die Eigenschaften von Draht Nummer %d", index)

		fmt.Println()
		val, err = c.getWireValue()
		if err != nil {
			return err
		}

		err = c.handleWire(val, bombState)
		if err != nil {
			return err
		}

		fmt.Println()
		index++
	}
}

func (c *ComplexWiresModule) getWireValue() (int, error) {
	result := 0

	hasRed, err := communication.ConfirmNoDefault("Hat der Draht Rotanteil?")
	if err != nil {
		return 0, err
	}
	if hasRed {
		result = result | complexWiresHasRed
	}

	hasBlue, err := communication.ConfirmNoDefault("Hat der Draht Blauanteil?")
	if err != nil {
		return 0, err
	}
	if hasBlue {
		result = result | complexWiresHasBlue
	}

	hasLED, err := communication.ConfirmNoDefault("Leuchtet die runde LED oben?")
	if err != nil {
		return 0, err
	}
	if hasLED {
		result = result | complexWiresHasLED
	}

	hasStar, err := communication.ConfirmNoDefault("Leuchtet der Stern unten?")
	if err != nil {
		return 0, err
	}
	if hasStar {
		result = result | complexWiresHasStar
	}

	return result, nil
}

func (c *ComplexWiresModule) handleWire(val int, bombState *state.BombState) error {
	switch val {
	// all four
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasLED | complexWiresHasStar:
		c.handleN()
	// three
	case complexWiresHasBlue | complexWiresHasLED | complexWiresHasStar:
		return c.handleP()
	case complexWiresHasRed | complexWiresHasLED | complexWiresHasStar:
		return c.handleB(bombState)
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasStar:
		return c.handleP()
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasLED:
		return c.handleS(bombState)
	// two
	case complexWiresHasRed | complexWiresHasBlue:
		return c.handleS(bombState)
	case complexWiresHasRed | complexWiresHasLED:
		return c.handleB(bombState)
	case complexWiresHasRed | complexWiresHasStar:
		c.handleD()
	case complexWiresHasBlue | complexWiresHasLED:
		return c.handleP()
	case complexWiresHasBlue | complexWiresHasStar:
		c.handleN()
	case complexWiresHasLED | complexWiresHasStar:
		return c.handleB(bombState)
	// only one
	case complexWiresHasRed:
		return c.handleS(bombState)
	case complexWiresHasBlue:
		return c.handleS(bombState)
	case complexWiresHasLED:
		c.handleN()
	case complexWiresHasStar:
		c.handleD()
	// none
	case 0:
		c.handleD()
	}
	return nil
}

func (c *ComplexWiresModule) handleD() {
	communication.Tell("Draht durchtrennen!\nAlle exakt gleichen Drähte, egal wo, auch durchtrennen.")
}

func (c *ComplexWiresModule) handleN() {
	communication.Tell("Draht NICHT durchtrennen!\nAlle exakt gleichen Drähte, egal wo, nicht mehr mitteilen.")
}

func (c *ComplexWiresModule) handleS(bombState *state.BombState) error {
	serial, err := bombState.Serial.GetLastDigit()
	if err != nil {
		return err
	}

	isSerialLastDigitEven := serial%2 == 0
	if isSerialLastDigitEven {
		c.handleD()
		return nil
	}
	c.handleN()
	return nil
}

func (c *ComplexWiresModule) handleP() error {
	var (
		hasParallelPort bool
		err             error
	)
	if c.hasParallelPortCache != nil {
		hasParallelPort = *c.hasParallelPortCache
	} else {
		hasParallelPort, err = communication.ConfirmNoDefault("Gibt es einen Parallel-Port-Anschluss?")
		if err != nil {
			return err
		}
		c.hasParallelPortCache = &hasParallelPort
	}

	if hasParallelPort {
		c.handleD()
		return nil
	}
	c.handleN()
	return nil
}

func (c *ComplexWiresModule) handleB(bombState *state.BombState) error {
	batterieCount, err := bombState.Batteries.GetCount()
	if err != nil {
		return err
	}

	hasMoreThanOneBattery := batterieCount > 1
	if hasMoreThanOneBattery {
		c.handleD()
		return nil
	}
	c.handleN()
	return nil
}
