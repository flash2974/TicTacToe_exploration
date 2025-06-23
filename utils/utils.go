package utils

import (
	"fmt"
	"strings"
)

func Show_tree(node *Node, max_depth int) {
	if max_depth == -1 {
		max_depth = INF
	}

	indent := strings.Repeat("\t", node.Depth)
	indent += fmt.Sprintf("val=%d dept=%d -- %v nexturn=%d win=%v MAXIMISING=%v", node.Value, node.Depth, node.Grid, node.NextTurn, node.Win, node.NextTurn == 1)

	fmt.Println(indent)

	if node.Depth <= max_depth {
		for _, child := range node.Childs {
			Show_tree(child, max_depth)
		}
	}
}

func Show_Branch0(node *Node) {
	indent := strings.Repeat("\t", node.Depth)
	indent += fmt.Sprintf("val=%d dept=%d -- %v nextturn=%d win=%v", node.Value, node.Depth, node.Grid, node.NextTurn, node.Win)

	fmt.Println(indent)

	if len(node.Childs) != 0 {
		Show_Branch0(node.Childs[0])
	}
}

func Show_Depth(node *Node, depth int) {
	if node.Depth+1 < depth {
		Show_Depth(node.Childs[0], depth)
	} else {
		return
	}
	indent := strings.Repeat("\t", node.Depth)
	indent += fmt.Sprintf("val=%d dept=%d -- %v nextturn=%d win=%v", node.Value, node.Depth, node.Grid, node.NextTurn, node.Win)

	fmt.Println(indent)

	for _, child := range node.Childs {
		Show_Depth(child, depth)
	}
}

func GetElement(grid *Grid, row int, col int) int {
	return grid[Array_to_idx(row, col)]
}

func SetElement(grid *Grid, row int, col int, value int) {
	grid[Array_to_idx(row, col)] = value
}

func IdxToArray(index int) []int {
	col := index % Size
	row := (int)(index-col) / Size

	return []int{row, col}
}

func Array_to_idx(row int, col int) int {
	return Size*row + col
}

func GameFinished(grid Grid) int {
	directions := [][]int{
		{0, 1},  // → droite
		{1, 0},  // ↓ bas
		{1, 1},  // ↘ diagonale
		{1, -1}, // ↙ diagonale inversée
	}

	if GridFull(grid) {
		return 0
	}

	for row := range Size {
		for col := range Size {
			startVal := GetElement(&grid, row, col)

			if startVal == 0 {
				continue
			}

			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				count := 1

				nx, ny := row+dx, col+dy

				for count < AlignToWin &&
					nx >= 0 && nx < Size &&
					ny >= 0 && ny < Size &&
					GetElement(&grid, nx, ny) == startVal {

					count++
					nx += dx
					ny += dy
				}

				if count == AlignToWin {
					return startVal
				}
			}
		}
	}
	return -1
}
