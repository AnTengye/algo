package dp

import "testing"

func Test_bl(t *testing.T) {
	type args struct {
		wgt []int
		val []int
		i   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "70",
			args: args{
				wgt: []int{10, 20, 30, 40, 50},
				val: []int{50, 120, 150, 210, 240},
				i:   5,
				c:   70,
			},
			want: 380,
		},
		{
			name: "90",
			args: args{
				wgt: []int{10, 20, 30, 40, 50},
				val: []int{50, 120, 150, 210, 240},
				i:   5,
				c:   90,
			},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bl(tt.args.wgt, tt.args.val, tt.args.i, tt.args.c); got != tt.want {
				t.Errorf("bl() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bl1(tt.args.wgt, tt.args.val, tt.args.i, tt.args.c); got != tt.want {
				t.Errorf("bl() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.wgt, tt.args.val, tt.args.i, tt.args.c); got != tt.want {
				t.Errorf("bl() = %v, want %v", got, tt.want)
			}
		})
	}
}
