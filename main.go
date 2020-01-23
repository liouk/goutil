package main

import (
	"fmt"
	"goutil/cli"
)

func main() {
	selected, err := cli.Prompt("Fancy?", 1, "yes", "no", "maybe")
	if err != nil {
		panic(err)
	}

	fmt.Println(selected)
}
