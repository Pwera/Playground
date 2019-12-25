package main

import (
	"github.com/pwera/Playground/src/main/go/snippets/_new/blockchain"
	"github.com/pwera/Playground/src/main/go/snippets/_new/io"
	"github.com/pwera/Playground/src/main/go/snippets/_new/logger"
)

func main() {
	//printProcessList.ExamplePrintProcess()
	io.ExampleIo()
	{
		//Blockchain sample
		block := blockchain.NewGenesisBlock()
		newBlockchain := blockchain.NewBlockchain(block)
		transaction := blockchain.Transaction{Hash: []byte{076}}
		newBlockchain.AddBlock(transaction)
	}
	{
		// Print current Stacktrace
		// go stacktrace.ExampleStackTrace()
	}
	{
		// Logger
		logger.Console.Debug("Sample1 D")
		logger.Console.Info("Sample1 I")
		logger.Console.Error("Sample1 E")
		logger.Console.SetCurrentLevel(logger.LevelError)
		logger.Console.Debug("Sample1 D")
		logger.Console.Info("Sample1 I")
		logger.Console.Error("Sample1 E")
	}

}
