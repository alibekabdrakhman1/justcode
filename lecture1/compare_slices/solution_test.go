package compare_slices

import "testing"

func TestCompare(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{a: []int{1, 2, 3, 4, 6}, b: []int{4, 3, 2, 6, 1}},
			want: true,
		},
		{
			name: "false",
			args: args{
				a: []int{1, 2, 2, 3, 4, 5},
				b: []int{1, 2, 3, 4, 5, 6},
			},
			want: false,
		},
		{
			name: "lengths not equal",
			args: args{
				a: []int{1, 2, 3, 4, 6},
				b: []int{4, 3, 2, 6, 1, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
