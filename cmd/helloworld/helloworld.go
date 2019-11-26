package helloworld

import "fmt"

//return hello greeting
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello There!"
	}
	return fmt.Sprintf("Hello %v!", user)
}
