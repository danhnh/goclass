package ex3

import (
	"reflect"
	"testing"
	"unsafe"
)

type testcase struct {
	want  interface{}
	cases []interface{}
}

func TestIsEmpty(t *testing.T) {
	var tempString = "temp"
	var tempNil *struct{}
	var temp []int
	testCases := []testcase{
		{
			want: true, cases: []interface{}{
				false, "", 0, uint(0), float64(0), []int{}, []uint{}, []uint8{}, []float64{}, nil, tempNil, chan string(nil), map[string]string(nil), temp,
			},
		},
		{
			want: false, cases: []interface{}{
				true, "a", 1, uint(10), float64(2), -1, []int{1, 2, 3}, [3]int{1, 2, 3}, []uint{1, 2, 3}, []uint8{1, 2, 3}, []float64{1}, &tempString, unsafe.Pointer(&tempString), map[string]string{"a": "a"}, func() {},
			},
		},
	}
	for c := range testCases {
		for v := range testCases[c].cases {
			if r := IsEmpty(testCases[c].cases[v]); r != testCases[c].want {
				t.Errorf("\tIsEmpty(%T) has error or incorrect value %t when %s is expected \n", testCases[c].cases[v], r, testCases[c].want)
			}
		}
	}
}

func TestLast(t *testing.T) {
	testErrorCase := []interface{}{
		nil, [0]int{}, []int{}, map[string]string{}, map[string]string{"z": "3", "a": "1", "b": "2"},
	}
	testCases := []testcase{
		{
			want: 1, cases: []interface{}{
				[1]int{1}, [2]int{2, 1},
			},
		},
		{
			want: 2, cases: []interface{}{
				[1]int{2}, []int{1, 2, 3, 2},
			},
		},
		{
			want: "3", cases: []interface{}{
				[3]string{"1", "2", "3"}, []string{"3"},
			},
		},
	}
	for p := range testErrorCase {
		t.Run("test error cases", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("\tLast(%T) not raise error\n", testErrorCase[p])
				}
			}()
			Last(testErrorCase[p])
		})
	}
	for c := range testCases {
		for v := range testCases[c].cases {
			if r := Last(testCases[c].cases[v]); r != testCases[c].want {
				t.Errorf("\tLast(%T) has error or incorrect value %s when %s is expected \n", testCases[c].cases[v], r, testCases[c].want)
			}
		}
	}
}

func TestMap(t *testing.T) {
	type parameter struct {
		p1, p2 interface{}
	}
	testErrorCase := []parameter{
		{nil, nil},
		{1, nil},
		{[0]int{}, nil},
		{[]int{}, nil},
		{map[string]string{}, nil},
		{nil, func() {}},
	}
	testCases := []testcase{
		{
			want: []int{}, cases: []interface{}{
				parameter{[0]int{}, func(v int) int {
					return v * 2
				}},
			},
		},
		{
			want: []int{2}, cases: []interface{}{
				parameter{[1]int{1}, func() int {
					return 2
				}},
			},
		},
		{
			want: []int{2,4,6}, cases: []interface{}{
				parameter{[3]int{1, 2, 3}, func(v int) int {
					return v*2
				}},
			},
		},
		{
			want: []float32{2,4,6}, cases: []interface{}{
				parameter{[3]int{1, 2, 3}, func(v int) float32 {
					return float32(v)*2
				}},
			},
		},
	}
	for p := range testErrorCase {
		t.Run("test error cases", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("\tMap(%T,%T) not raise error\n", testErrorCase[p].p1, testErrorCase[p].p2)
				}
			}()
			Map(testErrorCase[p].p1, testErrorCase[p].p2)
		})
	}

	for c := range testCases {
		for v := range testCases[c].cases {
			if r := Map(testCases[c].cases[v].(parameter).p1, testCases[c].cases[v].(parameter).p2); !reflect.DeepEqual(r, testCases[c].want) {
				t.Errorf("\tMap(%T,%T) has error or incorrect value %T when %T is expected \n", testCases[c].cases[v].(parameter).p1, testCases[c].cases[v].(parameter).p2, r, testCases[c].want)
			}
		}
	}
}

func TestMax(t *testing.T) {
	testErrorCase := []interface{}{
		nil, [0]int{}, []int{}, map[string]string{}, [4]string{"5", "2", "1", "3"},
	}
	testCases := []testcase{
		{
			want: 1, cases: []interface{}{
				[1]int{1}, []int{1,0},
			},
		},
		{
			want: 2, cases: []interface{}{
				[1]int{2}, []int{1,0,2,0,1},
			},
		},
		{
			want: uint(5), cases: []interface{}{
			[1]uint{5}, []uint{1,0,2,5,1},
		},
		},
		{
			want: float32(5), cases: []interface{}{
			[1]float32{5}, []float32{1,0,2,5,1},
		},
		},
	}
	for p := range testErrorCase {
		t.Run("test error cases", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("\tMax(%T) not raise error\n", testErrorCase[p])
				}
			}()
			Max(testErrorCase[p])
		})
	}

	for c := range testCases {
		for v := range testCases[c].cases {
			if r := Max(testCases[c].cases[v]); r != testCases[c].want {
				t.Errorf("\tMax(%T) has error or incorrect value %t when %s is expected \n", testCases[c].cases[v], r, testCases[c].want)
			}
		}
	}
}

func TestIndexOf(t *testing.T) {
	r := IndexOf(nil, 0, 0)
	if r != -1 {
		t.Errorf("\tIndexOf(nil) has error or incorrect value %d\n", r)
	}
	r = IndexOf([0]int{}, 0, 0)
	if r != -1 {
		t.Errorf("\tIndexOf(nil) has error or incorrect value %d\n", r)
	}
	r = IndexOf(map[string]string{}, "a", 0)
	if r != -1 {
		t.Errorf("\tIndexOf(map[string]string{}) has error or incorrect value %d\n", r)
	}
	r = IndexOf([3]int{3, 1, 2}, 0, 0)
	if r != -1 {
		t.Errorf("\tIndexOf([3]int{3, 1, 2}, 0, 0) has error or incorrect value %d\n", r)
	}
	r = IndexOf([3]int{3, 1, 2}, 1, 0)
	if r != 1 {
		t.Errorf("\tIndexOf([3]int{3, 1, 2}, 0, 0) has error or incorrect value %d\n", r)
	}
	r = IndexOf([6]int{3, 1, 2, 1, 2, 3}, 1, 0)
	if r != 1 {
		t.Errorf("\tIndexOf([3]int{3, 1, 2}, 0, 0) has error or incorrect value %d\n", r)
	}
	r = IndexOf([6]int{3, 1, 2, 1, 2, 3}, 1, 2)
	if r != 3 {
		t.Errorf("\tIndexOf([3]int{3, 1, 2}, 0, 0) has error or incorrect value %d\n", r)
	}
	r = IndexOf([6]string{"3", "1", "2", "1", "2", "3"}, "3", 2)
	if r != 5 {
		t.Errorf("\tIndexOf([3]int{3, 1, 2}, 0, 0) has error or incorrect value %d\n", r)
	}
	r = IndexOf([6]map[string]string{{"3": "3"}, {"1": "3"}, {"2": "2"}, {"1": "1"}, {"2": "2"}, {"3": "3"}}, map[string]string{"3": "3"}, 0)
	if r != -1 {
		t.Errorf("\tIndexOf(Array of map) has error or incorrect value %d\n", r)
	}
}

// func TestHead(t *testing.T) {
// 	type args struct {
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Head(tt.args.v); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Head() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
