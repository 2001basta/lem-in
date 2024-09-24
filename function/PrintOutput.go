package lem_in

import "fmt"

func Output(ants int, paths [][]string) {
	arr := Choose(paths)
	paths = Rec(arr)
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
	l := 0
	stop := len(paths)
	an := 1

	// lem_in.Rec(paths)

	for l <= s {
		if ants <= 0 {
			break
		}
		// fmt.Println(ants)
		if ants < stop {
			stop = ants
		}

		temp := l
		tempan := an
		for {

			if i < len(paths) {
				if j < len(paths[i]) {
					if l < len(output) {
						output[l] = append(output[l], fmt.Sprintf("L%v-%v", an, paths[i][j]))
						an++
					} else {
						out = append(out, fmt.Sprintf("L%v-%v", an, paths[i][j]))
						an++
					}
				}
				i++
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

func printoutput(output [][]string) {
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			fmt.Print(output[i][j])
			if j < len(output[i])-1 {
				fmt.Print(" ")
			}
		}
		if i < len(output)-1 {
			fmt.Println()
		}

	}
}
