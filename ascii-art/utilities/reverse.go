package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Reverse function
func Reverse(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	content := strings.Split(string(data), "\n")
	reversedSymbols := getReversedSymbols(content)
	fonts := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}

	for _, font := range fonts {
		str := getConvertedStr(reversedSymbols, GetSymbolsFrom(font))

		if str != " " {
			fmt.Println(str)
			return
		}
	}
}

func getReversedSymbols(content []string) (result [][]string) {
	currentSymbols := [][]string{}
	for i := 0; i < len(content); i += 9 {
		currentSymbols = append(currentSymbols, content[i:i+8])
	}

	for _, word := range currentSymbols {
		startingIndex := 0
		spaces := 0
	loop:
		for x := 0; x < len(word[0]); x++ {
			line := make([]string, 8)
			for y := 0; y < 8; y++ {
				if string(word[y][x]) != " " {
					continue loop
				} else if y == 7 {
					spaces++
				} else {
					continue
				}
			}

			// Check for space
			if x-startingIndex == 0 || spaces < 6 && spaces > 1 {
				continue
			}

			for i := 0; i < 8; i++ {
				line[i] = word[i][startingIndex : x+1]
			}
			spaces = 0
			result = append(result, line)
			startingIndex = x + 1
		}
	}

	return
}

func getConvertedStr(symbols, targetSymbols [][]string) (str string) {
	for _, symbol := range symbols {
	loop:
		for index := 0; index < len(targetSymbols); index++ {
			targetSymbol := targetSymbols[index]
			if len(symbol) != len(targetSymbol) {
				continue
			}

			for i := range symbol {
				if symbol[i] != targetSymbol[i] {
					continue loop
				}
			}

			str += string(rune(index + 32))
			break
		}
	}

	return
}
