package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"
)

type reel struct {
	idx    int
	symbol rune
}

type App struct {
	symbols        []rune
	reelCount      int
	currentSymbols [3]rune
	prizes         prizes
}

func (a *App) Run(betAmount float64) {
	spinDur := time.Duration(a.reelCount*1000) * time.Millisecond
	now := time.Now()
	stopTime := now.Add(spinDur)

	ch := make(chan *reel, 3)
	defer close(ch)

	for i := range a.currentSymbols {
		reelSpinDur := time.Duration(i+1) * time.Second
		go a.spinReel(ch, i, now.Add(reelSpinDur))
	}

	ctx, cancel := context.WithDeadline(context.Background(), stopTime)
	defer cancel()

	var (
		output     string
		prevOutput string
	)

	for {
		select {
		case reel := <-ch:
			a.currentSymbols[reel.idx] = reel.symbol

			for i, symbol := range a.currentSymbols {
				if symbol == 0 {
					continue
				}

				if i == len(a.currentSymbols)-1 {
					output += fmt.Sprintf("%c", symbol)
					break
				}

				output += fmt.Sprintf("%c|", symbol)
			}

			fmt.Println(output)
			time.Sleep(50 * time.Millisecond)

			fmt.Print("\033[A")
			fmt.Print("\033[2K")

			prevOutput = output
			output = ""
		case <-ctx.Done():
			prize := a.prizes.calculatePrize(betAmount, a.currentSymbols)
			if prize > 0 {
				fmt.Printf("%s You win %.2f\n", prevOutput, prize)
			} else {
				fmt.Printf("%s You lose\n", prevOutput)
			}
			time.Sleep(200 * time.Millisecond)
			return
		}
	}
}

func (a *App) spinReel(ch chan *reel, idx int, stopTime time.Time) {
	max := big.NewInt(int64(len(a.symbols)))
	randBigInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbolIdx := randBigInt.Int64()

	ticker := time.NewTicker(50 * time.Millisecond)

	for tick := range ticker.C {
		if tick.After(stopTime) {
			break
		}

		ch <- &reel{
			idx:    idx,
			symbol: a.symbols[symbolIdx],
		}

		if symbolIdx == int64(len(a.symbols))-1 {
			symbolIdx = 0
		} else {
			symbolIdx++
		}
	}
}

// NewDefault returns App with default symbols (🍒, 🍋, 🍊, 🍇, 🍉, 🍕, 🍀, 💎, and 🔔)
// and 3 reels.
func NewDefault() *App {
	return &App{
		symbols: []rune{
			'🍒',
			'🍋',
			'🍊',
			'🍇',
			'🍉',
			'🐱',
			'🍀',
			'💎',
			'🔔',
		},
		reelCount: 3,
		prizes: prizes{
			matched1: defaultMatched1PrizeCaltr,
			matched2: defaultMatched2PrizeCaltr,
			matched3: defaultMatched3PrizeCaltr,
		},
	}
}
