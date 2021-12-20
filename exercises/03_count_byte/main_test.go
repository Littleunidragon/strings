package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		input  string
		input2 byte
		want   int
	}{
		{"", 102, 0},
		{"abc", 97, 1},
		{"aaa", 97, 3},
		{"242", 97, 0},
		{"a", 97, 1},
		{"   abc   ", 97, 1},
	} {
		t.Run(tc.input, func(t *testing.T) {
			if got := CountByte(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
