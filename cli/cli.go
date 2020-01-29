// Package cli provides various command-line interface utilities
package cli

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Prompt displays a CLI prompt to the user. It prints the specified message,
// and gives a list of options, with a preselected default (-1 to not set any default).
// It reads the user input from the specified input stream (e.g. os.Stdin) and
// returns the selected option as a string.
func Prompt(stream io.Reader, message string, defaultIdx int, options ...string) (string, error) {
	if len(options) < 1 {
		return "", errors.New("no options specified")
	}

	validOptions := map[string]bool{}

	var buf bytes.Buffer
	buf.WriteString(message)

	buf.WriteString(" (")
	for i, o := range options {
		validOptions[strings.ToLower(o)] = true

		if i == defaultIdx {
			buf.WriteString(strings.Title(o))
		} else {
			buf.WriteString(o)
		}

		if i < len(options)-1 {
			buf.WriteString("/")
		}
	}
	buf.WriteString(") ")

	reader := bufio.NewReader(stream)
	for {
		fmt.Print(buf.String())
		selected, _ := reader.ReadString('\n')
		selected = strings.TrimSpace(selected)

		if selected == "" {
			return options[defaultIdx], nil
		}

		if valid, _ := validOptions[strings.ToLower(selected)]; valid {
			return selected, nil
		}
	}
}
