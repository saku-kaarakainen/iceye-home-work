package hand

import "testing"

func TestIsFourOfKind_OK(t *testing.T) {
	vals := map[rune]int{
		'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9,
		'8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
	}

	testCase := []struct {
		value string
		exRes bool
		exVal int
		act   func(map[rune]int, string) (bool, int)
		mName string
	}{
		{
			value: "22223",
			exRes: true,
			exVal: 2,
			act:   IsFourOfKind,
			mName: "IsFourOfKind",
		},

		{
			value: "22233",
			exRes: false,
			exVal: 0,
			act:   IsFourOfKind,
			mName: "IsFourOfKind",
		},

		{
			value: "22233",
			exRes: true,
			exVal: 2,
			act:   IsFullHouse,
			mName: "IsFullHouse",
		},

		{
			value: "22234",
			exRes: false,
			exVal: 2,
			act:   IsFullHouse,
			mName: "IsFullHouse",
		},

		{
			value: "22234",
			exRes: true,
			exVal: 2,
			act:   IsTripple,
			mName: "IsTripple",
		},

		{
			value: "22334",
			exRes: false,
			exVal: 0,
			act:   IsTripple,
			mName: "IsTripple",
		},

		{
			value: "22334",
			exRes: true,
			exVal: 3,
			act:   IsTwoPairs,
			mName: "IsTwoPairs",
		},

		{
			value: "22345",
			exRes: false,
			exVal: 2,
			act:   IsTwoPairs,
			mName: "IsTwoPairs",
		},

		{
			value: "22345",
			exRes: true,
			exVal: 2,
			act:   IsPair,
			mName: "IsPair",
		},
		{
			value: "23456",
			exRes: false,
			exVal: 0,
			act:   IsPair,
			mName: "IsPair",
		},
		{
			value: "23456",
			exRes: true,
			exVal: 6,
			act:   IsHighCard,
			mName: "IsHighCard",
		},
	}

	for _, tc := range testCase {
		actualResult, actualValue := tc.act(vals, tc.value)

		// Check if the actual result and value match the expected result and value
		if actualResult != tc.exRes || actualValue != tc.exVal {
			t.Errorf(`
Method: %s
	- arg0: cardValues
	- arg1: '%s'

Actual: (%v, %d)
Expected: (%v, %d)

`,
				tc.mName, tc.value,
				actualResult, actualValue,
				tc.exRes, tc.exVal)
		}
	}
}
