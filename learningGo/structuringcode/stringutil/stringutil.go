package stringutil

func FullName(first, last string) (string, int) {
	full := first + " " + last
	length := len(full)
	return full, length
}

func FullNameNakedReturn(first, last string) (full string, length int) {
	full = first + " " + last
	length = len(full)
	return
}