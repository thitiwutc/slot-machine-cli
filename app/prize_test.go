package app

import "testing"

func TestPrizes_CalculatePrize(t *testing.T) {
	type args struct {
		betAmount float64
		symbols [3]rune
	}

	testCases := []struct {
		name string
		args args
		expect float64
	}{
		{
			name: "1CherryAtIndex0_Return2.5xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍒', '🍋', '🍀'},
			},
			expect: 2.5,
		},
		{
			name: "1CherryAtIndex1_Return2.5xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍋', '🍒', '🍀'},
			},
			expect: 2.5,
		},
		{
			name: "1CherryAtIndex2_Return2.5xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍀', '🍋', '🍒'},
			},
			expect: 2.5,
		},
		{
			name: "2WatermelonsAtIndex0And1_Return3xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍉', '🍉', '🍀'},
			},
			expect: 3,
		},
		{
			name: "2WatermelonsAtIndex2And3_Return3xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍀', '🍉', '🍉'},
			},
			expect: 3,
		},
		{
			name: "2WatermelonsAtIndex2And3_Return3xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🍉', '🍀', '🍉'},
			},
			expect: 3,
		},
		{
			name: "3Bells_Return5xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🔔', '🔔', '🔔'},
			},
			expect: 5,
		},
		{
			name: "3Diamonds_Return10xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'💎', '💎', '💎'},
			},
			expect: 10,
		},
		{
			name: "3CatFaces_Return100xBetAmount",
			args: args{
				betAmount: 1,
				symbols: [3]rune{'🐱', '🐱', '🐱'},
			},
			expect: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prizes := prizes{
				matched1: defaultMatched1PrizeCaltr,
				matched2: defaultMatched2PrizeCaltr,
				matched3: defaultMatched3PrizeCaltr,
			}

			got := prizes.calculatePrize(tc.args.betAmount, tc.args.symbols)
			if got != tc.expect {
				t.Errorf("expect: %v, got: %v", tc.expect, got)
			}
		})
	}
}
