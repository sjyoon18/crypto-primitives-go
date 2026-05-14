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
