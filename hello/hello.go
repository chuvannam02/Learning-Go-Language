package main

// import "fmt"
// import "rsc.io/quote"
import (
	"fmt"

	"example.com/greetings"
	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println(reverse.String("Hello"))
}
