package communication

import (
	"fmt"
	"github.com/gookit/color"
)

func TellSprint(text string) string {
	colorPrinter := color.New(color.FgYellow, color.BgBlack, color.Bold)
	return colorPrinter.Sprint(text)
}

func TellSprintf(text string, values ...interface{}) string {
	return TellSprint(fmt.Sprintf(text, values...))
}

func Tell(text string) {
	fmt.Println(TellSprint(text))
}

func Tellf(text string, values ...interface{}) {
	fmt.Print(TellSprintf(text, values...))
}
