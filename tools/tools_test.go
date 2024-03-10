// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/19 13:26:12
// * Proj: work
// * Pack: tools
// * File: tool_test.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package tools

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	type args[T any] struct {
		slice []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceHasElem(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		elem  T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5},
				elem:  3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceHasElem(tt.args.slice, tt.args.elem); got != tt.want {
				t.Errorf("SliceHasElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceCntElem(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		elem  T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5},
				elem:  3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceCntElem(tt.args.slice, tt.args.elem); got != tt.want {
				t.Errorf("SliceCntElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceDelElemAll(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		elem  T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "",
			args: args[int]{
				slice: []int{1, 2, 3, 3, 5},
				elem:  3,
			},
			want: []int{1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceDelElemAll(tt.args.slice, tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceDelElemAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlicesDifferSet(t *testing.T) {
	type args[T comparable] struct {
		slice1 []T
		slice2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "",
			args: args[int]{
				slice1: []int{1, 2, 3, 4, 5},
				slice2: []int{4, 5},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SlicesDifferSet(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlicesDifferSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlicesInterSet(t *testing.T) {
	type args[T comparable] struct {
		slice1 []T
		slice2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "",
			args: args[int]{
				slice1: []int{1, 2, 3, 4, 5},
				slice2: []int{4, 5},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SlicesInterSet(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlicesInterSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceStrsToInts(t *testing.T) {
	type args struct {
		array []string
	}
	type testCase[T constraints.Integer] struct {
		name    string
		args    args
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args{
				array: []string{"1", "2", "3", "4", "5"},
			},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceStrsToInts[int](tt.args.array)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceStrsToInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceStrsToInts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIntsToStrs(t *testing.T) {
	type args[T constraints.Integer] struct {
		array []T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want []string
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				array: []int{1, 2, 3, 4, 5},
			},
			want: []string{"1", "2", "3", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIntsToStrs(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceIntsToStrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixIntsToStrs(t *testing.T) {
	type args[T constraints.Integer] struct {
		arrays [][]T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want [][]string
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				arrays: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			},
			want: [][]string{{"1", "2", "3", "4", "5"}, {"6", "7", "8", "9", "10"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixIntsToStrs(tt.args.arrays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixIntsToStrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixStrsToInts(t *testing.T) {
	type args struct {
		arrays [][]string
	}
	type testCase[T constraints.Integer] struct {
		name    string
		args    args
		want    [][]T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args{
				arrays: [][]string{{"1", "2", "3", "4", "5"}, {"6", "7", "8", "9", "10"}},
			},
			want:    [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatrixStrsToInts[int](tt.args.arrays)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatrixStrsToInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixStrsToInts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "hello",
			},
			want: []byte("hello"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				b: []byte("hello"),
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomToken(t *testing.T) {
	token := GetRandomToken()
	t.Log(token)
}

func TestGetRandomString(t *testing.T) {
	type args struct {
		l    int
		with WithBase
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				l:    6,
				with: WithNumnber,
			},
			want: "0123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandomString(tt.args.l, tt.args.with)
			t.Log(got)

		})
	}
}

func TestSliceStrsToInts1(t *testing.T) {
	strs := []string{"1", "2", "3", "4", "5", "6"}
	var it8 = make([]int8, len(strs))
	var it16 = make([]int16, len(strs))
	var it32 = make([]int32, len(strs))
	var it64 = make([]int64, len(strs))
	var uit8 []uint8
	var uit16 []uint16
	var uit32 []uint32
	var uit64 []uint64

	SliceStrsToInts1(strs, it8)
	SliceStrsToInts1(strs, it16)
	SliceStrsToInts1(strs, it32)
	SliceStrsToInts1(strs, it64)
	//SliceStrsToInts1(strs, uit8)
	//SliceStrsToInts1(strs, uit16)
	//SliceStrsToInts1(strs, uit32)
	//SliceStrsToInts1(strs, uit64)

	t.Log(it8)
	t.Log(it16)
	t.Log(it32)
	t.Log(it64)
	t.Log(uit8)
	t.Log(uit16)
	t.Log(uit32)
	t.Log(uit64)
}
