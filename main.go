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

	flag.Usage = func() {
		fmt.Println("Usage: slot [options] <bet amount>")
		fmt.Println()
		fmt.Println("Minimum bet amount: 1")
		fmt.Println("Maximum bet amount: 1,000,000")
		fmt.Println()
		fmt.Println("Prizes:")
		fmt.Println()
		fmt.Println("  One cherry (🍒) wins 2.5x bet amount")
		fmt.Println("  Two watermelons (🍉🍉) wins 3x bet amount")
		fmt.Println("  Three bells (🔔🔔🔔) wins 5x bet amount")
		fmt.Println("  Three diamonds (💎💎💎) wins 10x bet amount")
		fmt.Println("  Three cat faces (🐱🐱🐱) wins 100x bet amount")
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

	if len(os.Args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	betAmount, err := strconv.ParseFloat(os.Args[1], 64)
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

	app.Run(betAmount)
}
