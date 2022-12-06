package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

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
	for _, line := range getInput(file) {
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

func main() {
	elves := getElves("input.txt")
	mostCalories := getTopCalories(elves)
	topThreeCalories := getTopThreeCalories(elves)
	fmt.Println("Most calories:", mostCalories)
	fmt.Println("Top 3 calories:", topThreeCalories)
	fmt.Println("Top 3 SUM:", getCaloriesSum(topThreeCalories))
}
