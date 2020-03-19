package customWriter

import (
	"fmt"
	"io"
	"net/http"
)

type customWriter struct{}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Print(string(bs))
	return len(bs), nil
}

func ExampleCustomWriter() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Err %v", err)
	}
	fmt.Println(resp)
	io.Copy(customWriter{}, resp.Body)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("%v %v\n", k, v)
	}
}
