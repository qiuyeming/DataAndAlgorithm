package simpleselect

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	var intarr = []int{2, 3, 1, 5, 2, 6, 4}
	Sort(intarr)
	assert.Equal(t, []int{1, 2, 2, 3, 4, 5, 6}, intarr)
	intarr = append(intarr,0)
	Sort(intarr)
	assert.Equal(t, []int{0, 1, 2, 2, 3, 4, 5, 6}, intarr)
}
