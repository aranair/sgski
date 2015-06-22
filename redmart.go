package main

import (
	"bufio"
	"os"
	"strconv"

	s "strings"
)

type Node struct {
	xAxis  int32
	yAxis  int32
	Value  int32
	Trails []Node
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
	for r := range m {
		rowStr, _ := reader.ReadString('\n')
		rowArr := s.Split(trimLine(rowStr), " ")
		m[r] = make([]Node, colCount)
		for c := range rowArr {
			m[r][c] = NewNode(rowArr[c], r, c)
		}
	}

	for r := range m {
		for c := range m[r] {
			m[r][c].fillNode()
		}
	}
}

func (n *Node) fillNode(m [][]Node) {
	// top
	if n.xAxis != 0 {
	}
	// bottom
	if n.xAxis == len(m) {
	}

	// left
	if n.yAxis == 0 {
	}
	// right
	if n.yAxis == len(m[0])-1 {
	}
}

func NewNode(s string, x, y int32) Node {
	intVal, _ := strconv.Atoi(s)
	return Node{
		Value: int32(intVal),
		xAxis: x,
		yAxis: y,
	}
}

func trimLine(str string) string {
	return s.Replace(str, "\n", "", -1)
}
