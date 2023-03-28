package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	/** read an argument from the command line and save it in a variable called filename */
	var filename string
	var currentCalorieCount int = 0
	var elfCalorieTotals []int
	filename = os.Args[1]
	fmt.Printf("Received filename %s\n", filename)

	// Open the file
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	fmt.Println("Opened the file!")

	// Get the file lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// trim whitespace
		trimmed := strings.TrimSpace(strings.ReplaceAll(scanner.Text(), "\n", ""))
		// If we have an empty line, add the elf's calorie count to the list and reset the current calorie count
		if len(trimmed) == 0 {
			elfCalorieTotals = append(elfCalorieTotals, currentCalorieCount)
			currentCalorieCount = 0
		} else {
			// Add the current line's calorie count to the current calorie count
			integer, err := strconv.Atoi(trimmed)
			check(err)
			currentCalorieCount += integer
		}
	}
	// Do it one more time for EOF
	elfCalorieTotals = append(elfCalorieTotals, currentCalorieCount)

	check(scanner.Err())

	// iterate over elf totals and print them
	for idx, calorieCount := range elfCalorieTotals {
		fmt.Printf("Elf #%d total: %d\n", idx+1, calorieCount)
	}

	// Find the elf with the most calories
	var maxCalories int = 0
	var maxCaloriesIdx int = 0
	for idx, calorieCount := range elfCalorieTotals {
		if calorieCount > maxCalories {
			maxCalories = calorieCount
			maxCaloriesIdx = idx
		}
	}

	fmt.Printf("Elf #%d had the most calories with %d\n", maxCaloriesIdx+1, maxCalories)

	// find the sum of the calories for the 3 elves with the most calories
	var topThreeCalories [3]int
	var topThreeCaloriesIdx [3]int
	for idx, calorieCount := range elfCalorieTotals {
		if calorieCount > topThreeCalories[0] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = topThreeCalories[0]
			topThreeCalories[0] = calorieCount
			topThreeCaloriesIdx[2] = topThreeCaloriesIdx[1]
			topThreeCaloriesIdx[1] = topThreeCaloriesIdx[0]
			topThreeCaloriesIdx[0] = idx
		} else if calorieCount > topThreeCalories[1] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = calorieCount
			topThreeCaloriesIdx[2] = topThreeCaloriesIdx[1]
			topThreeCaloriesIdx[1] = idx
		} else if calorieCount > topThreeCalories[2] {
			topThreeCalories[2] = calorieCount
			topThreeCaloriesIdx[2] = idx
		}
	}

	// Find the sum of the calories of the top 3 elves with the most calories

	// Print the top 3 elves
	fmt.Println(topThreeCalories, topThreeCaloriesIdx)
	var sumOfTopThree int = 0
	for _, topCalorieAmount := range topThreeCalories {
		sumOfTopThree += topCalorieAmount
	}
	fmt.Println("the sum of the top 3 calories is ", sumOfTopThree)
}
