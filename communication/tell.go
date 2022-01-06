package communication

import "fmt"

func Tell(text string) {
	fmt.Println("->", text)
}

func Tellf(text string, values ...interface{}) {
	fmt.Printf("-> "+text+"\n", values...)
}
