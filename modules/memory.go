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
	maxNum   int
	tellOnly bool
}

func (m MemoryModule) Name() string {
	return "Memory"
}

func (m MemoryModule) String() string {
	return m.Name()
}

func NewMemoryModule() *MemoryModule {
	return &MemoryModule{
		maxNum: 4,
	}
}

func (m *MemoryModule) Reset() error {
	m.steps = []MemoryStep{}
	m.tellOnly = false
	return nil
}

func (m *MemoryModule) Solve() error {
	rounds := 5

	var (
		numberShown int
		err         error
	)
	for i := 0; i < rounds; i++ {
		numberShown, err = communication.AskInt("Wie lautet die große Zahl?")
		if err != nil {
			return err
		}
		if numberShown < 1 || numberShown > m.maxNum {
			communication.Tellf("Die Zahl %d gibt es nicht. Nochmal!\n\n", numberShown)
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
				err = m.pushPos(2)
			case 2:
				err = m.pushPos(2)
			case 3:
				err = m.pushPos(3)
			case 4:
				err = m.pushPos(4)
			}
		// Stufe 2
		case 1:
			switch numberShown {
			case 1:
				err = m.pushValue(4)
			case 2:
				err = m.pushPosLike(1) // 1-based
			case 3:
				err = m.pushPos(1)
			case 4:
				err = m.pushPosLike(1) // 1-based
			}
		// Stufe 3
		case 2:
			switch numberShown {
			case 1:
				err = m.pushValueLike(2) // 1-based
			case 2:
				err = m.pushValueLike(1) // 1-based
			case 3:
				err = m.pushPos(3)
			case 4:
				err = m.pushValue(4)
			}
		// Stufe 4
		case 3:
			switch numberShown {
			case 1:
				err = m.pushPosLike(1) // 1-based
			case 2:
				err = m.pushPos(1)
			case 3:
				err = m.pushPosLike(2)
			case 4:
				err = m.pushPosLike(2)
			}
		// Stufe 5
		case 4:
			switch numberShown {
			case 1:
				err = m.pushValueLike(1) // 1-based
			case 2:
				err = m.pushValueLike(2) // 1-based
			case 3:
				err = m.pushValueLike(4) // 1-based
			case 4:
				err = m.pushValueLike(3) // 1-based
			}
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MemoryModule) pushPos(pos int) error {
	var (
		value int
		err   error
	)
	for {
		value, err = communication.AskInt(fmt.Sprintf("Wie lautet die %d. kleine Zahl?", pos))
		if err != nil {
			return err
		}
		if value > 0 && value <= m.maxNum {
			break
		}
		communication.Tellf("Die Zahl %d gibt es nicht. Nochmal!", value)
	}

	m.steps = append(m.steps, MemoryStep{
		Position: pos,
		Value:    value,
	})
	communication.Tell(fmt.Sprintf("Diese Zahl drücken. Es ist die %d an Position %d von links.", value, pos))
	fmt.Println()
	return nil
}

func (m *MemoryModule) pushValue(value int) error {
	pos, err := communication.AskInt(fmt.Sprintf("An welcher Position steht die kleine %d?", value))
	if err != nil {
		return err
	}
	m.steps = append(m.steps, MemoryStep{
		Position: pos,
		Value:    value,
	})
	communication.Tell(fmt.Sprintf("Diese Zahl drücken. Es ist die %d an Position %d von links.", value, pos))
	fmt.Println()
	return nil
}

func (m *MemoryModule) pushPosLike(step int) error {
	pos := m.steps[step-1].Position

	if m.tellOnly {
		communication.Tell(fmt.Sprintf("%d. Zahl von links drücken.", pos))
		return nil
	}

	return m.pushPos(pos)
}

func (m *MemoryModule) pushValueLike(step int) error {
	value := m.steps[step-1].Value

	if m.tellOnly {
		communication.Tell(fmt.Sprintf("Kleine %d drücken.", value))
		return nil
	}

	return m.pushValue(value)
}
