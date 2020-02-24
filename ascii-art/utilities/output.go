package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Output - создает файл
func Output(fileName, word string) {
	content, _ := ioutil.ReadFile(fileName)
	err := ioutil.WriteFile(fileName, []byte(string(content)+word), 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
