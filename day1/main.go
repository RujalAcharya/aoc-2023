package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part_one()
	part_two()
}

func part_one() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sum := 0
	text := "nil"

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = scanner.Text()
		sum += part_one_calc(text)
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part_two() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sum := 0
	text := "nil"

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Printf("%s: %d\n", text, part_two_calc(text))
		sum += part_two_calc(text)
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func part_one_calc(text string) int {

	first := 0
	second := 0

	for i := 0; i < len(text); i++ {
		if text[i] > '/' && text[i] < ':' {
			first = int(text[i] - '0')
			i = len(text)
		}
	}

	for i := len(text) - 1; i >= 0; i-- {
		if text[i] > '/' && text[i] < ':' {
			second = int(text[i] - '0')
			i = 0
		}
	}

	return 10*first + second
}

func part_two_calc(text string) int {
	first := 0
	first_pos := 0
	second := 0
	second_pos := 0

	num_map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i < len(text); i++ {
		if text[i] > '/' && text[i] < ':' {
			first = int(text[i] - '0')
			first_pos = i
			i = len(text)
		}
	}

	for i := len(text) - 1; i >= 0; i-- {
		if text[i] > '/' && text[i] < ':' {
			second = int(text[i] - '0')
			second_pos = i
			i = 0
		}
	}

	for key, element := range num_map {
		if strings.Contains(text, key) {
			index := strings.Index(text, key)
			if index < first_pos {
				first_pos = index
				first = element
			} else if index > second_pos {
				second_pos = index
				second = element
			}
		}
	}

	return 10*first + second
}
