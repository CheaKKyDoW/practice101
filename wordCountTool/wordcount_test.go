package wordcounttool

import "testing"

func TestCountValidWords(t *testing.T) {
	testsCase := []struct {
		input    string
		expected int
	}{
		{"This is Form16 submis$ion date", 3},
		{"apple orange Banana Kiwi123", 4},
		{"a an is it at", 0},
		{"cat d#g m@use do%g", 1},
		{"123 456 7890 $$$ @@@", 0},
		{"aaa eee ooo iii uuu", 0},
		{"bcd fgh xyz mnq", 0},
		{"Form123 name23 g8a9z", 3},
		{"abc defg $%# 23a aa3 1a2e", 2},
		{"", 0},
		{"Valid", 1},
	}
	for _, tc := range testsCase {
		got := countValidWords(tc.input)
		if got != tc.expected {
			t.Errorf("input: %q, got: %d, want: %d", tc.input, got, tc.expected)
		}
	}
}
