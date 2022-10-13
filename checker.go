package reload

func vowels(a rune) bool {
	return a == 'i' || a == 'e' || a == 'a' || a == 'o' || a == 'u' || a == 'h'
}

func Punct(a rune) bool {
	return a == '.' || a == ',' || a == '!' || a == '?' || a == ':' || a == ';'
}

func Checker(str []string, bimer []bool) []string {
	t := make([]string, 0)
	for i, j := range str {
		if Hexer(j) || Biner(j) || Upper(j) || Lower(j) || Capper(j) || Index(Cap_ind, i) || Capper_d(j) || Index(Low_ind, i) || Lower_d(j) || Index(Up_ind, i) || Upper_d(j) || !bimer[i] {
			continue
		} else {
			t = append(t, j)
		}
	}
	return t
}

// func Checker(str [][]string, bimer [][]bool) [][]string {
// 	tt := make([][]string, 0)
// 	for l, k := range str {
// 		t := make([]string, 0)
// 		for i, j := range k {
// 			if Hexer(j) || Biner(j) || Upper(j) || Lower(j) || Capper(j) || Index(Cap_ind, l, i) || Capper_d(j) || Index(Low_ind, l, i) || Lower_d(j) || Index(Up_ind, l, i) || Upper_d(j) {
// 				continue
// 			} else if !bimer[l][i] {
// 				continue
// 			} else {
// 				t = append(t, j)
// 			}
// 		}
// 		tt = append(tt, t)
// 	}
// 	return tt
// }

func Vowel(str []string) {
	for i, j := range str {
		if j == "a" && i+1 != len(str) {
			runes := []rune(str[i+1])
			if vowels(runes[0]) {
				str[i] = "an"
			}
		}
	}
}

// func Vowel(str [][]string) {
// 	for _, k := range str {
// 		for i, j := range k {
// 			if j == "a" && i+1 != len(str) {
// 				runes := []rune(k[i+1])
// 				if vowels(runes[0]) {
// 					k[i] = "an"
// 				}
// 			}
// 		}
// 	}
// }

func Puncter(str []string) []bool {
	bimer := make([]bool, 0)
	for i, j := range str {
		flag := false
		for _, k := range j {
			if Punct(k) {
				flag = false
			} else {
				flag = true
				break
			}
		}
		bimer = append(bimer, flag)
		if !flag && i-1 >= 0 {
			str[i-1] += str[i]
		}
	}
	return bimer
}

// func Puncter(str [][]string) [][]bool {
// 	bimer := make([][]bool, 0)
// 	for _, l := range str {
// 		tt := make([]bool, 0)
// 		for i, j := range l {
// 			flag := false
// 			for _, k := range j {
// 				if Punct(k) {
// 					flag = false
// 					break
// 				} else {
// 					flag = true
// 				}
// 			}
// 			tt = append(tt, flag)
// 			if !flag && i-1 >= 0 {
// 				l[i-1] += l[i]
// 			}
// 		}
// 		bimer = append(bimer, tt)
// 	}
// 	return bimer
// }
