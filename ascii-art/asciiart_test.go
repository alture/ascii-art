package main

import (
	"os/exec"
	"strings"
	"testing"
)

type ASCIITestCase struct {
	Arg      string
	Expected []string
}

func TestAsciiArt(t *testing.T) {
	exec.Command("go", "build", "main.go")
	for _, item := range generateTestCase() {
		cmd := exec.Command("./ascii-art", item.Arg)
		bytes, _ := cmd.CombinedOutput()
		out := string(bytes)
		expectedStr := strings.Join(item.Expected[:], "")
		if out != expectedStr {
			t.Errorf("For \"%v\" expected\n%v\ngot \n%v", item.Arg, expectedStr, out)
		}
	}
}

func generateTestCase() []*ASCIITestCase {
	return []*ASCIITestCase{
		{"hello", []string{
			" _                 _          \n",
			"| |            | | | |         \n",
			"| |__     ___  | | | |   ___   \n",
			"|  _ \\   / _ \\ | | | |  / _ \\  \n",
			"| | | | |  __/ | | | | | (_) | \n",
			"|_| |_|  \\___| |_| |_|  \\___/  \n",
			"                               \n",
			"                               \n",
		}},
		{"HELLO", []string{
			" _    _   ______   _        _         ____   \n",
			"| |  | | |  ____| | |      | |       / __ \\  \n",
			"| |__| | | |__    | |      | |      | |  | | \n",
			"|  __  | |  __|   | |      | |      | |  | | \n",
			"| |  | | | |____  | |____  | |____  | |__| | \n",
			"|_|  |_| |______| |______| |______|  \\____/  \n",
			"                                             \n",
			"                                             \n",
		}},
	}
}
