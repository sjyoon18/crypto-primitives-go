package aesviz

import "fmt"

func printTitle(title string) {
	fmt.Println(title)
	for i := 0; i < len(title); i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func printSection(title string) {
	fmt.Printf("\n%s\n", title)
}

func printRoundHeader(round int) {
	fmt.Printf("\n< Round %d >\n", round)
}

func printByteMatrix(
	values [4][4]byte,
	format string,
) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Printf(format, values[r][c])
		}
		fmt.Println()
	}
}

func printUint8Matrix(
	values [4][4]uint8,
	format string,
) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Printf(format, values[r][c])
		}
		fmt.Println()
	}
}

func printBoolMatrix(
	values [4][4]bool,
	trueSymbol string,
	falseSymbol string,
) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if values[r][c] {
				fmt.Print(trueSymbol)
			} else {
				fmt.Print(falseSymbol)
			}
		}
		fmt.Println()
	}
}
