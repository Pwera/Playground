package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/vmihailenco/msgpack"
	"gopkg.in/yaml.v2"
)

func encodeYaml(w io.Writer, i interface{}) {
	yaml.NewEncoder(w).Encode(&i)
}
func encodeXML(w io.Writer, i interface{}) {
	xml.NewEncoder(w).Encode(&i)
}
func encodeMsgPack(w io.Writer, i interface{}) {
	msgpack.NewEncoder(w).Encode(&i)
}
func encodeJson(w io.Writer, i interface{}) {
	json.NewEncoder(w).Encode(&i)
}

type UserType int

const (
	Personal UserType = iota
	Business
	Institution
)

type Target struct {
	ID       string   `json:"id,omitempty"`
	Usertype UserType `json:"usertype,omitempty"`
}
type UnknowType struct {
	foo interface{}
}

func encodeKnownTypes() {
	f, err := os.Create("file")
	if err != nil {
		panic(err)
	}
	target := &Target{ID: "id1", Usertype: Institution}
	encodeYaml(f, target)
	encodeXML(f, target)
	encodeMsgPack(f, target)
	encodeJson(f, target)
}
func encodeUnknownTypes() {
	var unknowntypes []interface{}
	jsonStr := `[{"foo":123},{"foo":444}]`
	err := json.Unmarshal([]byte(jsonStr), &unknowntypes)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}
	for i := range unknowntypes {
		switch v := unknowntypes[i].(type) {
		case int:
			fmt.Printf("%d int : %d\n", i, v)
		case string:
			fmt.Printf("%d str : %s\n", i, v)
		case nil:
			fmt.Printf("%d null : %s\n", i, v)
		case map[string]interface{}:
			fmt.Printf("%d map[string]interface {} : %s\n", i, v)
		default:
			fmt.Printf("%d : %T\n", i, v)
		}
	}
}
func main() {
	encodeKnownTypes()
	encodeUnknownTypes()
}
