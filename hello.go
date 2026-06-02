// This tells the GO, we have an executable here
/*
   The main package must have func main() {...}
*/
package main

import (
	"examples/hello/greetings"
	"fmt"
	// Import your published module!
	"github.com/sujeet-banerjee/go-well-examples-coffee"
)

func main() {
	fmt.Println("Hello, World!")

	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	// From the custom module (git repo)
	fmt.Println(coffee.Brew())
}
