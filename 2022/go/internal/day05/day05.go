package day05

import (
	"strconv"
	"strings"
)

func decomposeOrder(movementOrder string) (quantity, source, destination int) {
	items := strings.Split(movementOrder, " ")

	quantity, _ = strconv.Atoi(items[1])
	source, _ = strconv.Atoi(items[3])
	destination, _ = strconv.Atoi(items[5])

	return quantity, source, destination
}

func popCrate(crateStack *[]string) string {
	numberOfCrates := len(*crateStack)
	topMostCratePosition := numberOfCrates - 1
	topMostCrate := (*crateStack)[topMostCratePosition]
	*crateStack = (*crateStack)[:topMostCratePosition]
	return topMostCrate
}

func pushCrate(crate string, crateStack *[]string) {
	*crateStack = append(*crateStack, crate)
}

func popCrates(quantity int, crateStack *[]string) []string {
	numberOfCrates := len(*crateStack)
	topMostCrates := (*crateStack)[numberOfCrates-quantity : numberOfCrates]
	*crateStack = (*crateStack)[:numberOfCrates-quantity]
	return topMostCrates
}

func pushCrates(crates *[]string, crateStack *[]string) {
	*crateStack = append(*crateStack, *crates...)
}

func decomposeManifest(manifest *[]string) *[][]string {
	// Pad the end of the 1st row with one space so we can reason in terms of
	// '[C] ' (that is, all stacks have 4 characters) even in the last stack
	numberOfStacks := (len((*manifest)[0]) + 1) / 4
	maximumStackHeight := len(*manifest)
	stacks := make([][]string, numberOfStacks)
	// Go through the manifest from bottom to top:
	// 0: '....[D]....'
	// 1: '[N].[C]....'
	// 2: '[Z].[M].[P]'
	for height := maximumStackHeight; height > 0; height-- {
		line := (*manifest)[height-1]
		// Pad the end of the line as described above
		line += " "
		for stackNumber := 0; stackNumber < numberOfStacks; stackNumber++ {
			// Each stack has 4 characters: [, the crate letter, ] and a space.
			// '012345678901'
			// '[A].[B].[C].' (spaces replaced with dots for better understanding)
			// So, to get the 1st stack's crate number we need to get the character at position (0*4)+1 = 1 ('A').
			// To get the 2nd stack's crate number we need to get the character at position (1*4)+1 = 5 ('B').
			// To get the 3rd stack's crate number we need to get the character at position (2*4)+1 = 9 ('C').
			crate := line[(stackNumber*4)+1]
			if crate != ' ' {
				pushCrate(string(crate), &stacks[stackNumber])
			}
		}
	}
	return &stacks
}

func readShipArrangement(lines []string) (crates, orders []string) {
	crates = []string{}
	orders = []string{}

	readingCrates := true

	for _, line := range lines {
		if line == "" {
			if readingCrates {
				// First time we encounter an empty line, so it is between the crates and the orders.
				readingCrates = false
				continue
			} else {
				// Since we already read the crates, then this is the second time we encounter an empty line, so this is the eof.
				break
			}
		}

		if readingCrates {
			crates = append(crates, line)
		} else {
			orders = append(orders, line)
		}
	}

	// The list line of crates has the crates numbers, so ignore those
	crates = crates[:len(crates)-1]

	return crates, orders
}

func PartOne(lines []string) string {
	crates, orders := readShipArrangement(lines)

	shipHold := *decomposeManifest(&crates)

	for _, order := range orders {
		quantity, source, destination := decomposeOrder(order)

		for i := 0; i < quantity; i++ {
			crate := popCrate(&shipHold[source-1])
			pushCrate(crate, &shipHold[destination-1])
		}
	}

	var stack []string

	result := new(strings.Builder)
	result.Grow(len(shipHold))
	for stackNumber := 0; stackNumber < len(shipHold); stackNumber++ {
		stack = shipHold[stackNumber]
		positionOfTopCrate := len(stack) - 1
		result.WriteString(stack[positionOfTopCrate])
	}

	return result.String()
}

func PartTwo(lines []string) string {
	stacks, orders := readShipArrangement(lines)

	shipHold := *decomposeManifest(&stacks)

	for _, order := range orders {
		quantity, source, destination := decomposeOrder(order)
		crates := popCrates(quantity, &shipHold[source-1])
		pushCrates(&crates, &shipHold[destination-1])
	}

	var stack []string

	result := new(strings.Builder)
	result.Grow(len(shipHold))
	for stackNumber := 0; stackNumber < len(shipHold); stackNumber++ {
		stack = shipHold[stackNumber]
		positionOfTopCrate := len(stack) - 1
		result.WriteString(stack[positionOfTopCrate])
	}

	return result.String()
}
