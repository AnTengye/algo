package dp

import "testing"

func TestSolution1(t *testing.T) {
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
				wgt: []int{1, 2, 5},
				i:   3,
				c:   4,
			},
			want: 380,
		},
		{
			name: "90",
			args: args{
				wgt: []int{1, 2, 5},
				i:   3,
				c:   10,
			},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution1(tt.args.wgt, tt.args.val, tt.args.i, tt.args.c); got != tt.want {
				t.Errorf("Solution1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	type args struct {
		coins []int
		i     int
		a     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "70",
			args: args{
				coins: []int{1, 2, 5},
				i:     3,
				a:     4,
			},
			want: 2,
		},
		{
			name: "90",
			args: args{
				coins: []int{1, 2, 5},
				i:     3,
				a:     12,
			},
			want: 3,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := Solution2(tt.args.coins, tt.args.i, tt.args.a); got != tt.want {
	//			t.Errorf("Solution2() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution2(tt.args.coins, tt.args.i, tt.args.a); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	type args struct {
		coins []int
		i     int
		a     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "70",
			args: args{
				coins: []int{1, 2, 5},
				i:     3,
				a:     5,
			},
			want: 4,
		},
		{
			name: "90",
			args: args{
				coins: []int{1, 2},
				i:     2,
				a:     5,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution3(tt.args.coins, tt.args.i, tt.args.a); got != tt.want {
				t.Errorf("Solution3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution4(t *testing.T) {
	type args struct {
		s []string
		t []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "today to apple",
			args: args{
				//today
				s: []string{"t", "o", "d", "a", "y"},
				//apple
				t: []string{"a", "p", "p", "l", "e"},
			},
			want: 5,
		},
		{
			name: "summery to stay",
			args: args{
				//summery
				s: []string{"s", "u", "m", "m", "e", "r", "y"},
				//stay
				t: []string{"s", "t", "a", "y"},
			},
			want: 5,
		},
		{
			name: "bag to pack",
			args: args{
				//bag
				s: []string{"b", "a", "g"},
				//pack
				t: []string{"p", "a", "c", "k"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution4(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Solution4() = %v, want %v", got, tt.want)
			}
		})
	}
}
