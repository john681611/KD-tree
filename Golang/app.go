package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/samber/lo"
)

type Node struct {
	position []int
	location   float64 // the median
	isLeaf bool
	leftChild  *Node
	rightChild *Node
}

// [[x, y], [x2, y2]]
func kdtree(coords [][]int, depth int) Node {
	if len(coords) == 1 {
		return Node{
			position: coords[0],
			isLeaf: true,
		}
	}

	axis := depth % 2 // 0 is x-axis, else y
	
	median := GetAxisMedian(coords, axis)

	fmt.Println(coords, depth, median)
	leftChildNode := kdtree(lo.Filter(coords, func(x []int, index int) bool {
		return float64(x[axis]) <= median
	}), depth+1)

	rightChildNode := kdtree(lo.Filter(coords, func(x []int, index int) bool {
		return float64(x[axis]) > median
	}), depth+1)

	return Node{
		location:      median,
		leftChild:     &leftChildNode,
		rightChild: &rightChildNode,
	}

}

func GetAxisMedian(coords [][]int, axis int) float64 {
	axisArray := lo.Map(coords, func(x []int, index int) int {
		return x[axis]
	})
	return GetMedian(axisArray)
}

func GetMedian(axisArray []int) float64 {
	sort.Ints(axisArray)
	arrayLen := len(axisArray)
	if arrayLen % 2 == 0 {
		halfPoint := arrayLen/2
		return float64(axisArray[halfPoint] + axisArray[halfPoint-1])/2
		} else {
		return float64(axisArray[int(math.Floor(float64(arrayLen)/2))])
	}
}
