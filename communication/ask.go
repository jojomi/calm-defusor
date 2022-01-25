package communication

import (
	"bufio"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/jojomi/calm-defusor/ktane"
	"github.com/jojomi/go-script/v2/interview"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"strings"
)

func AskSprint(text string) string {
	colorPrinter := color.New(color.FgLightBlue, color.BgBlack, color.Bold)
	return colorPrinter.Sprint(text)
}

func AskSprintf(text string, values ...interface{}) string {
	return AskSprint(fmt.Sprintf(text, values...))
}

func AskPrintf(text string, values ...interface{}) {
	fmt.Print(AskSprintf(text, values...))
}

func AskInt(question string) (int, error) {
	var (
		input  string
		i      int
		err    error
		reader = bufio.NewReader(os.Stdin)
	)
	for {
		fmt.Printf("(?) %s ", AskSprint(question))
		input, err = reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		i, err = strconv.Atoi(strings.TrimSpace(input))
		// good int?
		if err == nil {
			break
		}
		log.Error().Err(err).Str("raw", input).Msg("invalid input")
	}
	return i, nil
}

func AskString(question string) (string, error) {
	var (
		input  string
		err    error
		reader = bufio.NewReader(os.Stdin)
	)
	fmt.Printf("(?) %s ", AskSprint(question))
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func ConfirmNoDefault(question string) (result bool, err error) {
	return interview.ConfirmNoDefault(AskSprint(question))
}

func ChooseOneString(question string, options []string) (result string, err error) {
	return interview.ChooseOneString(AskSprint(question), options)
}

func ChooseOneColor(question string, options []ktane.Color) (ktane.Color, error) {
	return ChooseOneWithMapper[ktane.Color](question, options, func(color ktane.Color) string {
		return color.BySysLocaleForTerminal()
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
		Message: AskSprint(question),
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
