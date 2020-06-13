package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 7, 5, 3, 1}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 5, els[2])
	assert.EqualValues(t, 7, els[3])
	assert.EqualValues(t, 9, els[4])

}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{9, 7, 5, 3, 1}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 5, els[2])
	assert.EqualValues(t, 7, els[3])
	assert.EqualValues(t, 9, els[4])

}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)

}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkGoNativeSort100000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
