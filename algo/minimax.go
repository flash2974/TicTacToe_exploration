package algo

import (
	"exploration/utils"
)

func MiniMax(node *utils.Node) int {
	result := utils.GameFinished(node.Grid)
	if result != -1 {
		switch result {
		case 0:
			return 0 // match nul
		case 1:
			return 1 // bot gagne
		case 2:
			return -1 // humain gagne
		}
	}

	isMaximising := node.NextTurn == 1

	bestValue := -utils.INF
	if !isMaximising {
		bestValue = utils.INF
	}

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

		score := MiniMax(child)

		if isMaximising {
			bestValue = max(bestValue, score)
		} else {
			bestValue = min(bestValue, score)
		}

		child.Value = score
		node.Childs = append(node.Childs, child)
	}
	return bestValue
}

func MiniMax_AlphaBeta(node *utils.Node, a int, b int) int {
	result := utils.GameFinished(node.Grid)
	if result != -1 {
		switch result {
		case 0:
			return 0 // match nul
		case 1:
			return 1 // bot gagne
		case 2:
			return -1 // humain gagne
		}
	}

	isMaximising := node.NextTurn == 1

	bestValue := -utils.INF
	if !isMaximising {
		bestValue = utils.INF
	}

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

		score := MiniMax_AlphaBeta(child, a, b)

		if isMaximising {
			bestValue = max(bestValue, score)

			if bestValue >= b {
				break
			}
			a = max(a, bestValue)
		} else {
			bestValue = min(bestValue, score)

			if bestValue <= a {
				break
			}
			b = min(b, bestValue)
		}

		child.Value = score
		node.Childs = append(node.Childs, child)
	}
	return bestValue
}
