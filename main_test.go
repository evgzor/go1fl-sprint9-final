package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	require.Len(t, generateRandomElements(-10), 0)
	require.Len(t, generateRandomElements(0), 0)
	require.Len(t, generateRandomElements(1), 1)
}

func TestMaximum(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource).Int()
	assert.Equal(t, rng, maximum([]int{rng}))
	assert.Equal(t, 0, maximum(nil))
	assert.Equal(t, maximum([]int{0, 1, 1, 1, 1, 0}), 1)
	assert.Equal(t, maximum([]int{0, -1, 1, 1, -1000000, 0}), 1)
}

func TestMaximumChunks(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource).Int() + 1
	assert.Equal(t, rng, maxChunks([]int{rng}))
	assert.Equal(t, 0, maxChunks(nil))
	assert.Equal(t, maxChunks([]int{0, 1, 1, 1, 1, 0, 1, rng}), rng)
	assert.Equal(t, maxChunks([]int{0, -1, -10, -1, -1000000, 0, 0, 0}), 0)
}

func TestBoth(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	randomNumber := randSource.Int63()%10 + 1
	for range randomNumber {
		elements := generateRandomElements(SIZE)
		require.Equal(t, maximum(elements), maxChunks(elements))
	}
}
