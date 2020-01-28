package main

import (
	"fmt"
	"goutil/cli"
	"os"
)

func main() {
	selected, err := cli.Prompt(os.Stdin, "choose:", 0, "yes", "no")
	if err != nil {
		panic(err)
	}

	fmt.Println("user selected:", selected)
}
