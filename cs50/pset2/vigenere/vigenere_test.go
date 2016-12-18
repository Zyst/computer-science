package main

import "testing"

func Test_vigenereShift(t *testing.T) {
	type args struct {
		letter byte
		offset byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
	  {"AA", args{'A', 'A'}, 'A'},
	  {"MN", args{'M', 'b'}, 'N'},
	  {"MN", args{'M', 'B'}, 'N'},
	  {"aa", args{'a', 'a'}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vigenereShift(tt.args.letter, tt.args.offset); got != tt.want {
				t.Errorf("vigenereShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vigenereCipher(t *testing.T) {
	type args struct {
		encrypt string
		key     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Meet me at the park at eleven am",
			args{"Meet me at the park at eleven am", "bacon"},
			"Negh zf av huf pcfx bt gzrwep oz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vigenereCipher(tt.args.encrypt, tt.args.key); got != tt.want {
				t.Errorf("vigenereCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
