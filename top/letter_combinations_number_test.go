package top

import (
	"reflect"
	"sort"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				digits: "23",
			},
			want: []string{"ad","ae","af","bd","be","bf","cd","ce","cf"},
		},{
			name: "2",
			args: args{
				digits: "",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits);
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}