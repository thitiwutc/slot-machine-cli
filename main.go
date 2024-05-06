package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/thitiwut-c/slot-machine-cli/app"
)

func main() {
	helpFlag := flag.Bool("h", false, "Display the help message")
	spinCountFlag := flag.Int("n", 1, "Number of times the slot matchine spins")

	flag.Usage = func() {
		fmt.Println("Usage: slot [options] <bet amount>")
		fmt.Println()
		fmt.Println("Minimum bet amount: 1")
		fmt.Println("Maximum bet amount: 1,000,000")
		fmt.Println()
		fmt.Println("Prizes:")
		fmt.Println()
		fmt.Println("  One cherry (🍒) wins 1.5x bet amount")
		fmt.Println("  Two watermelons (🍉🍉) wins 3x bet amount")
		fmt.Println("  A cat and a fish (🐱🐟) wins 4x bet amount")
		fmt.Println("  A combination of 3 fruits and vegetables (🍒🍋🍊🍇🍉) wins 5x bet amount")
		fmt.Println("  Three bells (🔔🔔🔔) wins 10x bet amount")
		fmt.Println("  Three diamonds (💎💎💎) wins 30x bet amount")
		fmt.Println("  Three cat (🐱🐱🐱) wins 100x bet amount")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println()
		flag.PrintDefaults()
	}

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	arg := flag.Arg(0)

	betAmount, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	if betAmount < 1 || betAmount > 1_000_000 {
		flag.Usage()
		os.Exit(1)
	}

	app := app.NewDefault()

	for i := 0; i < *spinCountFlag; i++ {
		app.Run(betAmount)
	}
}
