package main

import "testing"

func TestCountrune(t *testing.T) {
	for _, tc := range []struct {
		input  string
		input2 rune
		want   int
	}{
		{"", 76, 0},
		{"abc", 97, 1},
		{"aaa", 97, 3},
		{"242", 97, 0},
		//{"â", 0162, 1},
		//{"   ččabcčč   ", 0205, 4},
	} {
		t.Run(tc.input, func(t *testing.T) {
			if got := CountRune(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
