package app

import (
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
	symbols    []rune
	reelStates []rune
}

func (a *App) Run() {
	now := time.Now()

	ch := make(chan *reel)

	for i := range a.reelStates {
		randDur, err := rand.Int(rand.Reader, big.NewInt(50))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		baseDur := time.Duration(i+1) * time.Second
		extraDur := time.Duration(randDur.Int64()) * 10 * time.Millisecond

		reelSpinDur := baseDur + extraDur
		go a.spinReel(ch, i, now.Add(reelSpinDur), i == len(a.reelStates)-1)
	}

	var (
		output     string
		prevOutput string
	)

LOOP:
	for {
		select {
		case reel, ok := <-ch:
			if !ok {
				break LOOP
			}

			a.reelStates[reel.idx] = reel.symbol

			for i, symbol := range a.reelStates {
				if symbol == 0 {
					continue
				}

				if i == len(a.reelStates)-1 {
					output += fmt.Sprintf("%c", symbol)
					break
				}

				output += fmt.Sprintf("%c|", symbol)
			}

			fmt.Print(output)
			time.Sleep(50 * time.Millisecond)

			fmt.Print("\r\033[2K")

			prevOutput = output
			output = ""
		}
	}

	fmt.Println(prevOutput)
}

func (a *App) spinReel(ch chan *reel, idx int, stopTime time.Time, lastReel bool) {
	max := big.NewInt(int64(len(a.symbols)))
	randBigInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbolIdx := randBigInt.Int64()

	// Use the last symbol as a starting position. If exists.
	if a.reelStates[idx] != 0 {
		for i, symbol := range a.reelStates {
			if symbol == a.reelStates[idx] {
				symbolIdx = int64(i)
				break
			}
		}
	}

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

	if lastReel {
		close(ch)
	}
}

// NewDefault returns App with default symbols and 3 reels.
func NewDefault(reelCount int) *App {
	return &App{
		symbols: []rune{
			'🍒',
			'🍋',
			'🍊',
			'🍇',
			'🍉',
			'🐱',
			'🐟',
			'🍀',
			'💎',
			'🔔',
		},
		reelStates: make([]rune, reelCount),
	}
}
