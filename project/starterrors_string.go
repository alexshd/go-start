// Code generated by "stringer -type=StartErrors"; DO NOT EDIT.

package project

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NameError-11]
}

const _StartErrors_name = "NameError"

var _StartErrors_index = [...]uint8{0, 9}

func (i StartErrors) String() string {
	i -= 11
	if i < 0 || i >= StartErrors(len(_StartErrors_index)-1) {
		return "StartErrors(" + strconv.FormatInt(int64(i+11), 10) + ")"
	}
	return _StartErrors_name[_StartErrors_index[i]:_StartErrors_index[i+1]]
}