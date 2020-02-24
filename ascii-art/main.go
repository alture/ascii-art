package main

import (
	"fmt"
	"os"
	"strings"

	utilities "./utilities"
)

// Flag structure
type Flag struct {
	Key, Value string
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	symbols := utilities.GetSymbolsFrom(utilities.FS(args))
	output := convertSymbols([][]int{}, symbols, args[0])
	draw(output, symbols, flag(args))
}

// Подготавливаем
func convertSymbols(output [][]int, symbols [][]string, word string) [][]int {
	line := []int{}
	newLineIndex := strings.Index(word, "\\n")
	for index, item := range word {
		if newLineIndex == index {
			output = append(output, line)
			output = convertSymbols(output, symbols, word[index+2:])
			break
		} else {
			line = append(line, int(item)-32)
		}
	}

	if newLineIndex == -1 {
		output = append(output, line)
	}

	return output
}

// Рисуем или отталкиваемся от флага
func draw(output [][]int, symbols [][]string, flag *Flag) {
	var str string
	for _, line := range output {
		for level := 0; level < 8; level++ {
			for _, index := range line {
				str += symbols[index][level]
			}

			str += "\n"
		}
	}

	switch flag.Key {
	case "--output":
		utilities.Output(flag.Value, str)
	case "--reverse":
		utilities.Reverse(flag.Value)
	case "--align":
		utilities.Align(flag.Value, str)
	default:
		fmt.Print(str)
	}
}

// Достаем flag
func flag(args []string) *Flag {
	flag := &Flag{}
	for index, item := range args {
		midIndex := strings.Index(item, "=")
		if midIndex != -1 {
			if index == 0 && item[:midIndex] != "--reverse" {
				continue
			}

			flag.Key, flag.Value = item[:midIndex], item[midIndex+1:]
			break
		}
	}

	return flag
}
