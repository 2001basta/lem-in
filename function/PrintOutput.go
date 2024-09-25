package lem_in

import (
	"fmt"
	"strconv"
)

func Output2(ants int, paths [][]string) [][]string {
	arr := Choose(paths)
	paths = ChosePath(arr)

	// fmt.Println(paths)
	rule := make([][]string, len(paths))
	memoir := 1
	for i := 0; i < len(paths); i++ {
		ar := make([]string, len(paths[i]))
		rule[i] = ar
	}
	k := 0
	str := ""
	all_rule := [][]string{}
	for {
		if k > len(paths)*ants+1 {
			break
		}
		for j := 0; j < len(paths); j++ {
			if memoir <= ants {
				str = "L" + strconv.Itoa(memoir) + "-"
				memoir++
			}
			rule[j] = mouveAnt(rule[j])
			rule[j][0] = str
			str = ""
			all_rule = append(all_rule, rule[j])
			k++
		}
	}
	return all_rule
}

func mouveAnt(ar []string) []string {
	arr := make([]string, len(ar))
	arr[0] = ""
	for i := 0; i < len(ar); i++ {
		if i < len(arr)-1 {
			arr[i+1] = ar[i]
		}
	}
	return arr
}

func Printlnlinks(pat, ar [][]string) {
	arr := Choose(pat)
	pat = ChosePath(arr)
	fmt.Println(pat)
	a := 0
	for j := 0; j < len(ar); j++ {
		if !checkArry(ar[j]) {
			continue
		}
		for k := len(ar[j])-1; k >=0; k-- {
			if ar[j][k] != "" {
				if k < len(pat[a]) {
					fmt.Print(ar[j][k] + pat[a][k] + " ")
				}
			}
		}
		a++
		if a >= len(pat) {
			fmt.Println()
			a = 0
		}
	}
}

func checkArry(ar []string) bool {
	for _, i := range ar {
		if i != "" {
			return true
		}
	}
	return false
}
