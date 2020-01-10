package main

import (
	"github.com/pwera/Playground/src/main/go/snippets/app1/db"
	"log"
)

func main(){
	connector := db.NewConnector()
	options, err := connector.LoadOptions()
	defer connector.CloseDb()
	if err != nil{
		log.Fatalf("Couldn't load option %v", err)
	}
	_=options

}
