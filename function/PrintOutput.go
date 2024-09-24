package lem_in

import (
	"fmt"
	"sync"
	"time"
)

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
						// an++
					} else {
						out = append(out, fmt.Sprintf("L%v-%v", an, paths[i][j]))
						// an++
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

			if j > len(paths[i]) {
				// stop = true
				// an = tempan
				break
			}

		}
		ants = ants - len(paths)
		tempan += stop
		// an = tempan
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

var (
	antsPaths = make(map[int][]string)
	output    = [][]string{}
	mu        sync.Mutex
	p         int = 0
	st        int = 0
)

func Output2(ants int, paths [][]string) {
	arr := Choose(paths)
	paths = Rec(arr)
	t := 0
	fmt.Println(paths)
	for i := 1; i <= ants; i++ {
		antsPaths[i] = paths[t]
		t++
		if t == len(paths) {
			t = 0
		}
	}

	fmt.Println(antsPaths)
	for in,_ := range antsPaths {
		go mouveAnt(in, p)
		if in%len(paths) != 0 {
			p++
		}
		

		time.Sleep(time.Second * 1 / 8)
	}
	for _, p := range output {
		fmt.Println(p)
	}
}

func mouveAnt(ant int, j int) {
	mu.Lock()
	fmt.Println(ant, j)
	for i := 0; i < len(antsPaths[ant]); i++ {
		if j < len(output) {
			output[j] = append(output[j], fmt.Sprintf("L%v-%v", p, antsPaths[ant][i]))
		} else {
			output = append(output, []string{fmt.Sprintf("L%v-%v", p, antsPaths[ant][i])})
		}
		j++
	}
	mu.Unlock()
}
