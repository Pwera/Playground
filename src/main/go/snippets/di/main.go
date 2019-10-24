package main

import (
	"github.com/Pwera/Playground/src/main/go/snippets/di/logging"
	"github.com/sarulabs/di"
)

func main() {

	defer logging.Logger.Sync()

	builder, err := di.NewBuilder()
	if err != nil {
		logging.Logger.Fatal(err.Error)
	}

}
