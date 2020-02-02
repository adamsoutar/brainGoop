package main

type instruction int
const (
	increment instruction = iota
	decrement
	left
	right
	startLoop
	endLoop
	input
	output
)

func charToInstruction (chr int32) instruction {
	// Ranging over a string gives you ascii numbers
	switch chr {
	case 43:
		return increment
	case 45:
		return decrement
	case 60:
		return left
	case 62:
		return right
	case 91:
		return startLoop
	case 93:
		return endLoop
	case 44:
		return input
	case 46:
		return output
	}
	return -1
}