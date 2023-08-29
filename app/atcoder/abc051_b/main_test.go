package main

import (
	"testing"
)

func Test_countEqualS(t *testing.T) {
	type args struct {
		s int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常_(k, s)==(2, 2)",
			args: args{
				k: 2,
				s: 2,
			},
			want: 6,
		},
		{
			name: "正常_(k, s)==(5, 15)",
			args: args{
				k: 5,
				s: 15,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEqualS(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("countEqualS() = %v, want %v", got, tt.want)
			}
		})
	}
}
