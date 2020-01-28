package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrompt(t *testing.T) {
	var stdin bytes.Buffer

	options := []string{"yes", "no", "depends"}

	stdin.Write([]byte("yes\n"))
	selected, err := Prompt(&stdin, "", 1, options...)
	assert.Equal(t, "yes", selected)
	assert.Nil(t, err)

	stdin.Reset()
	stdin.Write([]byte("no\n"))
	selected, err = Prompt(&stdin, "", 1, options...)
	assert.Equal(t, "no", selected)
	assert.Nil(t, err)

	stdin.Reset()
	stdin.Write([]byte("depends\n"))
	selected, err = Prompt(&stdin, "", 1, options...)
	assert.Equal(t, "depends", selected)
	assert.Nil(t, err)

	stdin.Reset()
	stdin.Write([]byte("\n"))
	selected, err = Prompt(&stdin, "", 1, options...)
	assert.Equal(t, options[1], selected)
	assert.Nil(t, err)

	stdin.Reset()
	stdin.Write([]byte("maybe\n"))
	selected, err = Prompt(&stdin, "", 1, options...)
	assert.Equal(t, options[1], selected)
	assert.Nil(t, err)

	stdin.Reset()
	stdin.Write([]byte("maybe\n"))
	selected, err = Prompt(&stdin, "", 1)
	assert.Equal(t, "", selected)
	assert.NotNil(t, err)
}
