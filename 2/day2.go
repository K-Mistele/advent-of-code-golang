package main

import (
	"bufio"
	"fmt"
	"os"
)

// Play - Enum-like structure for rock/paper/scissors
type Play int

const (
	ROCK Play = iota
	PAPER
	SCISSORS
)

// Outcome - Enum-like structure for the outcome of a game of rock/paper/scissors
type Outcome int

const (
	YOU_WON Outcome = iota
	THEY_WON
	DRAW
)

// Error handler function
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Get the winner of a game of rock/paper/scissors given their choice and your choice
func getWinner(theirPlay Play, yourPlay Play) Outcome {

	switch theirPlay {

	// If they chose Rock
	case ROCK:

		// Handle your choices
		switch yourPlay {
		case ROCK:
			return DRAW
		case PAPER:
			return YOU_WON
		case SCISSORS:
			return THEY_WON
		}
		break

	// If they chose Paper
	case PAPER:
		switch yourPlay {
		case ROCK:
			return THEY_WON
		case PAPER:
			return DRAW
		case SCISSORS:
			return YOU_WON
		}
		break

	// If they chose Scissors
	case SCISSORS:
		switch yourPlay {
		case ROCK:
			return YOU_WON
		case PAPER:
			return THEY_WON
		case SCISSORS:
			return DRAW

		}

	}
	// If we still don't have a valid result, panic
	panic(fmt.Sprintf("Invalid Play: %d / %d", theirPlay, yourPlay))
}

// Get the score for a round
func scoreForRound(yourPlay Play, outcome Outcome) (totalPoints int) {
	totalPoints = 0
	switch yourPlay {
	case ROCK:
		totalPoints += 1
		break
	case PAPER:
		totalPoints += 2
		break

	case SCISSORS:
		totalPoints += 3
		break

	default: // should never happen
		panic(fmt.Sprintf("Invalid Play: %d", yourPlay))
	}
	switch outcome {
	case YOU_WON:
		totalPoints += 6
		break
	case THEY_WON:
		break
	case DRAW:
		totalPoints += 3
		break
	default: // should never happen
		panic(fmt.Sprintf("Invalid Outcome: %d", outcome))
	}
	return totalPoints
}

// Get convert a character to a Play
func playForChar(char string) Play {
	switch char {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	default:
		panic(fmt.Sprintf("Invalid character: %s", char))
	}
}

// Get the outcome the round needs to have per the strategy guide
func outcomeForChar(char string) Outcome {
	switch char {
	case "X":
		return THEY_WON
	case "Y":
		return DRAW
	case "Z":
		return YOU_WON
	default:
		panic(fmt.Sprintf("Invalid character: %s", char))
	}
}

// Given their play and the desired outcome, return your play
func getYourPlayForDesiredOutcome(theirPlay Play, desiredOutcome Outcome) Play {
	switch theirPlay {
	case ROCK:
		switch desiredOutcome {
		case THEY_WON:
			return SCISSORS
		case DRAW:
			return ROCK
		case YOU_WON:
			return PAPER
		}
	case PAPER:
		switch desiredOutcome {
		case THEY_WON:
			return ROCK
		case DRAW:
			return PAPER
		case YOU_WON:
			return SCISSORS
		}
	case SCISSORS:
		switch desiredOutcome {
		case THEY_WON:
			return PAPER
		case DRAW:
			return SCISSORS
		case YOU_WON:
			return ROCK
		}
	}
	panic(fmt.Sprintf("Invalid Play: %d / Outcome %d", theirPlay, desiredOutcome))
}

// Get their play and the desired outcome for the char
func getPlaysForLine(line string) (theirPlay Play, desiredOutcome Outcome) {
	theirPlay = playForChar(string(line[0]))
	desiredOutcome = outcomeForChar(string(line[2]))
	return
}

func main() {
	var totalScore int = 0
	var filename string

	fmt.Printf("ROCK:%d, PAPER:%d, SCISSORS:%d\n", ROCK, PAPER, SCISSORS)
	fmt.Printf("YOU_WIN: %d, THEY_WIN: %d, DRAW: %d\n", YOU_WON, THEY_WON, DRAW)

	// get the file name as the first positional argument
	filename = os.Args[1]
	fmt.Printf("Received filename %s\n", filename)

	// try to open the file, and check the error using check()
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			fmt.Println("Skipping empty line!")
			continue
		}
		fmt.Printf("Line: %s\n", line)
		theirPlay, desiredOutcome := getPlaysForLine(line)
		yourPlay := getYourPlayForDesiredOutcome(theirPlay, desiredOutcome)
		outcome := getWinner(theirPlay, yourPlay)
		totalScore += scoreForRound(yourPlay, outcome)
	}
	fmt.Println("Total Score: ", totalScore)
}
