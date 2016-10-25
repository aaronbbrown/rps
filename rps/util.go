package rps

// a sane modulo operator that works like every other damn language
// see https://github.com/golang/go/issues/448
// https://groups.google.com/forum/#!topic/golang-nuts/xj7CV857vAg
func saneModInt(x, y int) int {
	result := x % y
	if result < 0 {
		result += y
	}
	return result
}
