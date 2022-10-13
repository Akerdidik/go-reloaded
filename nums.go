package reload

import (
	"fmt"
	"strconv"
)

func Numer(str []string) {
	for i, j := range str {
		if Hexer(j) && i-1 >= 0 {
			tt, err := strconv.ParseInt(str[i-1], 16, 64)
			if err != nil {
				fmt.Println(err.Error())
			}
			t := strconv.Itoa(int(tt))
			Indexes = append(Indexes, i)
			str[i-1] = t
		} else if Biner(j) && i-1 >= 0 {
			tt, err := strconv.ParseInt(str[i-1], 2, 64)
			if err != nil {
				fmt.Println(err.Error())
			}
			t := strconv.Itoa(int(tt))
			Indexes = append(Indexes, i)
			str[i-1] = t
		}
	}
}

// func Numer(str [][]string) {
// 	for _, j := range str {
// 		for k, l := range j {
// 			if Hexer(l) && k-1 >= 0 {
// 				tt, err := strconv.ParseInt(j[k-1], 16, 64)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					os.Exit(1)
// 				}
// 				t := strconv.Itoa(int(tt))
// 				Indexes = append(Indexes, k)
// 				j[k-1] = t
// 			} else if Biner(l) && k-1 >= 0 {
// 				tt, err := strconv.ParseInt(j[k-1], 2, 64)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					os.Exit(1)
// 				}
// 				t := strconv.Itoa(int(tt))
// 				Indexes = append(Indexes, k)
// 				j[k-1] = t
// 			}
// 		}
// 	}
// }

func Indexer(str []string) {
	for i, j := range str {
		if Capper_d(j) && i+1 != len(str) {
			vre, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
			if err != nil {
				fmt.Println(err.Error())
			}
			Caps = append(Caps, vre)
			Cap_ind = append(Cap_ind, i+1)
		} else if Lower_d(j) && i+1 != len(str) {
			vre, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
			if err != nil {
				fmt.Println(err.Error())
			}
			Lows = append(Lows, vre)
			Low_ind = append(Low_ind, i+1)
		} else if Upper_d(j) && i+1 != len(str) {
			vre, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
			if err != nil {
				fmt.Println(err.Error())
			}
			Ups = append(Ups, vre)
			Up_ind = append(Up_ind, i+1)
		}
	}
}

// func Indexer(str [][]string) {
// 	for _, l := range str {
// 		c := make([]int, 0)
// 		u := make([]int, 0)
// 		lo := make([]int, 0)
// 		for i, j := range l {
// 			if Capper_d(j) && i+1 != len(l) {
// 				vre, err := strconv.Atoi(l[i+1][:len(l[i+1])-1])
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					os.Exit(1)
// 				}
// 				Caps = append(Caps, vre)
// 				c = append(c, i+1)
// 			} else if Lower_d(j) && i+1 != len(l) {
// 				vre, err := strconv.Atoi(l[i+1][:len(l[i+1])-1])
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					os.Exit(1)
// 				}
// 				Lows = append(Lows, vre)
// 				lo = append(lo, i+1)
// 			} else if Upper_d(j) && i+1 != len(l) {
// 				vre, err := strconv.Atoi(l[i+1][:len(l[i+1])-1])
// 				if err != nil {
// 					fmt.Println(err.Error())
// 					os.Exit(1)
// 				}
// 				Ups = append(Ups, vre)
// 				u = append(u, i+1)
// 			}
// 		}
// 		Cap_ind = append(Cap_ind, c)
// 		Low_ind = append(Low_ind, lo)
// 		Up_ind = append(Up_ind, u)
// 	}
// }
