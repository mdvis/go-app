package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var echo chan []string = make(chan []string)
	list := os.Args[1:]
	for _, v := range list {
		go readFile(v, echo)
	}
	fmt.Println(<-echo)
}

func readFile(path string, data chan<- []string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(f)
	readLine(content, data)
}

func readLine(file string, data chan<- []string) {
	stringReader := strings.NewReader(file)
	reader := bufio.NewReader(stringReader)
	lineList := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineList = append(lineList, string(line))
	}
	data <- lineList
}
