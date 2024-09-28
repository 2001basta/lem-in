package lem_in

import (
	"fmt"
	"strconv"
)

var Ant int

func Output2(ants int, paths [][]string) {
	Ant = ants
	arr := Choose(paths)
	paths = ChosePath(arr)
	paths = SortPath(paths)
	fmt.Println(paths)
	rule := make([][]string, len(paths))
	memoir := 1
	for i := 0; i < len(paths); i++ {
		ar := make([]string, len(paths[i]))
		rule[i] = ar
	}
	m := affectPaths(Ant, paths)
	str := ""
	boo := false
	cont := -1
	for {
		if len(m) > cont+1 {
			for a := range m {
				if m[a] > 0 {
					cont++
					m[a]--
				}
			}
		}
		x := 0
		for j := 0; j < len(paths); j++ {
			if memoir <= ants {
				str = "L" + strconv.Itoa(memoir) + "-"
				memoir++
			}
			//fmt.Println(rule[j],cont,j)

			rule[j] = mouveAnt(rule[j])
			if x <= cont && cont > -1 {
				rule[x][0] = str
				str = ""
				x++
			}
		}

		for z := 0; z < len(rule); z++ {
			for n := 0; n < len(rule[z]); n++ {
				if rule[z][n] != "" {
					fmt.Print(rule[z][n] + paths[z][n] + " ")
					boo = true
				}
			}
		}

		if !boo {
			break
		}
		fmt.Println()
		boo = false
		cont = -1

	}
}

// ///////////////////////////////////////////////////////////
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
