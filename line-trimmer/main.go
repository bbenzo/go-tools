package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
// read in a file line by line and cut from end to start
func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    start, _ := strconv.Atoi(os.Args[2])
    end, _ := strconv.Atoi(os.Args[3])
    for scanner.Scan() {
    	fmt.Println(string([]rune(scanner.Text())[start:end]))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}