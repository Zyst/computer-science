package main

import "testing"

func Test_getInitials(t *testing.T) {
	type args struct {
		fullname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Erick", args{"Erick Romero"}, "ER"},
		{"Erick For Real", args{"Erick Eduardo Romero Vargas Rebeil"}, "EERVR"},
		{"AGS", args{"Angel Gomez Salazar"}, "AGS"},
		{"Zen Dev", args{"Zen Dev"}, "ZD"},
		{"Zamyla Chan", args{"Zamyla Chan"}, "ZC"},
		{"robert thomas bowden", args{"robert thomas bowden"}, "RTB"},
		{"Joseph Gordon-Levitt", args{"Joseph Gordon-Levitt"}, "JG"},
		{"Conan O’Brien", args{"Conan O’Brien"}, "CO"},
		{"David J. Malan", args{"David J. Malan"}, "DJM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInitials(tt.args.fullname); got != tt.want {
				t.Errorf("getInitials() = %v, want %v", got, tt.want)
			}
		})
	}
}
