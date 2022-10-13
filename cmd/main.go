package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reload"
	"strings"
)

func main() {
	Args := os.Args[1:]
	if len(Args) != 2 {
		return
	}
	Sample := Args[0]
	Result := Args[1]
	var file1 []byte
	pwd, _ := os.Getwd()
	if _, err := os.Stat(Sample); err == nil {
		t, ERR1 := ioutil.ReadFile(pwd + "/" + Sample)
		if ERR1 != nil {
			fmt.Println(ERR1.Error())
		}
		file1 = t
	} else {
		fmt.Println(err.Error())
	}
	temp := strings.Split(string(file1), " ")
	// rese := make([][]string, 0)
	// for _, j := range temp {
	// 	tt := strings.Split(j, " ")
	// 	rese = append(rese, tt)
	// }
	str := temp
	reload.Numer(str)
	reload.Single(str)
	bimer := reload.Puncter(str)
	fmt.Println(bimer)

	reload.Indexer(str)
	reload.NonSingle(str)
	t := reload.Checker(str, bimer)
	str = t
	reload.Vowel(str)
	dir := os.Mkdir("result", 0700)
	if dir != nil {
		fmt.Println("Directory exists so:")
	} else {
		fmt.Println("Successfully created directory!")
	}
	file2, err := os.Create("result/" + Result)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("You can read it by --> cat result/result.txt")
	}
	defer file2.Close()
	res := ""
	for i, j := range str {
		for k, l := range j {
			if reload.Punct(l) && k+1 != len(j) {
				r := []rune(j)
				if !reload.Punct(r[k+1]) {
					res += string(r[k])
					res += string(" ")
					continue
				}
			}
			res += string(l)
		}
		if i+1 != len(str) {
			res += " "
		}
	}
	// fmt.Println(str)
	// for l, k := range str {
	// 	for i, j := range k {
	// 		// for h, f := range j {
	// 		// 	r := []rune(j)
	// 		// 	if reload.Punct(f) && h+1 != len(j) {
	// 		// 		if !reload.Punct(r[h+1]) {
	// 		// 			res += string(r[h])
	// 		// 			res += " "
	// 		// 			continue
	// 		// 		}
	// 		// 	}
	// 		// 	res += string(f)
	// 		// }
	// 		res += j
	// 		if i+1 != len(k) {
	// 			res += " "
	// 		}
	// 	}
	// 	if l+1 != len(str) {
	// 		res += "\n"
	// 	}
	// }
	_, err2 := file2.WriteString(res)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(reload.Caps)
	fmt.Println(reload.Lows)
	fmt.Println(reload.Up_ind)
	for _, j := range str {
		for i, _ := range j {
			fmt.Println(i)
		}
	}
}
