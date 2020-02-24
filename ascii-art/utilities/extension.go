package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	GetSymbolsFrom - Создаем двумерный массив, из определенного файла
	в формате [Cимвол1[8(каждый подсимвол является этажом)], Cимвол2[8]...94])
*/
func GetSymbolsFrom(fileName string) (symbols [][]string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	content := strings.Split(string(data), "\n")[1:]
	for i := 0; i < len(content); i += 9 {
		symbols = append(symbols, content[i:i+8])
	}

	return
}

// GetIndexesOfSymbols function
func GetIndexesOfSymbols(symbols [][]string, indexes [][]int, word string) [][]int {
	line := []int{}
	newLineIndex := strings.Index(word, "\r")
	for index, item := range word {
		if newLineIndex == index {
			indexes = append(indexes, line)
			indexes = GetIndexesOfSymbols(symbols, indexes, word[index+2:])
			break
		} else {
			line = append(line, int(item)-32)
		}
	}

	if newLineIndex == -1 {
		indexes = append(indexes, line)
	}

	return indexes
}
