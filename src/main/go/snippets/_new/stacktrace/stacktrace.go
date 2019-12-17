package stacktrace

import (
	"fmt"
	"runtime/debug"
)

func ExampleStackTrace() {
	debug.PrintStack()
	defer debug.PrintStack()
	done := make(chan bool)

	go func(done chan bool) {
		debug.PrintStack()
		done <- true
		close(done)
	}(done)
	<-done

	starckTrace := debug.Stack()
	fmt.Printf("%v", string(starckTrace))
	//os.Stdout.Write(debug.Stack())

}
