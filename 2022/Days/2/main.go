// Package day2 contains the solutions to the puzzles for Day 2 of the
// Advent of Code 2022.
package day2

import (
	"fmt"
	"strings"

	"github.com/techdufus/AdventOfCode/2022/helpers"
)

const (
	inputFile = "./Days/2/input.txt"
	gameWin   = 6
	gameTie   = 3
	gameLoss  = 0
)

var (
	// opponentWinningMoves key is the move the opponent made and the value is the playerMove
	// that it beats.
	opponentWinningMoves = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	opponentLosesTo = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	// opponentMoves key is the move the opponent made and the value is it's points
	opponentMovePoints = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	opponentTiesTo = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	// playerWinningMoves key is the move the player made and the value is the opponentMove
	// that it beats.
	playerWinningMoves = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	playerLosesTo = map[string]string{
		"X": "B",
		"Y": "C",
		"Z": "A",
	}
	// playerMoves key is the move the player made and the value is it's points
	playerMovePoints = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	desiredState = map[string]string{
		"X": "Lose",
		"Y": "Tie",
		"Z": "Win",
	}
)

func getMoves(file string) [][]string {
	moves := [][]string{}
	for _, line := range helpers.GetInput(file) {
		newLine := strings.TrimSpace(line)
		move := strings.Split(newLine, " ")
		moves = append(moves, move)
	}
	return moves
}

func getDesiredMoves(file string) [][]string {
	moves := [][]string{}
	for _, line := range helpers.GetInput(file) {
		newLine := strings.TrimSpace(line)
		move := strings.Split(newLine, " ")
		newMove := desiredMove(move[0], move[1])
		moves = append(moves, newMove)
	}
	return moves
}

func desiredMove(opponentMove string, desiredMove string) []string {
	var newMove string
	switch desiredState[desiredMove] {
	case "Tie":
		newMove = opponentTiesTo[opponentMove]
	case "Win":
		newMove = opponentLosesTo[opponentMove]
	case "Lose":
		newMove = opponentWinningMoves[opponentMove]
	}

	return []string{opponentMove, newMove}
}

func getWinner(opponentMove string, playerMove string) string {
	if opponentMovePoints[opponentMove] == playerMovePoints[playerMove] {
		return "Tie"
	}
	if opponentMove == playerWinningMoves[playerMove] {
		return "Player"
	}
	return "Opponent"
}

func getPoints(opponentMove string, playerMove string) int {
	points := playerMovePoints[playerMove]
	winner := getWinner(opponentMove, playerMove)
	switch winner {
	case "Tie":
		points += gameTie
	case "Player":
		points += gameWin
	case "Opponent":
		points += gameLoss
	}
	return points
}

func getPartOne(file string) int {
	moves := getMoves(file)
	var points int
	for _, move := range moves {
		points += getPoints(move[0], move[1])
	}
	return points
}

func getPartTwo(file string) int {
	moves := getDesiredMoves(file)
	var points int
	for _, move := range moves {
		points += getPoints(move[0], move[1])
	}
	return points
}

func Answers() {
	timer := helpers.StartTimer()
	fmt.Println("--- Day 2 Answers ---")
	fmt.Println("    Part 1:", getPartOne(inputFile))
	fmt.Println("    Part 2:", getPartTwo(inputFile))
	fmt.Println("  Execution Time:", helpers.GetRuntimeMs(timer))
}
