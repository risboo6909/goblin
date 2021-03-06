package misc

import (
	"math"
	"reflect"
	"testing"
	"runtime"
)

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	}
	return false
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundFloat(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return round(f * shift) / shift;
}

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

func assertEqual(t *testing.T, x, y interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	if reflect.TypeOf(x).Kind() == reflect.Slice && reflect.TypeOf(x).Kind() == reflect.Slice {
		if !cmpSlices(x, y) {
			t.Fatalf("Asserion failed in %v, %v != %v",
				runtime.FuncForPC(pc).Name(), x, y)
		}
	} else {
		if x != y {
			t.Fatalf("Asserion failed in %v, %v != %v",
				runtime.FuncForPC(pc).Name(), x, y)
		}
	}
}

func minIntPair(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxIntPair(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// the idea has been taken from
// https://github.com/bradfitz/iter/
func intRange(n int) []int {
	return make([]int, n)
}
