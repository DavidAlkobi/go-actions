// BEGIN: 8f7d6e5b5c4a
package main

import "testing"

func TestSumInts(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "sum of positive integers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "sum of negative integers",
			nums: []int{-1, -2, -3, -4, -5},
			want: -15,
		},
		{
			name: "sum of mixed integers",
			nums: []int{-1, 2, -3, 4, -5},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumInts(tt.nums...); got != tt.want {
				t.Errorf("sumInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END: 8f7d6e5b5c4a
