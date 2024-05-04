package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/thitiwut-c/slot-machine-cli/app"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: slot <bet_amount>")
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

	app := app.NewDefault()

	app.Run(betAmount)
}
