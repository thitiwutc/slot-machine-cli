package app

type prizeCalculator func(betAmount float64, symbols [3]rune) float64

type prizes struct {
	matched1 []prizeCalculator
	matched2 []prizeCalculator
	matched3 []prizeCalculator
}

var defaultMatched1PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol == '🍒' {
				return betAmount * 2.5
			}
		}

		return 0
	},
}

var defaultMatched2PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		var symbolCount int

		for _, symbol := range symbols {
			if symbol == '🍉' {
				symbolCount++
			}
		}

		if symbolCount == 2 {
			return betAmount * 3
		}

		return 0
	},
}

var defaultMatched3PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol != '🔔' {
				return 0
			}
		}

		return betAmount * 5
	},
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol != '💎' {
				return 0
			}
		}

		return betAmount * 10
	},
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol != '🐱' {
				return 0
			}
		}

		return betAmount * 100
	},
}
