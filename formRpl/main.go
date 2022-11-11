package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	list := os.Args[1:]

	var echo chan []string = make(chan []string, len(list))

	for _, v := range list {
		wg.Add(1)
		go readFile(v, echo)
	}

	go func() {
		for k := range echo {
			fmt.Println(k)
			wg.Done()
		}
	}()

	wg.Wait()
}

func readFile(path string, data chan<- []string) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	readLine(content, data)

	fmt.Println(path)
}

func readLine(file string, data chan<- []string) {
	stringReader := strings.NewReader(file)
	reader := bufio.NewReader(stringReader)
	lineList := []string{}

	for {
		line, _, err := reader.ReadLine()

		formIdx := strings.Index(string(line), "form")
		hasForm := formIdx != -1
		if hasForm {
			fmt.Println(hasForm)
		}
		fmt.Println()

		if err != nil {
			break
		}

		lineList = append(lineList, string(line))
	}

	data <- lineList
}
