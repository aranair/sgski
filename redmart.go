package main

import (
	"bufio"
	"os"
	"strconv"

	s "strings"
)

type Node struct {
	xAxis int32
	yAxis int32
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
			nm.fillNode(r, c)
		}
	}
}

func getTrail(m *[][]Node, r, c int32) {
}

func (nm *NodeMap) fillNode(r, c int32) {
	m := nm.m
	n := m[r][c]

	var trail []int32

	// not top-most
	if n.xAxis > 0 {
		top := m[r-1][c]
		if top.Value < n.Value && top.Trail == nil {
			trails = append(trails, m.fillNode(r-1, c))
		}
	}

	// not bottom-most
	if n.xAxis == len(m) {
		bottom := m[r+1][c]
		if bottom.Value < n.Value && bottom.Trail == nil {
			trails = append(trails, m.fillNode(r+1, c))
		}
	}

	// not left-most
	if n.yAxis == 0 {
		left := m[r][c-1]
		if left.Value < n.Value && left.Trail == nil {
			trails = append(trails, m.fillNode(r, c-1))
		}
	}

	// not right-most
	if n.yAxis == len(m[0])-1 {
		right := m[r][c+1]

		if right.Value < n.Value && right.Trail == nil {
			trails = append(trails, m.fillNode(r, c+1))
		}
	}

	// pick best trail
	return trail
}

func NewNode(s string, x, y int32) Node {
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
