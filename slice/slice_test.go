package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	sli := []string{"0", "1", "2"}
	assert.True(t, Contains(sli, "0"))
	assert.True(t, Contains(sli, "1"))
	assert.True(t, Contains(sli, "2"))
	assert.False(t, Contains(sli, ""))
	assert.False(t, Contains(sli, "3"))

	sli = []string{}
	assert.False(t, Contains(sli, ""))
	assert.False(t, Contains(sli, "0"))

	var nilsli []string
	assert.False(t, Contains(nilsli, ""))
	assert.False(t, Contains(nilsli, "0"))
}

func ExampleContains() {
	sli := []string{"one", "two", "three"}
	fmt.Println(Contains(sli, "four"), Contains(sli, "three"))
	// Output: false true
}

func TestFindAndRemove(t *testing.T) {
	sli, removed := FindAndRemove([]string{"0", "1", "2"}, "0")
	assert.True(t, removed)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "1")
	assert.Equal(t, sli[1], "2")

	sli, removed = FindAndRemove([]string{"0", "1", "2"}, "1")
	assert.True(t, removed)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "2")

	sli, removed = FindAndRemove([]string{"0", "1", "2"}, "2")
	assert.True(t, removed)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "1")

	sli, removed = FindAndRemove([]string{"0", "1", "2"}, "-1")
	assert.False(t, removed)
	assert.Equal(t, len(sli), 3)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "1")
	assert.Equal(t, sli[2], "2")

	sli, removed = FindAndRemove([]string{"0", "1", "2"}, "3")
	assert.False(t, removed)
	assert.Equal(t, len(sli), 3)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "1")
	assert.Equal(t, sli[2], "2")

	sli, removed = FindAndRemove([]string{}, "3")
	assert.False(t, removed)
	assert.Equal(t, len(sli), 0)

	sli, removed = FindAndRemove(nil, "3")
	assert.False(t, removed)
	assert.Equal(t, len(sli), 0)
}

func TestRemoveAtIndex(t *testing.T) {
	sli := RemoveAtIndex([]string{"0", "1", "2"}, 0)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "1")
	assert.Equal(t, sli[1], "2")

	sli = RemoveAtIndex([]string{"0", "1", "2"}, 1)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "2")

	sli = RemoveAtIndex([]string{"0", "1", "2"}, 2)
	assert.Equal(t, len(sli), 2)
	assert.Equal(t, sli[0], "0")
	assert.Equal(t, sli[1], "1")

	sli = RemoveAtIndex([]string{"0", "1", "2"}, -1)
	assert.Nil(t, sli)

	sli = RemoveAtIndex([]string{"0", "1", "2"}, 3)
	assert.Nil(t, sli)

	sli = RemoveAtIndex([]string{}, 0)
	assert.Nil(t, sli)

	sli = RemoveAtIndex(nil, 0)
	assert.Nil(t, sli)
}
