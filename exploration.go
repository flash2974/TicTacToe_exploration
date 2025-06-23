package main

import (
	"exploration/algo"
	"exploration/utils"
	"fmt"
	"time"
)

func benchmark_ab() {
	var root utils.Node
	grid := utils.MakeGrid()
	root.Grid = grid
	root.NextTurn = 1
	root.Depth = 0
	start := time.Now()
	root.Value = algo.MiniMax_AlphaBeta(&root, -utils.INF, utils.INF)
	elapsed := time.Since(start)

	fmt.Printf("temps ab : %v\n", elapsed)
}

func benchmark_normal() {
	var root utils.Node
	grid := utils.MakeGrid()
	root.Grid = grid
	root.NextTurn = 1
	root.Depth = 0
	start := time.Now()
	root.Value = algo.MiniMax(&root)
	elapsed := time.Since(start)

	fmt.Printf("temps normal : %v\n", elapsed)
}

func exploration(grid utils.Grid) utils.Node {
	var root utils.Node
	root.Grid = grid
	root.NextTurn = 1
	root.Depth = 0
	root.Value = algo.MiniMax_AlphaBeta(&root, -utils.INF, utils.INF)
	// root.Value = algo.NegaMax(&root, -1)
	return root
}

func play(turn int) {
	grid := utils.MakeGrid()

	for {
		utils.ShowGrid(grid)
		status := utils.GameFinished(grid)

		if status > 0 {
			fmt.Printf("Le joueur %d a gagné !\n", status)
			break
		} else if status == 0 && utils.GridFull(grid) {
			fmt.Println("Match nul !")
			break
		}
		

		if turn == 1 {
			root := exploration(grid)
			best_move := utils.FindBestMove(root)
			utils.SetElement(&grid, best_move[0], best_move[1], 1)

		} else {
			var num int
			valid := false

			for !valid {
				fmt.Print("Entrez un nombre : ")
				fmt.Scanln(&num)

				if num < utils.Size*utils.Size && num >= 0 && grid[num] == 0 {
					valid = true
				}
				if !valid {
					fmt.Println("Impossible. réessayez")
				}

			}
			row_col := utils.IdxToArray(num)
			utils.SetElement(&grid, row_col[0], row_col[1], 2)
		}

		turn = 3 - turn
	}

	utils.ShowGrid(grid)
}

// bot : joue les 1 (x)
func main() {
	play(1)
	// benchmark_ab()
	// benchmark_normal()

}
