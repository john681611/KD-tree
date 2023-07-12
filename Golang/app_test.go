package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedians(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input []int
		output float64
	}{
		{[]int{1} ,1},
		{[]int{1,2,3} ,2},
		{[]int{1,2} ,1.5},
		{[]int{1,2,3,4,5} ,3},
		{[]int{1,2,3,4,5,6} ,3.5},
		{[]int{1,3,1,4,6,5} ,3.5},
	}
	for _, test := range tests {
		test := test 
		t.Run(fmt.Sprintf("Test Median %v", test.output), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.output, GetMedian(test.input))
		})
	}
}


func TestConstructKDTreeLeaf(t *testing.T) {
	points := [][]int{{0,0}}
	got:= kdtree(points, 0)

	assert.Equal(t, []int{0,0}, got.position)
	assert.True(t, got.isLeaf)
}

func TestConstructKDTree(t *testing.T) {
	points := [][]int{{0,0}, {10,10}, {20, 20}}
	got:= kdtree(points, 0)

	assert.Equal(t, float64(10), got.location)
	assert.False(t, got.isLeaf)

	assert.False(t, got.leftChild.isLeaf)
	assert.Equal(t, float64(5), got.leftChild.location)

	assert.True(t, got.rightChild.isLeaf)
	assert.Equal(t, []int{20,20}, got.rightChild.position)

	assert.True(t, got.leftChild.leftChild.isLeaf)
	assert.Equal(t, []int{0,0}, got.leftChild.leftChild.position)

	assert.True(t, got.leftChild.rightChild.isLeaf)
	assert.Equal(t, []int{10,10}, got.leftChild.rightChild.position)
}