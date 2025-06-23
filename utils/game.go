package utils

import (
	"fmt"
	"strings"
)

const Size = 4
const AlignToWin = 3
const INF = 1000000

type Grid [Size * Size]int

type Node struct {
	Grid     Grid
	Depth    int
	Value    int
	NextTurn int
	Win      int
	Move     [2]int
	Childs   []*Node
}

func MakeGrid() Grid {
	var grid Grid

	for i := range grid {
		grid[i] = 0
	}

	return grid
}

func ShowGrid(grid Grid) {
	m := map[int]string{0: "₀", 1: "₁", 2: "₂", 3: "₃", 4: "₄", 5: "₅", 6: "₆", 7: "₇", 8: "₈", 9: "₉"}

	for row := range Size {
		// 1ère ligne de la rangée

		fmt.Print("|")
		for col := range Size {
			idx := Array_to_idx(row, col)
			switch grid[idx] {
			case 0:
				fmt.Printf("     | ")
			case 1:
				fmt.Printf("\033[32m \\/  \033[0m| ")
			case 2:
				fmt.Printf("\033[31m /\\  \033[0m| ")
			}
		}
		fmt.Println()

		// 2ème ligne de la rangée
		fmt.Print("|")
		for col := range Size {
			idx := Array_to_idx(row, col)
			switch grid[idx] {
			case 0:
				if idx <= 9 {
					fmt.Printf("   %v | ", m[idx])
				} else {
					tens := idx / 10
					units := idx % 10
					fmt.Printf("   %v%v| ", m[tens], m[units])
				}

			case 1:
				fmt.Printf("\033[32m /\\  \033[0m| ")
			case 2:
				fmt.Printf("\033[31m \\/ \033[0m | ")
			}
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-------", Size)) // séparateur de ligne
	}
	fmt.Println()
}

func FindBestMove(root Node) [2]int {
	bestNode := Node{Value: -999}

	for _, child := range root.Childs {
		if child.Value > bestNode.Value {
			bestNode = *child
		}
	}

	return bestNode.Move
}

func FindPossibleMoves(grid Grid) [][]int {
	var possible_moves [][]int

	for i, element := range grid {
		if element == 0 {
			possible_moves = append(possible_moves, IdxToArray(i))
		}

	}
	return possible_moves
}

func GridFull(grid Grid) bool {
	for _, element := range grid {
		if element == 0 {
			return false
		}
	}
	return true
}
