package reload

import "strings"

func Single(str []string) {
	for i, j := range str {
		if Capper(j) && i-1 >= 0 {
			runes := []rune(str[i-1])
			if runes[0] >= 'a' && runes[0] <= 'z' {
				runes[0] -= 32
			}
			str[i-1] = string(runes)
		} else if Upper(j) && i-1 >= 0 {
			runes := []rune(str[i-1])
			for k := 0; k < len(runes); k++ {
				if runes[k] >= 'a' && runes[k] <= 'z' {
					runes[k] -= 32
				}
			}
			str[i-1] = string(runes)
		} else if Lower(j) && i-1 >= 0 {
			runes := []rune(str[i-1])
			for k := 0; k < len(runes); k++ {
				if runes[k] >= 'A' && runes[k] <= 'Z' {
					runes[k] += 32
				}
			}
			str[i-1] = string(runes)
		}
	}
}

// func Single(str [][]string) {
// 	for _, l := range str {
// 		for i, j := range l {
// 			if Capper(j) && i-1 >= 0 {
// 				runes := []rune(l[i-1])
// 				if runes[0] >= 'a' && runes[0] <= 'z' {
// 					runes[0] -= 32
// 				}
// 				l[i-1] = string(runes)
// 			} else if Upper(j) && i-1 >= 0 {
// 				runes := []rune(l[i-1])
// 				for k := 0; k < len(runes); k++ {
// 					if runes[k] >= 'a' && runes[k] <= 'z' {
// 						runes[k] -= 32
// 					}
// 				}
// 				l[i-1] = string(runes)
// 			} else if Lower(j) && i-1 >= 0 {
// 				runes := []rune(l[i-1])
// 				for k := 0; k < len(runes); k++ {
// 					if runes[k] >= 'A' && runes[k] <= 'Z' {
// 						runes[k] += 32
// 					}
// 				}
// 				l[i-1] = string(runes)
// 			}
// 		}
// 	}
// }
// func NonSingle(str [][]string) {
// 	Keys["(cap,"] = Caps
// 	Keys["(low,"] = Lows
// 	Keys["(up,"] = Ups
// 	capper := 0
// 	for _, f := range str {
// 		for i := 0; i < len(Keys["(cap,"]); i++ {
// 			for j, k := range f {
// 				if Capper_d(k) && j-Keys["(cap,"][capper] >= 0 {
// 					for l := j - 1; l >= j-Keys["(cap,"][capper]; l-- {
// 						runes := []rune(f[l])
// 						if runes[0] >= 'a' && runes[0] <= 'z' {
// 							runes[0] -= 32
// 						}
// 						f[l] = string(runes)
// 					}
// 				}
// 			}
// 			capper++
// 		}
// 	}
// 	lower := 0
// 	for _, f := range str {
// 		for i := 0; i < len(Keys["(low,"]); i++ {
// 			for j, k := range f {
// 				if Lower_d(k) && j-Keys["(low,"][lower] >= 0 {
// 					for l := j - 1; l >= j-Keys["(low,"][lower]; l-- {
// 						f[l] = strings.ToLower(f[l])
// 					}
// 				}
// 			}
// 			lower++
// 		}
// 	}
// 	upper := 0
// 	for _, f := range str {
// 		for i := 0; i < len(Keys["(up,"]); i++ {
// 			for j, k := range f {
// 				if Upper_d(k) && j-Keys["(up,"][upper] >= 0 {
// 					for l := j - 1; l >= j-Keys["(up,"][upper]; l-- {
// 						f[l] = strings.ToUpper(f[l])
// 					}
// 				}
// 			}
// 			upper++
// 		}
// 	}
// }

func NonSingle(str []string) {
	Keys["(cap,"] = Caps
	Keys["(low,"] = Lows
	Keys["(up,"] = Ups
	capper := 0
	for i := 0; i < len(Keys["(cap,"]); i++ {
		for j, k := range str {
			if Capper_d(k) && j-Keys["(cap,"][capper] >= 0 {
				for l := j - 1; l >= j-Keys["(cap,"][capper]; l-- {
					runes := []rune(str[l])
					if runes[0] >= 'a' && runes[0] <= 'z' {
						runes[0] -= 32
					}
					str[l] = string(runes)
				}
			}
		}
		capper++
	}
	lower := 0
	for i := 0; i < len(Keys["(low,"]); i++ {
		for j, k := range str {
			if Lower_d(k) && j-Keys["(low,"][lower] >= 0 {
				for l := j - 1; l >= j-Keys["(low,"][lower]; l-- {
					str[l] = strings.ToLower(str[l])
				}
			}
		}
		lower++
	}
	upper := 0
	for i := 0; i < len(Keys["(up,"]); i++ {
		for j, k := range str {
			if Upper_d(k) && j-Keys["(up,"][upper] >= 0 {
				for l := j - 1; l >= j-Keys["(up,"][upper]; l-- {
					str[l] = strings.ToUpper(str[l])
				}
			}
		}
		upper++
	}
}
