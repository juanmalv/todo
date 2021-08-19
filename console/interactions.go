package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please write down your next todo")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		return text
	}
}
