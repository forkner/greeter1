package greeter1

import (
	"fmt"

	"github.com/forkner/greetings"
)

func SayHello() {
	fmt.Printf("Greeter1: %s!\n", greetings.Greeting)
}
