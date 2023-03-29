package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Get the priority from an item character, where a is 1 and z is 26, A is 27 and Z is 52
func getPriorityFromItem(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 96)
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 38)
	} else {
		panic(fmt.Sprintf("Invalid character: %c", char))
	}
}

// Separate a rucksack into two compartments
func getSeparateRuckSackCompartments(line string) (string, string) {
	// split the string into two strings
	strLen := len(line)
	if (strLen % 2) != 0 {
		panic(fmt.Sprintf("Odd-length rucksack definition: %s", line))
	}

	return line[0 : strLen/2], line[strLen/2:]
}

// Find the item that appears in each
func findItemInAllThree(firstRucksack string, secondRucksack string, thirdRucksack string) (itemFoundInAll rune) {

	// map used to track number of unique rucksacks that a given rune appears in
	itemCountMap := make(map[rune]int)

	for _, char := range firstRucksack {
		if itemCountMap[char] == 1 {
			continue
		} // don't count it twice if it already appears in this half
		itemCountMap[char]++
	}

	// Create a map to track runes that unique appear in this rucksack, and then track
	secondRucksackMap := make(map[rune]int)
	for _, char := range secondRucksack {
		// If we have seen it in the seconond rucksack, skip
		if secondRucksackMap[char] == 1 {
			continue
		}

		// If we haven't seen it in this rucksack yet, add it to this rucksack's map and the global count
		secondRucksackMap[char]++
		itemCountMap[char]++
	}

	// If we have seen a rune in two rucksacks already, we've found it
	for _, char := range thirdRucksack {
		if itemCountMap[char] == 2 {
			return char
		}
	}
	panic(fmt.Sprintf("No item found in all rucksacks: %s | %s | %s", firstRucksack, secondRucksack, thirdRucksack))
}

func main() {

	// get the first command-line argument (the filename) and open it
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prioritySum int = 0
	var canScan bool = true
	// get three lines at a time NOTE assumes line count will be a multiple of 3
	for canScan {
		canScan = scanner.Scan()
		firstLine := scanner.Text()
		canScan = scanner.Scan()
		secondLine := scanner.Text()
		canScan = scanner.Scan()
		thirdLine := scanner.Text()
		fmt.Printf("%s | %s | %s\n", firstLine, secondLine, thirdLine)
		if len(firstLine) == 0 || len(secondLine) == 0 || len(thirdLine) == 0 {
			fmt.Printf("Empty line found, breaking\n")
			break
		}

		itemInAllThree := findItemInAllThree(firstLine, secondLine, thirdLine)
		prioritySum += getPriorityFromItem(itemInAllThree)
	}

	fmt.Printf("Sum of priorities: %d\n", prioritySum)

}
