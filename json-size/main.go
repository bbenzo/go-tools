package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		panic("reading file failed: " + err.Error())
	}

	fmt.Println(strconv.Itoa(len(buf)) + " bytes")
}
