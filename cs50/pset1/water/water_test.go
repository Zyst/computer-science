package main

import "testing"

func Test_minutesToWaterBottles(t *testing.T) {
	type args struct {
		min int
	}
	tests := []struct {
		name    string
		minutes int
		want    int
	}{
		{"120", 10, 120},
		{"1200", 100, 1200},
		{"180", 15, 180},
		{"3996", 333, 3996},
		{"12000", 1000, 12000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minutesToWaterBottles(tt.minutes); got != tt.want {
				t.Errorf("minutesToWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
