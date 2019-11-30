package gomod_sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		numbers []interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Check sum",
			args: args{
				numbers: []interface{}{1, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
