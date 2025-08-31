package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thitiwut-c/slot-machine-cli/app"
)

func main() {
	helpFlag := flag.Bool("h", false, "Display the help message")
	spinCountFlag := flag.Int("nspin", 1, "Number of times the slot matchine spins")
	reelCount := flag.Int("nreel", 3, "Number of reels to spin")
	symbols := flag.String("symbols", "🍒🍋🍊🍇🍉🐶🐱🦆🦓🐟🍀💎🔔", "Symbol set in reel")

	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage = func() {
		fmt.Println("Usage: slot [options]")
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

	if *spinCountFlag < 0 || *reelCount < 0 {
		flag.Usage()
		os.Exit(1)
	}

	app := app.NewDefault(*symbols, *reelCount)

	for i := 0; i < *spinCountFlag; i++ {
		app.Run()
	}
}
