package modules

type Module interface {
	Name() string
	Solve() error
}

func mapList[T, U any](in []T, mapperFunc func(t T) U) []U {
	result := make([]U, len(in))
	for i, t := range in {
		result[i] = mapperFunc(t)
	}
	return result
}
