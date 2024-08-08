package solutions

import "testing"

func Test_countRoutes(t *testing.T) {
	type args struct {
		locations []int
		start     int
		finish    int
		fuel      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1DP",
			args: args{
				locations: []int{5, 2, 1},
				start:     0,
				finish:    2,
				fuel:      40,
			},
			want: 89844,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRoutes(tt.args.locations, tt.args.start, tt.args.finish, tt.args.fuel); got != tt.want {
				t.Errorf("countRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRoutesRecursive(t *testing.T) {
	type args struct {
		locations []int
		start     int
		finish    int
		fuel      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1Rec",
			args: args{
				locations: []int{5, 2, 1},
				start:     0,
				finish:    2,
				fuel:      40,
			},
			want: 89844,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRoutesRecursive(tt.args.locations, tt.args.start, tt.args.finish, tt.args.fuel); got != tt.want {
				t.Errorf("countRoutesRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
