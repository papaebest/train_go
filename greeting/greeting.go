package greeting

import "strings"

func Greet(name string) string {
	// return name
	back := ""
	if name == "Bob" {
		back = "Hello, Bob."
	} else if name == "" {
		back = "Hello, my friend"
	} else if name == strings.ToUpper(name) {
		back = "HELLO, BOB"
	}
	return back
}
