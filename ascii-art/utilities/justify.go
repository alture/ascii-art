package utilities

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Size - terminal  struct
type Size struct {
	Width, Height int
}

// Align - set align of text
func Align(align, str string) {
	size := getTerminalSize()
	strLines := strings.Split(str, "\n")
	switch align {
	case "center":
		printWithSpace((size.Width+len(strLines[0]))/2, strLines)
	case "right":
		printWithSpace(size.Width, strLines)
	default:
		fmt.Print(str)
	}
}

func getTerminalSize() *Size {
	size := &Size{}
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdout
	out, _ := cmd.Output()

	midIndex := bytes.Index(out, []byte(" "))
	size.Width, _ = strconv.Atoi(string(out[midIndex+1 : len(out)-1]))
	size.Height, _ = strconv.Atoi(string(out[:midIndex]))

	return size
}

func printWithSpace(spaceWidth int, strLines []string) {
	for _, line := range strLines {
		fmt.Println(fmt.Sprintf("%"+strconv.Itoa(spaceWidth)+"v", line))
	}
}
