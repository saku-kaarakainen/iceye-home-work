package main

type testCase struct {
	Hand1  string
	Hand2  string
	Winner string
}

func geTestCases() []testCase {
	return []testCase{
		{Hand1: "AAAQQ", Hand2: "QQAAA", Winner: "Tie"},
		{Hand1: "53QQ2", Hand2: "Q53Q2", Winner: "Tie"},
		{Hand1: "53888", Hand2: "88385", Winner: "Tie"},
		{Hand1: "QQAAA", Hand2: "AAAQQ", Winner: "Tie"},
		{Hand1: "Q53Q2", Hand2: "53QQ2", Winner: "Tie"},
		{Hand1: "88385", Hand2: "53888", Winner: "Tie"},
		{Hand1: "AAAQQ", Hand2: "QQQAA", Winner: "Hand 1"},
		{Hand1: "Q53Q4", Hand2: "53QQ2", Winner: "Hand 1"},
		{Hand1: "53888", Hand2: "88375", Winner: "Hand 1"},
		{Hand1: "33337", Hand2: "QQAAA", Winner: "Hand 1"},
		{Hand1: "22333", Hand2: "AAA58", Winner: "Hand 1"},
		{Hand1: "33389", Hand2: "AAKK4", Winner: "Hand 1"},
		{Hand1: "44223", Hand2: "AA892", Winner: "Hand 1"},
		{Hand1: "22456", Hand2: "AKQJT", Winner: "Hand 1"},
		{Hand1: "99977", Hand2: "77799", Winner: "Hand 1"},
		{Hand1: "99922", Hand2: "88866", Winner: "Hand 1"},
		{Hand1: "9922A", Hand2: "9922K", Winner: "Hand 1"},
		{Hand1: "99975", Hand2: "99965", Winner: "Hand 1"},
		{Hand1: "99975", Hand2: "99974", Winner: "Hand 1"},
		{Hand1: "99752", Hand2: "99652", Winner: "Hand 1"},
		{Hand1: "99752", Hand2: "99742", Winner: "Hand 1"},
		{Hand1: "99753", Hand2: "99752", Winner: "Hand 1"},
		{Hand1: "88822", Hand2: "QQ777", Winner: "Hand 1"},
		{Hand1: "99662", Hand2: "88776", Winner: "Hand 1"},
		{Hand1: "QQQAA", Hand2: "AAAQQ", Winner: "Hand 2"},
		{Hand1: "53QQ2", Hand2: "Q53Q4", Winner: "Hand 2"},
		{Hand1: "88375", Hand2: "53888", Winner: "Hand 2"},
		{Hand1: "QQAAA", Hand2: "33337", Winner: "Hand 2"},
		{Hand1: "AAA58", Hand2: "22333", Winner: "Hand 2"},
		{Hand1: "AAKK4", Hand2: "33389", Winner: "Hand 2"},
		{Hand1: "AA892", Hand2: "44223", Winner: "Hand 2"},
		{Hand1: "AKQJT", Hand2: "22456", Winner: "Hand 2"},
		{Hand1: "77799", Hand2: "99977", Winner: "Hand 2"},
		{Hand1: "88866", Hand2: "99922", Winner: "Hand 2"},
		{Hand1: "9922K", Hand2: "9922A", Winner: "Hand 2"},
		{Hand1: "99965", Hand2: "99975", Winner: "Hand 2"},
		{Hand1: "99974", Hand2: "99975", Winner: "Hand 2"},
		{Hand1: "99652", Hand2: "99752", Winner: "Hand 2"},
		{Hand1: "99742", Hand2: "99752", Winner: "Hand 2"},
		{Hand1: "99752", Hand2: "99753", Winner: "Hand 2"},
	}
}
