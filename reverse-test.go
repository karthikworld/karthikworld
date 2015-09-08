package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("employee.txt")

	if err == nil {

		lines := strings.Split(string(content), "\n")

		for i := 1; i <= len(lines); i++ {

			fmt.Println(lines[len(lines)-i])

		}

	} else {

		log.Fatal(err)
	}
}
