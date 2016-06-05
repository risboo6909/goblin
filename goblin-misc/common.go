package misc

import (
	"math"
	"reflect"
)

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func cmpSlices(a, b interface{}) bool {
	xs := interfaceSlice(a)
	ys := interfaceSlice(b)

	if len(xs) != len(ys) {
		return false
	}

	for i, v := range xs {
		if v != ys[i] {
			return false
		}
	}

	return true
}

func minIntPair(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxIntPair(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
