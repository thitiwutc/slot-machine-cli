package app

type prizeCalculator func(betAmount float64, symbols [3]rune) float64

var defaultMatched1PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol == '🍒' {
				return betAmount * 1.5
			}
		}

		return 0
	},
}

var defaultMatched2PrizeCaltr = []prizeCalculator{
	func(betAmount float64, symbols [3]rune) float64 {
		var (
			catCount  int
			fishCount int
		)

		for _, symbol := range symbols {
			if symbol == '🐱' {
				catCount++
			}

			if symbol == '🐟' {
				fishCount++
			}
		}

		if catCount >= 1 && fishCount >= 2 {
			return betAmount * 4
		}

		return 0
	},
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
			if symbol != '🐱' {
				return 0
			}
		}

		return betAmount * 100
	},
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol != '💎' {
				return 0
			}
		}

		return betAmount * 30
	},
	func(betAmount float64, symbols [3]rune) float64 {
		for _, symbol := range symbols {
			if symbol != '🔔' {
				return 0
			}
		}

		return betAmount * 10
	},
	func(betAmount float64, symbols [3]rune) float64 {
		var fruitVegCount int

		for _, symbol := range symbols {
			if symbol == '🍒' ||
				symbol == '🍋' ||
				symbol == '🍊' ||
				symbol == '🍇' ||
				symbol == '🍉' {
				fruitVegCount++
			}
		}

		if fruitVegCount == 3 {
			return betAmount * 5
		}

		return 0
	},
}

type prizes struct {
	matched1 []prizeCalculator
	matched2 []prizeCalculator
	matched3 []prizeCalculator
}

func (p prizes) calculatePrize(betAmount float64, symbols [3]rune) float64 {
	for _, prizeCaltr := range p.matched3 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	for _, prizeCaltr := range p.matched2 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	for _, prizeCaltr := range p.matched1 {
		if prize := prizeCaltr(betAmount, symbols); prize > 0 {
			return prize
		}
	}

	return 0
}
