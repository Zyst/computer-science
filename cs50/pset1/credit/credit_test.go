package main

import "testing"

func Test_checkCardLengthValidity(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"378282246310005"}, true},
		{"true", args{"371449635398431"}, true},
		{"true", args{"5555555555554444"}, true},
		{"true", args{"5105105105105100"}, true},
		{"true", args{"4111111111111111"}, true},
		{"true", args{"4012888888881881"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCardLengthValidity(tt.args.card); got != tt.want {
				t.Errorf("checkCardLengthValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInputIsNumbers(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"378282246310005"}, true},
		{"true", args{"371449635398431"}, true},
		{"true", args{"5555555555554444"}, true},
		{"true", args{"5105105105105100"}, true},
		{"true", args{"4111111111111111"}, true},
		{"true", args{"4012888888881881"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInputIsNumbers(tt.args.card); got != tt.want {
				t.Errorf("checkInputIsNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyCardWithLuhnAlgorithm(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"378282246310005"}, true},
		{"true", args{"371449635398431"}, true},
		{"true", args{"5555555555554444"}, true},
		{"true", args{"5105105105105100"}, true},
		{"true", args{"4111111111111111"}, true},
		{"true", args{"4012888888881881"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyCardWithLuhnAlgorithm(tt.args.card); got != tt.want {
				t.Errorf("verifyCardWithLuhnAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returnCardType(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"American Express", args{"378282246310005"}, "American Express"},
		{"American Express", args{"371449635398431"}, "American Express"},
		{"Mastercard", args{"5555555555554444"}, "Mastercard"},
		{"Mastercard", args{"5105105105105100"}, "Mastercard"},
		{"Visa", args{"4111111111111111"}, "Visa"},
		{"Visa", args{"4012888888881881"}, "Visa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnCardType(tt.args.card); got != tt.want {
				t.Errorf("returnCardType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCardAndReturnType(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"American Express", args{"378282246310005"}, "American Express"},
		{"American Express", args{"371449635398431"}, "American Express"},
		{"Mastercard", args{"5555555555554444"}, "Mastercard"},
		{"Mastercard", args{"5105105105105100"}, "Mastercard"},
		{"Visa", args{"4111111111111111"}, "Visa"},
		{"Visa", args{"4012888888881881"}, "Visa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCardAndReturnType(tt.args.card); got != tt.want {
				t.Errorf("validateCardAndReturnType() = %v, want %v", got, tt.want)
			}
		})
	}
}
