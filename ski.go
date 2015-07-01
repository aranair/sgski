package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	s "strings"
)

type Node struct {
	xAxis int
	yAxis int
	Value int32
	Trail []int32
}

type NodeMap struct {
	m [][]Node
}

func main() {
	// Read inputs
	f, _ := os.Open("bigmap")
	reader := bufio.NewReader(f)
	metaStr, _ := reader.ReadString('\n')
	meta := s.Split(trimLine(metaStr), " ")
	rowCount, _ := strconv.Atoi(meta[0])
	colCount, _ := strconv.Atoi(meta[1])

	// Construct 2d array of nodes
	m := make([][]Node, rowCount)
	nm := NodeMap{m: m}
	for r := range m {
		rowStr, _ := reader.ReadString('\n')
		rowArr := s.Split(trimLine(rowStr), " ")
		m[r] = make([]Node, colCount)
		for c := range rowArr {
			m[r][c] = NewNode(rowArr[c], r, c)
		}
	}

	// Fill nodes one by one and store the best trail
	var bestTrail []int32
	for r := range m {
		for c := range m[r] {
			t := nm.fillNode(r, c)
			if len(bestTrail) == 0 || isBetterTrail(bestTrail, t) {
				bestTrail = t
			}
		}
	}

	fmt.Println(bestTrail)
	fmt.Println(len(bestTrail))
	fmt.Println(calcDesc(bestTrail))
}

func (nm *NodeMap) goNext(curr *Node, r, c int) []int32 {
	m := nm.m

	// Out of bounds
	if r < 0 || r > len(m)-1 ||
		c < 0 || c > len(m[0])-1 {
		return nil
	}

	next := m[r][c]
	if next.Value < curr.Value {
		if next.Trail == nil {
			return nm.fillNode(r, c)
		} else {
			return next.Trail // already visited
		}
	}
	return nil
}

func (nm *NodeMap) fillNode(r, c int) []int32 {
	m := nm.m
	n := m[r][c]

	trails := make([][]int32, 4)
	trails = append(trails, nm.goNext(&n, r-1, c)) // up
	trails = append(trails, nm.goNext(&n, r+1, c)) // down
	trails = append(trails, nm.goNext(&n, r, c-1)) // left
	trails = append(trails, nm.goNext(&n, r, c+1)) // right

	totalLen := 0
	for _, t := range trails {
		totalLen += len(t)
	}

	// Cache best trail
	if totalLen == 0 {
		// leaf, cache [self.value] and return
		n.Trail = make([]int32, 1)
		n.Trail[0] = n.Value
	} else {
		// non-leaf, Cache best trail and return
		n.Trail = append(findBest(trails), n.Value)
	}

	return n.Trail
}

// Compares 2 trails and see which is better
func isBetterTrail(t1 []int32, t2 []int32) bool {
	return (len(t1) < len(t2) ||
		(len(t1) == len(t2) && calcDesc(t1) < calcDesc(t2)))
}

// Longest Path, or if same length more descend
func findBest(trails [][]int32) (bestTrail []int32) {
	for _, ts := range trails {
		if len(ts) > 0 &&
			(len(bestTrail) == 0 || isBetterTrail(bestTrail, ts)) {
			bestTrail = ts
		}
	}
	return
}

func calcDesc(a []int32) int32 {
	var smallest, biggest int32 = a[0], -1

	for _, v := range a {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	return biggest - smallest
}

func NewNode(s string, x, y int) Node {
	intVal, _ := strconv.Atoi(s)
	return Node{
		Trail: nil,
		Value: int32(intVal),
		xAxis: x,
		yAxis: y,
	}
}

func trimLine(str string) string {
	return s.Replace(str, "\n", "", -1)
}
