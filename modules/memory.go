package modules

import (
	"fmt"
	"github.com/jojomi/calm-defusor/communication"
)

type MemoryStep struct {
	Value    int
	Position int
}

type MemoryModule struct {
	steps    []MemoryStep
	tellOnly bool
}

func (m *MemoryModule) Name() string {
	return "Memory"
}

func NewMemoryModule() *MemoryModule {
	return &MemoryModule{}
}

func (m *MemoryModule) Solve() error {
	rounds := 5

	var (
		numberShown int
		err         error
	)
	for i := 0; i < rounds; i++ {
		numberShown, err = communication.AskInt("Große Zahl?")
		if err != nil {
			// retry
			i--
			continue
		}

		// last round?
		if i == rounds-1 {
			m.tellOnly = true
		}

		switch i {
		// Stufe 1
		case 0:
			switch numberShown {
			case 1:
				m.pushPos(2)
			case 2:
				m.pushPos(2)
			case 3:
				m.pushPos(3)
			case 4:
				m.pushPos(4)
			}
		// Stufe 2
		case 1:
			switch numberShown {
			case 1:
				m.pushValue(4)
			case 2:
				m.pushPosLike(1) // 1-based
			case 3:
				m.pushPos(1)
			case 4:
				m.pushPosLike(1) // 1-based
			}
		// Stufe 3
		case 2:
			switch numberShown {
			case 1:
				m.pushValueLike(2) // 1-based
			case 2:
				m.pushValueLike(1) // 1-based
			case 3:
				m.pushPos(3)
			case 4:
				m.pushValue(4)
			}
		// Stufe 4
		case 3:
			switch numberShown {
			case 1:
				m.pushPosLike(1) // 1-based
			case 2:
				m.pushPos(1)
			case 3:
				m.pushPosLike(2)
			case 4:
				m.pushPosLike(2)
			}
		// Stufe 5
		case 4:
			switch numberShown {
			case 1:
				m.pushValueLike(1) // 1-based
			case 2:
				m.pushValueLike(2) // 1-based
			case 3:
				m.pushValueLike(4) // 1-based
			case 4:
				m.pushValueLike(3) // 1-based
			}
		}
	}

	return nil
}

func (m *MemoryModule) pushPos(pos int) {
	value, _ := communication.AskInt(fmt.Sprintf("Wie lautet die %d. kleine Zahl?", pos))
	m.steps = append(m.steps, MemoryStep{
		Position: pos,
		Value:    value,
	})
	communication.Tell(fmt.Sprintf("Diese Zahl drücken. Es ist die %d an Position %d von links.", value, pos))
	fmt.Println()
}

func (m *MemoryModule) pushValue(value int) {
	pos, _ := communication.AskInt(fmt.Sprintf("An welcher Position steht die kleine %d?", value))
	m.steps = append(m.steps, MemoryStep{
		Position: pos,
		Value:    value,
	})
	communication.Tell(fmt.Sprintf("Diese Zahl drücken. Es ist die %d an Position %d von links.", value, pos))
	fmt.Println()
}

func (m *MemoryModule) pushPosLike(step int) {
	pos := m.steps[step-1].Position

	if m.tellOnly {
		communication.Tell(fmt.Sprintf("%d. Zahl von links drücken.", pos))
		return
	}

	m.pushPos(pos)
}

func (m *MemoryModule) pushValueLike(step int) {
	value := m.steps[step-1].Value

	if m.tellOnly {
		communication.Tell(fmt.Sprintf("Kleine %d drücken.", value))
		return
	}

	m.pushValue(value)
}
