package solutions

import "testing"

func Test_numberOfCombinations(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				"327",
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				"094",
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				"123213124231412",
			},
			want: 67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfCombinations(tt.args.nums); got != tt.want {
				t.Errorf("numberOfCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
