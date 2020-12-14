package top

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "4",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "5",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "6",
			args: args{
				s: "jfsafdndfafd",
			},
			want: "afdndfa",
		},
		{
			name: "6",
			args: args{
				s: "ababababa",
			},
			want: "ababababa",
		},
		{
			name: "7",
			args: args{
				s: "tattarrattat",
			},
			want: "tattarrattat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}