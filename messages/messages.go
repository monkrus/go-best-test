package messages

import "fmt"

// Greet is ...
func Greet(name string) string {

	return fmt.Sprintf("Hello, %v!\n , name")
}

func depart(name string) string {
	return fmt.Sprintf("Bye, %v!\n , name")
}
