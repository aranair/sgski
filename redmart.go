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
	// 4 4
	// 4 8 7 3
	// 2 5 9 3
	// 6 3 2 5
	// 4 4 1 6

	reader := bufio.NewReader(os.Stdin)
	metaStr, _ := reader.ReadString('\n')
	meta := s.Split(trimLine(metaStr), " ")
	rowCount, _ := strconv.Atoi(meta[0])
	colCount, _ := strconv.Atoi(meta[1])

	// Construct 2d array of nodes
	m := make([][]Node, rowCount)
	nm := NodeMap{m: m}

	// Building Nodes
	for r := range m {
		rowStr, _ := reader.ReadString('\n')
		rowArr := s.Split(trimLine(rowStr), " ")
		m[r] = make([]Node, colCount)
		for c := range rowArr {
			m[r][c] = NewNode(rowArr[c], r, c)
		}
	}

	// Filling Nodes
	var trails [][]int32
	for r := range m {
		for c := range m[r] {
			trails = append(trails, nm.fillNode(r, c))
		}
	}
	fmt.Println(findBest(trails))
}

func (nm *NodeMap) getNextTrail(curr *Node, r, c int) []int32 {
	m := nm.m

	// Out of bounds
	if r < 0 || r > len(m)-1 ||
		c < 0 || c > len(m[0])-1 {
		return nil
	}

	next := m[r][c]

	if next.Value < curr.Value && next.Trail == nil {
		return nm.fillNode(r, c)
	}
	return nil
}

func (nm *NodeMap) fillNode(r, c int) []int32 {
	m := nm.m

	// Out of bounds
	if r < 0 || r > len(m)-1 ||
		c < 0 || c > len(m[0])-1 {
		return nil
	}

	n := m[r][c]

	// Visited
	if n.Trail != nil {
		return n.Trail
	}

	trails := make([][]int32, 4)
	trails = append(trails, nm.getNextTrail(&n, r-1, c)) // up
	trails = append(trails, nm.getNextTrail(&n, r+1, c)) // down
	trails = append(trails, nm.getNextTrail(&n, r, c-1)) // left
	trails = append(trails, nm.getNextTrail(&n, r, c+1)) // right

	totalLen := 0
	for _, t := range trails {
		totalLen += len(t)
	}

	if totalLen == 0 {
		// leaf, return [self.value]
		n.Trail = make([]int32, 1)
		n.Trail[0] = n.Value
	} else {
		// Find best trail from up/down/left/right trails
		n.Trail = append(findBest(trails), n.Value)
	}

	return n.Trail
}

// Longest Path, or if same length more descend
func findBest(trails [][]int32) (bestTrail []int32) {
	for _, ts := range trails {
		if len(ts) > 0 &&
			(len(bestTrail) == 0 || len(bestTrail) < len(ts) ||
				(len(bestTrail) == len(ts) && calcDesc(bestTrail) < calcDesc(ts))) {
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
