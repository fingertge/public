// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/19 13:19:08
// * Proj: work
// * Pack: tools
// * File: tools.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package tools

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/exp/constraints"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

func ReverseSlice[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func SliceHasElem[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func SliceCntElem[T comparable](slice []T, elem T) int {
	cnt := 0
	for _, v := range slice {
		if v == elem {
			cnt++
		}
	}
	return cnt
}

func SliceDelElemAll[T comparable](slice []T, elem T) []T {
	var tmp []T
	for _, v := range slice {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func SlicesDifferSet[T comparable](slice1, slice2 []T) []T {
	mp := make(map[T]bool)
	for _, v := range slice1 {
		mp[v] = true
	}
	var tp []T
	for _, v := range slice2 {
		if ok := mp[v]; ok {
			delete(mp, v)
		}
	}
	for k := range mp {
		tp = append(tp, k)
	}
	return tp
}

func SlicesInterSet[T comparable](slice1, slice2 []T) []T {
	mp := make(map[T]bool)
	for _, v := range slice1 {
		mp[v] = true
	}
	var tp []T
	for _, v := range slice2 {
		if _, ok := mp[v]; ok {
			tp = append(tp, v)
		}
	}
	return tp
}

func SliceStrsToInts[T constraints.Integer](array []string) ([]T, error) {
	var ret = make([]T, len(array))
	for i, v := range array {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ret[i] = T(val)
	}
	return ret, nil
}

func SliceStrsToInts1[T constraints.Integer](array []string, tp []T) error {
	for i, v := range array {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		tp[i] = T(val)
	}
	return nil
}

func SliceIntsToStrs[T constraints.Integer](array []T) []string {
	var ret = make([]string, len(array))
	for i, v := range array {
		ret[i] = strconv.FormatInt(int64(v), 10)
	}
	return ret
}

func MatrixStrsToInts[T constraints.Integer](arrays [][]string) ([][]T, error) {
	var ret = make([][]T, len(arrays))
	for i, array := range arrays {
		var err error
		ret[i], err = SliceStrsToInts[T](array)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func MatrixIntsToStrs[T constraints.Integer](arrays [][]T) [][]string {
	var ret = make([][]string, len(arrays))
	for i, array := range arrays {
		ret[i] = SliceIntsToStrs[T](array)
	}
	return ret
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

const (
	Number         string = "0123456789"
	Lower          string = "abcdefghijklmnopqrstuvwxyz"
	Upper          string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Alpha          string = Lower + Upper
	NumberAndLower string = Number + Lower
	NumberAndUpper string = Number + Upper
	NumberAndAlpha string = Number + Alpha
)

type WithBase func() string

func WithNumnber() string {
	return Number
}

func WithLower() string {
	return Lower
}

func WithUpper() string {
	return Upper
}

func WithAlpha() string {
	return Alpha
}

func WithNumberAndLower() string {
	return NumberAndLower
}

func WithNumberAndUpper() string {
	return NumberAndUpper
}

func WithNumberAndAlpha() string {
	return NumberAndAlpha
}

func GetRandomString(l int, with WithBase) string {
	str := with()
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomToken() string {
	nowTime := time.Now().UnixNano()
	mdobj := md5.New()
	mdobj.Write(StringToBytes(GetRandomString(11, WithNumberAndAlpha) + strconv.FormatInt(nowTime, 10)))
	sum := mdobj.Sum(nil)
	return hex.EncodeToString(sum)
}
