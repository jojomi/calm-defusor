package communication

import "fmt"

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
