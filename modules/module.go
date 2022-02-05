package modules

import (
	"github.com/jojomi/calm-defusor/state"
)

type Module interface {
	Name() string
	Reset(bombState *state.BombState) error
	Solve(bombState *state.BombState) error
	String() string
}

func mapList[T, U any](in []T, mapperFunc func(t T) U) []U {
	result := make([]U, len(in))
	for i, t := range in {
		result[i] = mapperFunc(t)
	}
	return result
}

func maxMapper[T any, U int](ts []T, evaluator func(t T) U) U {
	var max U
	for i, t := range ts {
		if v := evaluator(t); i == 0 || v > max {
			max = v
		}
	}
	return max
}

func inArray[T comparable](array []T, elem T) bool {
	for _, t := range array {
		if t != elem {
			continue
		}
		return true
	}
	return false
}
