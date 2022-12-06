package day1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/techdufus/AdventOfCode/2022/helpers"
)

func getCaloriesSum(elf []int) int {
	total := 0
	for _, c := range elf {
		total += c
	}
	return total
}

func getElves(file string) [][]int {
	elves := [][]int{}
	elf := []int{}
	for _, line := range helpers.GetInput(file) {
		if line == "" {
			elves = append(elves, elf)
			elf = []int{}
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elf = append(elf, calories)
		}
	}
	elves = append(elves, elf)
	return elves
}

func getTopCalories(elves [][]int) int {
	var mostCalories int
	for _, elf := range elves {
		calories := getCaloriesSum(elf)
		if calories > mostCalories {
			mostCalories = calories
		}
	}
	return mostCalories
}

func getTopThreeCalories(elves [][]int) []int {
	var topThree []int
	for _, elf := range elves {
		calories := getCaloriesSum(elf)
		if len(topThree) < 3 {
			topThree = append(topThree, calories)
		} else {
			for i, c := range topThree {
				if calories > c {
					topThree = append(topThree[:i], append([]int{calories}, topThree[i:]...)...)
					topThree = topThree[:3]
					break
				}
			}
		}
	}
	return topThree
}

// Answers prints the answers to the day's puzzles
func Answers() {
	elves := getElves("./Days/1/input.txt")
	mostCalories := getTopCalories(elves)
	topThreeCalories := getTopThreeCalories(elves)
	fmt.Println("--- Day 1 Answers ---")
	fmt.Println("    Part1:", mostCalories)
	fmt.Println("    Part2:", getCaloriesSum(topThreeCalories))
}