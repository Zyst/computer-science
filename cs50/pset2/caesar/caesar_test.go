package main

import "testing"

func Test_caesarShift(t *testing.T) {
	type args struct {
		letter byte
		offset int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Y", args{'Y', 2}, 'A'},
		{"y", args{'y', 2}, 'a'},
		{"B", args{'B', 13}, 'O'},
		{"b", args{'b', 13}, 'o'},
		{"j", args{'w', 13}, 'j'},
		{"!", args{'!', 3}, '!'},
		{" ", args{' ', 8}, ' '},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesarShift(tt.args.letter, tt.args.offset); got != tt.want {
				t.Errorf("caesarShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_caesarCipher(t *testing.T) {
	type args struct {
		cipher string
		offset int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Cipher memes",
			args{"Be sure to drink your Ovaltine!", 13},
			"Or fher gb qevax lbhe Binygvar!",
		},
		{
			// I laughed IRL at this one
			"?",
			args{"uggc://jjj.lbhghor.pbz/jngpu?i=bUt5FWLEUN0", 13},
			"http://www.youtube.com/watch?v=oHg5SJYRHA0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesarCipher(tt.args.cipher, tt.args.offset); got != tt.want {
				t.Errorf("caesarCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
