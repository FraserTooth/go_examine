package main

import "fmt"

//return hello greeting
func hello(user string) string {
	if len(user) == 0 {
		return "Hello There!"
	} else {
		return fmt.Sprintf("Hello %v!", user)
	}
}
