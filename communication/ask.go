package communication

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/jojomi/calm-defusor/ktane_color"
)

func AskString(question string) (string, error) {
	// TODO implement
	return "", nil
}

func AskBool(question string) (bool, error) {
	fmt.Println("(?)", question)
	return true, nil
}

func AskInt(question string) (int, error) {
	fmt.Println("(?)", question)
	var i int
	_, err := fmt.Scanf("%d", &i)
	return i, err
}

func ChooseOneColor(question string, options []ktane_color.Color) (ktane_color.Color, error) {
	return ChooseOneWithMapper[ktane_color.Color](question, options, func(color ktane_color.Color) string {
		return color.BySysLocale()
	})
}

func ChooseOneStringable[T fmt.Stringer](question string, options []T) (T, error) {
	return ChooseOneWithMapper[T](question, options, func(t T) string {
		return t.String()
	})
}

func ChooseOneWithMapper[T any](question string, options []T, mapper func(t T) string) (T, error) {
	var tResult T

	// map options
	stringOpts := make([]string, len(options))
	for i, option := range options {
		stringOpts[i] = mapper(option)
	}

	prompt := &survey.Select{
		Message: question,
		Options: stringOpts,
	}

	var result string
	err := survey.AskOne(prompt, &result, nil)
	if err != nil {
		return tResult, err
	}

	// map back
	for i, option := range stringOpts {
		if option != result {
			continue
		}
		return options[i], nil
	}
	return tResult, fmt.Errorf("invalid selection: %s", result)
}
