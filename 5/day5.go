package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []crate
type crate string
type moveCommand struct {
	amount    int
	fromStack int
	toStack   int
}

// Check an error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// parse out a list of stacks and commands from a list of lines
func getStacksAndCommands(lines []string) (map[int]stack, []moveCommand) {
	var stacks map[int]stack = make(map[int]stack)
	// golang create a slice of pointers to moveCommand structs
	var commands []moveCommand = make([]moveCommand, 0)

	readingStacks := true
	for _, line := range lines {

		if readingStacks {
			/**
			Given a layout of lines like this:
			[W] [V]     [P]
			[B] [T]     [C] [B]     [G]
			[G] [S]     [V] [H] [N] [T]
			[Z] [B] [W] [J] [D] [M] [S]
			[R] [C] [N] [N] [F] [W] [C]     [W]
			[D] [F] [S] [M] [L] [T] [L] [Z] [Z]
			[C] [W] [B] [G] [S] [V] [F] [D] [N]
			[V] [G] [C] [Q] [T] [J] [P] [B] [M]
			 1   2   3   4   5   6   7   8   9
			 1   5   9  13  17  21  25  29  33
			parse the columns into stacks
			ll = 15
			ns =1
			newCrate = line[1] = W
			ns = 5
			newCrate = line[5] = V
			ns = 9
			newCrate = line[9] = ' ' (space) -> continue
			ns = 13
			newCrate = line[13] = P
			ns = 17 break
			*/

			curCharIdx := 1
			for curCharIdx+1 < len(line) {
				fmt.Println("New Stack\n-------------")
				fmt.Println("Current char", curCharIdx)

				currentStackId := curCharIdx/4 + 1
				fmt.Println("Current stack idx", currentStackId)
				currentStack := stacks[currentStackId]
				fmt.Println("Current Stack", currentStack)
				newCrate := line[curCharIdx]
				if newCrate == '1' {
					fmt.Println("Reached labels, skipping!")
					break
				} else if newCrate != ' ' {
					fmt.Printf("Adding crate %c to stack %d\n", newCrate, currentStackId)

					// NOTE that this looks really weird because we are conceptualizing a stack as a slice, with the
					//  top of the stack at the back, BUT, we parse the file (and therefore the stacks) from top to bottom
					// 	so, we add crates to the bottom of the stack, which is the front of the slice, as we read
					// 	down the file
					currentStack = append([]crate{crate(newCrate)}, currentStack...)
					stacks[currentStackId] = currentStack // stacks can't be updated weirdly
					fmt.Println("updated stack: ", currentStack)
				} else {
					fmt.Println("Empty crate! continuing")
				}
				curCharIdx += 4
			}

			// When we hit the line break, skip it and change modes
			if line == "" {
				readingStacks = false
				continue
			}
		} else {
			// parse out commands

			// parse out AMOUNT, FROM, and TO amounts from a string structured as "move AMOUNT from FROM to TO"
			tokens := strings.Split(line, " ")
			amount, err := strconv.Atoi(tokens[1])
			check(err)
			from, err := strconv.Atoi(tokens[3])
			check(err)
			to, err := strconv.Atoi(tokens[5])
			check(err)
			commands = append(commands, moveCommand{amount, from, to})
		}

	}
	fmt.Println("commands", commands[0:3])
	return stacks, commands
}

// moveCreatesFromStackToStack moves numCreatesToMove crates from the from stack (with ID fromStackId) to the to stack (with ID toStackId)
// 	Function makes use of pointers to update the map in place, no return value necessary
func moveCratesFromStackToStack(stacks *map[int]stack, fromStackId int, toStackId int, numCratesToMove int) {

	fromStack := (*stacks)[fromStackId]
	toStack := (*stacks)[toStackId]

	// move crates from fromStack to toStack
	for i := 0; i < numCratesToMove; i++ {

		// get the item from the back of the slice
		topItemOnFromStack := fromStack[len(fromStack)-1]

		// Add it to top/back of the destination stack - UNLIKE during stack parsing, we're not adding to the bottom
		toStack = append(toStack, topItemOnFromStack)

		// now, trim it off the back of the source stack
		fromStack = fromStack[:len(fromStack)-1]

	}

	// Commit changes back to the map now that we're done. slices are reference types and the references change when we append
	// 	so, we need to update the references in the map since the old ones are invalid
	(*stacks)[fromStackId] = fromStack
	(*stacks)[toStackId] = toStack
}

func main() {

	// Read in the filename from the command line
	if len(os.Args) < 2 {
		panic("Specify a filename!")
	}

	// Open the file
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	fmt.Println("Got filename", filename)

	// Create a scanner
	scanner := bufio.NewScanner(file)
	check(err)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		check(scanner.Err())
	}

	stacks, commands := getStacksAndCommands(lines)
	fmt.Printf("command %v\n", commands[0])
	fmt.Println("stacks", stacks)

	for _, command := range commands {
		moveCratesFromStackToStack(&stacks, command.fromStack, command.toStack, command.amount)
	}

	fmt.Println(stacks)
}
