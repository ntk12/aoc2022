package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	cases := []struct {
		op   string
		you  string
		want int
	}{
		{
			op:   "A",
			you:  "Y",
			want: 8,
		},
		{
			op:   "B",
			you:  "X",
			want: 1,
		},
		{
			op:   "C",
			you:  "Z",
			want: 6,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if got := Game(tt.op, tt.you); got != tt.want {
				t.Errorf("Game(%s, %s) = %v, want %v", tt.op, tt.you, got, tt.want)
			}
		})
	}
}

func TestGame2(t *testing.T) {
	tests := []struct {
		op   string
		end  string
		want int
	}{
		{
			op:   "A",
			end:  "Y",
			want: 4,
		},
		{
			op:   "B",
			end:  "X",
			want: 1,
		},
		{
			op:   "C",
			end:  "Z",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Game2(tt.op, tt.end); got != tt.want {
				t.Errorf("Game2(%s, %s) = %v, want %v", tt.op, tt.end, got, tt.want)
			}
		})
	}
}
