package reload

func Hexer(s string) bool {
	return s == "(hex)"
}

func Biner(s string) bool {
	return s == "(bin)"
}

func Upper_d(s string) bool {
	return s == "(up,"
}

func Upper(s string) bool {
	return s == "(up)"
}

func Lower_d(s string) bool {
	return s == "(low,"
}

func Lower(s string) bool {
	return s == "(low)"
}

func Capper_d(s string) bool {
	return s == "(cap,"
}

func Capper(s string) bool {
	return s == "(cap)"
}

func Index(arr []int, n int) bool {
	for _, j := range arr {
		if j == n {
			return true
		}
	}
	// [1 2],[3 4]
	// for _, j := range arr[n] {
	// 	if j == m {
	// 		fmt.Println("Found")
	// 		return true
	// 	}
	// }
	return false
}
