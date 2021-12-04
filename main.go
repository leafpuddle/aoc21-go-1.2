package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(inputf)

	inputs := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	result := 0

	for i := 0; i < len(inputs); i++ {
		if i == 0 || i >= len(inputs)-2 {
			continue
		}

		n1, err := strconv.ParseInt(inputs[i-1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.ParseInt(inputs[i], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		p1, err := strconv.ParseInt(inputs[i+1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		p2, err := strconv.ParseInt(inputs[i+2], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		if (b + p1 + p2) > (n1 + b + p1) {
			result++
		}
	}

	fmt.Print(result)
}
