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
var dt = make(chan []string, 225)
var a, b int

func main() {
	listPath := os.Args[1]

	readFile(listPath, func(s string) {
		wg.Add(1)
		readLine(s, handleFile)
	})

	for i := range dt {
		wg.Done()
		a++
		fmt.Println(a, "---------------------------------------------", i,
			"==========================================")
	}

	wg.Wait()
}

func handleFile(s string) string {
	go readFile(s, func(s string) {
		wg.Add(1)
		readLine(s, replaceKeyWord, dt)
	})
	return ""
}

func replaceKeyWord(s string) string {
	formIdx := strings.Index(s, "Form")
	importIdx := strings.Index(s, "import")
	hasForm := formIdx != -1
	hasImport := importIdx != -1

	if hasForm && hasImport {
		return strings.Replace(s, "Form", "", 1)
	}
	return s
}

type CB func(string)
type RCB func(string) string
type DT chan<- []string

func readLine(file string, callback RCB, data ...DT) {
	var d DT
	var l []string
	if len(data) != 0 {
		d = data[0]
	}
	stringReader := strings.NewReader(file)
	reader := bufio.NewReader(stringReader)
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}
		if d != nil {
			l = append(l, callback(string(line)))
		}
		callback(string(line))
	}
	if d != nil {
		wg.Add(1)
		b++
		fmt.Println("b", b)
		d <- l
	}
	wg.Done()
}

func readFile(path string, callback CB) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	content := string(f)
	callback(content)
}
