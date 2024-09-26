package lem_in

import (
	"fmt"
	"slices"
	"strings"
)

func Output(ants int, paths [][]string) {
	arr := Choose(paths)
	paths = ChosePath(arr)
	fmt.Println(paths)
	fmt.Println("ttttttttttt", paths)
	output := [][]string{}
	out := []string{}
	i := 0
	j := 0
	s := 0
	if ants%len(paths) == 0 {
		s = ants / len(paths)
	} else {
		s = ants/len(paths) + 1
	}
	fmt.Println(ants, s)
	l := 0
	stop := len(paths)
	an := 1
	for l <= s {
		if ants <= 0 {
			break
		}

		if ants < stop {
			stop = ants
		}

		temp := l
		tempan := an
		for {

			if i < len(paths) {
				if j < len(paths[i]) {
					if l < len(output) {
						// fmt.Println("fffffffffffffffffffffffffff",l,i,j)
						output[l] = append(output[l], fmt.Sprintf("L%v-%v", an, paths[i][j]))
						an++
						fmt.Println("ttttttttttttttttttt", output[l])
						fmt.Println("gggggggggggggggggggg", output)
					} else {
						out = append(out, fmt.Sprintf("L%v-%v", an, paths[i][j]))
						an++
						fmt.Println("yyyyyyyyyyyyyyyyyyyyy", out)
					}
				}

				i++
				fmt.Println("fffffffffffffff", out, i, stop)
				if i == stop {

					if l >= len(output) {
						tempsl := make([]string, len(out))
						copy(tempsl, out)
						output = append(output, tempsl)
						out = nil
					}
					an = tempan
					i = 0
					j++
					l++
				}
			}
			if i < len(paths) {
				if j > len(paths[i]) {
					// stop = true
					an = tempan
					break
				}
			}

		}
		ants = ants - len(paths)
		tempan += stop
		an = tempan
		i = 0
		j = 0
		temp++
		l = temp
	}
	printoutput(output)
}

var (
	output  = [][]string{}
	occuped = make(map[int][]string)
	// lastPath [][]string
)

func Pprint(ants int, paths [][]string) {
	arr := Choose(paths)
	paths = ChosePath(arr)
	// lastPath = paths
	affectPaths(ants, paths)

	pindex := 0
	j := 0
	for i := 1; i <= ants; i++ {

		moveAnts(i, antsPaths[i], j)

		pindex++
		if pindex == len(paths) {
			j++
			pindex = 0
		}

	}

	// fmt.Println(output)
	printoutput(output)
	fmt.Println(len(output))
}

func printoutput(output [][]string) {
	for i := 0; i < len(output); i++ {
		fmt.Println(strings.Join(output[i], " "))
	}
}

func moveAnts(ant int, path []string, j int) bool {
	if slices.Index(occuped[j], path[0]) != -1 {
		return false
	}
	for i := 0; i < len(path); i++ {
		toPrint := fmt.Sprintf("L%v-%v", ant, path[i])
		if j < len(output) {
			output[j] = append(output[j], toPrint)
			occuped[j] = append(occuped[j], path[i])
			j++
		} else {
			output = append(output, []string{toPrint})
			occuped[j] = append(occuped[j], path[i])
			j++
		}
	}
	return true
}

var antsPaths = make(map[int][]string)

// a         int = 1

func affectPaths(ants int, lastPath [][]string) {
	ar := [][]string{}
	a := 1
	boo := false

	for i := 1; i < len(lastPath); i++ {
		diff := len(lastPath[i]) - len(lastPath[i-1])
		ar = append(ar, lastPath[i-1])
		for a <= diff {
			for k := 0; k < len(ar); k++ {
				antsPaths[a] = ar[k]
				a++

			}
		}
		// fmt.Println(antsPaths)
	}
	for a <= ants {
		if !boo {
			ar = append(ar, lastPath[len(lastPath)-1])
			boo = true
		}
		for k := 0; k < len(ar); k++ {
			antsPaths[a] = ar[k]
			a++
			if a == ants+1 {
				break
			}
		}
	}

	// fmt.Println(antsPaths)
}
