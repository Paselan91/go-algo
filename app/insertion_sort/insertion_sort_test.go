package insertion_sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		arrays []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "正常 1~6までソートできること",
			args: args{
				arrays: []int{3, 4, 2, 5, 6, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.arrays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "正常 2変数をswapできること",
			args: args{
				a: 100,
				b: 200,
			},
			want:  200,
			want1: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swap(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
