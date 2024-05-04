package app

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"
)

type App struct {
	symbols        []rune
	mutex          sync.RWMutex
	currentSymbols [3]rune
}

func (a *App) Run() {
	spinDur := 3500 * time.Millisecond
	now := time.Now()
	stopTime := now.Add(spinDur)

	for i := range a.currentSymbols {
		reelSpinDur := time.Duration(i+1) * time.Second
		go a.spinReel(i, now.Add(reelSpinDur))
	}

	ticker := time.NewTicker(50 * time.Millisecond)

	for tick := range ticker.C {
		a.mutex.RLock()
		var output string
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
		a.mutex.RUnlock()

		fmt.Println(output)
		time.Sleep(50 * time.Millisecond)
		if tick.Before(stopTime) {
			fmt.Print("\033[A")
			fmt.Print("\033[2K")
		} else {
			break
		}
	}
}

func (a *App) spinReel(idx int, stopAt time.Time) {
	max := big.NewInt(int64(len(a.symbols)))
	randBigInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startIdx := randBigInt.Int64()

	for time.Now().Before(stopAt) {
		a.mutex.Lock()
		a.currentSymbols[idx] = a.symbols[startIdx]
		a.mutex.Unlock()
		time.Sleep(50 * time.Millisecond)

		if startIdx == int64(len(a.symbols))-1 {
			startIdx = 0
		} else {
			startIdx++
		}
	}
}

// NewDefault returns App with default symbols (🍒, 🍋, 🍊, 🍇, 🍉, 🍕, 🍀, 💎, and 🔔)
func NewDefault() *App {
	return &App{
		symbols: []rune{
			'🍒',
			'🍋',
			'🍊',
			'🍇',
			'🍉',
			'🍕',
			'🍀',
			'💎',
			'🔔',
		},
	}
}

// New returns App with the given symbols.
func New(symbols []rune) *App {
	return &App{
		symbols: symbols,
	}
}
