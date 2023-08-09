package main

import (
	"testing"
)

func TestPalindromeTwo(t *testing.T) {
	var tests = []struct {
		Sample string
		Output string
	}{
		{"Noel - sees Leon", "true"},
		{"A war at Tarawa!", "true"},
		{"Mr. Owl ate my metal worm", "true"},
		{"Was it a car or a cat I saw?", "true"},
		{"Rats live on no evil star!", "true"},
		{"Live on time, emit no evil", "true"},
		{"not true!", "false"},
	}

	for _, tt := range tests {
		t.Run(tt.Sample, func(t *testing.T) {
			want := tt.Output
			got := PalindromeTwo(tt.Sample)
			if want != got {
				t.Errorf("want %v got %v", want, got)
			}
		})
	}
}
