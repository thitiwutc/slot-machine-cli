package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thitiwut-c/slot-machine-cli/app"
)

func main() {
	reelCountFlag := flag.Int("c", 3, "Set the reel count, possible values are 3, 5, and 7")
	helpFlag := flag.Bool("h", false, "Display the help message")

	flag.Usage = func() {
		fmt.Println("Usage: slot [options]")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *reelCountFlag != 3 && *reelCountFlag != 5 && *reelCountFlag != 7 {
		flag.Usage()
		os.Exit(1)
	}

	app := app.New(&app.Config{
		ReelCount: *reelCountFlag,
	})

	app.Run()
}
