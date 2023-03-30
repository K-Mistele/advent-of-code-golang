package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check for Errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Check args
	if len(os.Args) < 2 {
		panic("Specify a filename!")
	}

	// Open the file
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var segmentsContainingEachOther int = 0
	// Open the file one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split it into halves which represent the sections
		sections := strings.Split(line, ",")
		if len(sections) < 2 {
			panic(fmt.Sprintf("Sections definition is not delimited by a comma!%s", line))
		}

		// Split each section into the end
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")
		if len(section1) != 2 || len(section2) != 2 {
			panic(fmt.Sprintf("Section definition is not delimited by a dash!%s", line))
		}

		// If the first section contains the second in entirety
		section1Start, err := strconv.Atoi(section1[0])
		check(err)
		section2Start, err := strconv.Atoi(section2[0])
		check(err)
		section1End, err := strconv.Atoi(section1[1])
		check(err)
		section2End, err := strconv.Atoi(section2[1])
		check(err)

		rangesOverlap := (section1Start <= section2Start && section2Start <= section1End) ||
			(section1Start <= section2End && section2End <= section1End) ||
			(section2Start <= section1Start && section1Start <= section2End) ||
			(section2Start <= section1End && section1End <= section2End)

		if rangesOverlap {
			segmentsContainingEachOther++

		}
	}

	fmt.Printf("Sections containing each other: %d\n", segmentsContainingEachOther)
}
