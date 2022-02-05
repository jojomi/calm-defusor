package state

import (
	"github.com/jojomi/calm-defusor/communication"
)

type Batteries struct {
	count *int
}

func (x *Batteries) SetCount(value int) {
	x.count = &value
}

func (x Batteries) HasCount() bool {
	return x.count != nil
}

func (x *Batteries) AskCount() error {
	count, err := communication.AskInt("Gesamtzahl Batterien?")
	if err == nil {
		x.SetCount(count)
	}
	return err
}

func (x *Batteries) GetCount() (int, error) {
	if !x.HasCount() {
		err := x.AskCount()
		if err != nil {
			return 0, err
		}
	}

	return *x.count, nil
}
