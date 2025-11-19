package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	cases := []struct {
		name string
		size int
		want int
	}{
		{"negative", -10, 0},
		{"zero", 0, 0},
		{"one", 1, 1},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := generateRandomElements(tc.size)
			require.Len(t, out, tc.want, tc.name)

		})
	}
}

func TestMaximum(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource).Int()
	cases := []struct {
		name  string
		array []int
		want  int
	}{
		{"one", []int{rng}, rng},
		{"nill", nil, 0},
		{"oneZero", []int{0, 1, 1, 1, 1, 0}, 1},
		{"one", []int{0, 100, 100, 100, 1000000, 1}, 1000000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := maximum(tc.array)
			require.Equal(t, out, tc.want, tc.name)

		})
	}
}

func TestMaxChunks(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource).Int() + 1

	cases := []struct {
		name  string
		array []int
		want  int
	}{
		{"one", []int{rng}, rng},
		{"nill", nil, 0},
		{"oneZero", []int{0, 1, 1, 1, 1, 0}, 1},
		{"one", []int{0, 100, 100, 100, 1000000, 1}, 1000000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := maxChunks(tc.array)
			assert.Equal(t, out, tc.want, tc.name)

		})
	}

}

func TestBoth(t *testing.T) {
	var randSource = rand.NewSource(time.Now().UnixNano())
	randomNumber := randSource.Int63()%10 + 1
	for range randomNumber {
		elements := generateRandomElements(SIZE)
		require.Equal(t, maximum(elements), maxChunks(elements))
	}
}
