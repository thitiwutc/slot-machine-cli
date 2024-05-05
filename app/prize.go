package app

type prizeCalculator func(betAmount float64, symbols [3]rune) float64

var defaultMatched1PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol == '🍒' {
				return betAmount * 2
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

type prizes struct {
	matched1 []prizeCalculator
	matched2 []prizeCalculator
	matched3 []prizeCalculator
}

func (p prizes) calculatePrize(betAmount float64, symbols [3]rune) float64 {
	for _, prizeCaltr := range p.matched1 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	for _, prizeCaltr := range p.matched2 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	for _, prizeCaltr := range p.matched3 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	return 0
}
