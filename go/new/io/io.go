package io

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ExampleIo() {
	payload := []byte{137, 80, 78, 71, 13, 10, 26, 10, 11, 12, 14}
	buf := make([]byte, 4)

	_, err := io.ReadFull(bytes.NewReader(payload), buf)
	if err != nil {
		log.Fatal("error reading png data")
	}

	fmt.Println(buf)

	//alternative version
	io.WriteString(os.Stdout, string(buf))

	lr := io.LimitReader(bytes.NewReader(payload), 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	tmpFile, err := ioutil.TempFile(".", "temp_")
	if err != nil {
		log.Fatal(err)
	}
	defer func(tmpFile *os.File) {
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
	}(tmpFile)

}
