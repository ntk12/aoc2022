package main

import "testing"

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
