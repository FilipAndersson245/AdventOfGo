package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution1() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	delta := map[byte]int{'L': -1, 'R': 1}

	p1, p2 := 0, 0
	dial := 50
	for s := range strings.FieldsSeq(string(input)) {
		turn, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}

		for range turn {
			if dial += delta[s[0]]; dial%100 == 0 {
				p2++
			}
		}

		if dial%100 == 0 {
			p1++
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
