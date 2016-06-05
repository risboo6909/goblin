package misc


func cmpSlices(a, b []Cell) bool {
	if len(a) != len(b) { return false }
	for i, v := range a { if v != b[i] { return false } }
	return true
}