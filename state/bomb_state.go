package state

type BombState struct {
	Serial    *Serial
	Batteries *Batteries
}

func NewBombState() *BombState {
	return &BombState{
		Serial:    &Serial{},
		Batteries: &Batteries{},
	}
}
