package services

func Reorder(a int, b int) (int, int) {
	if a > b {return b, a}; return a, b
}

