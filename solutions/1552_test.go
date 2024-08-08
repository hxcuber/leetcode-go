package solutions

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		position []int
		m        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				position: []int{79, 74, 57, 22},
				m:        4,
			},
			want: 5,
		},
		{
			name: "t2",
			args: args{
				position: []int{269826447, 974181916, 225871443, 189215924, 664652743, 592895362, 754562271, 335067223, 996014894, 510353008, 48640772, 228945137},
				m:        3,
			},
			want: 461712236,
		},
		{
			name: "t3",
			args: args{
				position: []int{1, 2, 3, 4, 7},
				m:        3,
			},
			want: 3,
		},
		{
			name: "t4",
			args: args{
				position: []int{4784, 9049, 3151, 7537, 2734, 1287, 2875, 6770, 9565, 6254, 6898, 2509, 6583},
				m:        13,
			},
			want: 128,
		},
		{
			name: "t5",
			args: args{
				position: []int{315472927, 838019418, 192618105, 12737314, 91324553, 870388082, 876868586, 421284402, 89843519, 758122835, 537247841, 25028360, 679362833, 964682238, 575887767, 352564662, 252976366, 757254613, 917719676, 363623981, 573425483, 663573618, 254765807, 523341581, 457593095},
				m:        22,
			},
			want: 2462284,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.position, tt.args.m); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
