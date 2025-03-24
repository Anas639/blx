package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForConfirmatino(msg string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s\n(y/n)\n", msg)
	ans, err := reader.ReadString('\n')

	if err != nil {
		return false
	}

	if len(ans) < 2 {
		// this will be an empty string with only the newline rune '\n'
		return false
	}

	return strings.ToLower(ans)[0] == 'y'
}
