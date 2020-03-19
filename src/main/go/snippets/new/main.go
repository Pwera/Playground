package main

import (
	"fmt"
	"github.com/pwera/new/blockchain"
	"github.com/pwera/new/ddd"
	"github.com/pwera/new/io"
	"github.com/pwera/new/logger"
	"github.com/pwera/new/printProcessList"
	"github.com/pwera/new/stacktrace"
	"github.com/pwera/new/wf"
)

func main() {
	if false {
		printProcessList.ExamplePrintProcess()
		io.ExampleIo()
	}

	if false {
		//Blockchain sample
		block := blockchain.NewGenesisBlock()
		newBlockchain := blockchain.NewBlockchain(block)
		transaction := blockchain.Transaction{Hash: []byte{076}}
		newBlockchain.AddBlock(transaction)
	}
	if false {
		// Print current Stacktrace
		go stacktrace.ExampleStackTrace()
	}
	if false {
		// Logger
		logger.Console.Debug("Sample1 D")
		logger.Console.Info("Sample1 I")
		logger.Console.Error("Sample1 E")
		logger.Console.SetCurrentLevel(logger.LevelError)
		logger.Console.Debug("Sample1 D")
		logger.Console.Info("Sample1 I")
		logger.Console.Error("Sample1 E")
	}
	//wf
	if false {
		app := wf.New()
		app.Use(func(ctx *wf.Context) {
			ctx.AddHeader("X-Info", "Hello")
		})
		app.Get("/", func(ctx *wf.Context) {
			ctx.Send("Hello World")
		})
		app.Post("/add/user", func(ctx *wf.Context) {
			name, _ := ctx.Query("name")
			if name == "" {
				ctx.Send("Whats your name?")
			} else {
				ctx.Send(fmt.Sprintf("Got Username: %s", name))
			}
		})
		app.Run()
		//curl -X POST http://localhost:3000/add/user -d 'name=alex' -H 'Content-Type: application/x-www-form-urlencoded'
	}

	if true {
		ddd.ExamppleUserRepository()

	}

}
