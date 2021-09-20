package sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first", want: 2, args: args{x: 1, y: 1}},
		{name: "second", want: 0, args: args{x: -1, y: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

