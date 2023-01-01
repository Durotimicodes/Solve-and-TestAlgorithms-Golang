package testalgorithms

import (
	algorithms "algorithms/Algo"
	"reflect"
	"testing"
)

func TestEqualizeArray(t *testing.T) {

	tests := []struct {
		name  string
		input []int32
		want  int32
	}{
		{name: "Case-One", input: []int32{1, 2, 2, 3}, want: int32(2)},
		{name: "Case-Two", input: []int32{3, 3, 2, 1, 3}, want: int32(2)},
		{name: "Case-Three", input: []int32{3, 1, 1, 2}, want: int32(2)},
		{name: "Case-Four", input: []int32{69, 86, 100, 69, 55, 83, 15, 69, 86, 69, 79, 16, 86, 24, 24, 55, 16, 69, 100, 79, 16, 83, 83, 79, 15, 15, 86, 16, 55, 18, 100, 100, 86, 16, 83, 79, 86, 83, 100, 83, 55, 15, 86, 86, 55, 100, 55, 18, 55, 100, 86, 69, 83, 24, 16, 55, 100, 16, 100, 24, 16, 55, 15, 79, 16, 18, 16, 16, 83, 83, 69, 18, 100, 79, 16, 24, 79, 16, 69, 86, 83, 79, 83, 18, 15, 100, 24, 83}, want: int32(75)},
		{name: "Case-Five", input: []int32{35, 65, 69, 28, 12, 69, 37, 80, 80, 50, 80, 50, 15, 57, 79, 69, 57, 65, 15, 69, 57, 50, 65, 2, 14, 64, 15, 65, 65, 5, 15, 64, 57, 37, 50, 12, 64, 37, 28, 20, 80, 80, 50}, want: 38},
	}

	for _, tc := range tests {
		got := algorithms.EqualizeArray(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.name, tc.want, got)
		}
	}
}

func TestCheckGemStone(t *testing.T) {

}
