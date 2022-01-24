package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
	"github.com/jojomi/go-script/v2/interview"
)

// Optimization: Ask for number of wires upfront and tell to cut last before talking about it if not yet solved

const (
	complexWiresHasRed = 1 << iota
	complexWiresHasBlue
	complexWiresHasLED
	complexWiresHasStar
)

type ComplexWiresModule struct {
	serialCache                *int
	hasParallelPortCache       *bool
	hasMoreThanOneBatteryCache *bool
}

func (c *ComplexWiresModule) Name() string {
	return "Komplexe Drähte"
}

func NewComplexWiresModule() *ComplexWiresModule {
	return &ComplexWiresModule{}
}

func (c *ComplexWiresModule) Reset() error {
	c.serialCache = nil
	c.hasParallelPortCache = nil
	c.hasMoreThanOneBatteryCache = nil
	return nil
}

func (c *ComplexWiresModule) Solve() error {
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

		err = c.handleWire(val)
		if err != nil {
			return err
		}

		fmt.Println()
		index++
	}
}

func (c *ComplexWiresModule) getWireValue() (int, error) {
	result := 0

	hasRed, err := interview.ConfirmNoDefault("Hat der Draht Rotanteil?")
	if err != nil {
		return 0, err
	}
	if hasRed {
		result = result | complexWiresHasRed
	}

	hasBlue, err := interview.ConfirmNoDefault("Hat der Draht Blauanteil?")
	if err != nil {
		return 0, err
	}
	if hasBlue {
		result = result | complexWiresHasBlue
	}

	hasLED, err := interview.ConfirmNoDefault("Leuchtet die runde LED oben?")
	if err != nil {
		return 0, err
	}
	if hasLED {
		result = result | complexWiresHasLED
	}

	hasStar, err := interview.ConfirmNoDefault("Leuchtet der Stern unten?")
	if err != nil {
		return 0, err
	}
	if hasStar {
		result = result | complexWiresHasStar
	}

	fmt.Println(result)

	return result, nil
}

func (c *ComplexWiresModule) handleWire(val int) error {
	switch val {
	// all four
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasLED | complexWiresHasStar:
		c.handleN()
	// three
	case complexWiresHasBlue | complexWiresHasLED | complexWiresHasStar:
		return c.handleP()
	case complexWiresHasRed | complexWiresHasLED | complexWiresHasStar:
		return c.handleB()
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasStar:
		return c.handleP()
	case complexWiresHasRed | complexWiresHasBlue | complexWiresHasLED:
		return c.handleS()
	// two
	case complexWiresHasRed | complexWiresHasBlue:
		return c.handleS()
	case complexWiresHasRed | complexWiresHasLED:
		return c.handleB()
	case complexWiresHasRed | complexWiresHasStar:
		return c.handleB()
	case complexWiresHasBlue | complexWiresHasLED:
		return c.handleP()
	case complexWiresHasBlue | complexWiresHasStar:
		c.handleN()
	case complexWiresHasLED | complexWiresHasStar:
		return c.handleB()
	// only one
	case complexWiresHasRed:
		return c.handleS()
	case complexWiresHasBlue:
		return c.handleS()
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
	communication.Tell("Draht durchtrennen! Alle exakt gleichen Drähte, egal wo, auch durchtrennen.")
}

func (c *ComplexWiresModule) handleN() {
	communication.Tell("Draht NICHT durchtrennen! Alle exakt gleichen Drähte, egal wo, nicht mehr mitteilen.")
}

func (c *ComplexWiresModule) handleS() error {
	var (
		serial int
		err    error
	)
	if c.serialCache != nil {
		serial = *c.serialCache
	} else {
		serial, err = communication.AskInt("Letzte Ziffer der Seriennummer?")
		if err != nil {
			return err
		}
		c.serialCache = &serial
	}

	if serial%2 == 0 {
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
		hasParallelPort, err = interview.ConfirmNoDefault("Gibt es einen Parallel-Port-Anschluss?")
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

func (c *ComplexWiresModule) handleB() error {
	var (
		hasMoreThanOneBattery bool
		err                   error
	)
	if c.hasMoreThanOneBatteryCache != nil {
		hasMoreThanOneBattery = *c.hasMoreThanOneBatteryCache
	} else {
		hasMoreThanOneBattery, err = interview.ConfirmNoDefault("Gibt es MEHR als eine Batterie?")
		if err != nil {
			return err
		}
		c.hasMoreThanOneBatteryCache = &hasMoreThanOneBattery
	}

	if hasMoreThanOneBattery {
		c.handleD()
		return nil
	}
	c.handleN()
	return nil
}
