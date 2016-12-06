package main

import "testing"

func Test_giveChange(t *testing.T) {
	type args struct {
		money float64
		coins int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "0.41", args{ 0.41, 0 }, 4 },
		{ "0.04", args{ 0.04, 0 }, 4 },
		{ "0.01", args{ 0.01, 0 }, 1 },
		{ "0.05", args{ 0.05, 0 }, 1 },
		{ "0.10", args{ 0.10, 0 }, 1 },
		{ "0.25", args{ 0.25, 0 }, 1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := giveChange(tt.args.money, tt.args.coins); got != tt.want {
				t.Errorf("giveChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
