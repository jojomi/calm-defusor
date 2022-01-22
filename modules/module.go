package modules

type Module interface {
	Name() string
	Solve() error
}
