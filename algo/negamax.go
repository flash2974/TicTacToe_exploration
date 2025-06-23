package algo

import (
	"exploration/utils"
)

func NegaMax_AlphaBeta(node *utils.Node, a int, b int, color int) int {
	/**
	color = 1 -> joueur veut maximiser
	coor = -1 -> minimiser
	*/
	result := utils.GameFinished(node.Grid)
	if result != -1 {
		switch result {
		case 0:
			return 0 // match nul
		case 1:
			return -color // bot gagne
		case 2:
			return color // humain gagne
		}
	}

	bestValue := -utils.INF

	moves := utils.FindPossibleMoves(node.Grid)
	for _, move := range moves {
		newGrid := node.Grid
		utils.SetElement(&newGrid, move[0], move[1], node.NextTurn)

		child := &utils.Node{
			Grid:     newGrid,
			NextTurn: 3 - node.NextTurn,
			Depth:    node.Depth + 1,
			Move:     [2]int{move[0], move[1]},
		}

		score := -NegaMax(child, -color)

		bestValue = max(bestValue, score)

		child.Value = score
		node.Childs = append(node.Childs, child)
	}
	return bestValue
}

func NegaMax(node *utils.Node, color int) int {
	/**
	color = 1 -> joueur veut maximiser
	coor = -1 -> minimiser
	*/
	result := utils.GameFinished(node.Grid)
	if result != -1 {
		switch result {
		case 0:
			return 0 // match nul
		case 1:
			return -color // bot gagne
		case 2:
			return color // humain gagne
		}
	}

	bestValue := -utils.INF

	moves := utils.FindPossibleMoves(node.Grid)
	for _, move := range moves {
		newGrid := node.Grid
		utils.SetElement(&newGrid, move[0], move[1], node.NextTurn)

		child := &utils.Node{
			Grid:     newGrid,
			NextTurn: 3 - node.NextTurn,
			Depth:    node.Depth + 1,
			Move:     [2]int{move[0], move[1]},
		}

		score := -NegaMax(child, -color)

		bestValue = max(bestValue, score)

		child.Value = score
		node.Childs = append(node.Childs, child)
	}
	return bestValue
}
