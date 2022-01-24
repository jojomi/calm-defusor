package communication

import (
	"fmt"
	"github.com/gookit/color"
)

func Tell(text string) {
	colorPrinter := color.New(color.FgYellow, color.BgBlack, color.Bold)
	colorPrinter.Println(text)
}

func Tellf(text string, values ...interface{}) {
	Tell(fmt.Sprintf(text, values...))
}
