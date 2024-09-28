package lem_in

func affectPaths(ants int, lastPath [][]string) map[int]int {
	ar := [][]string{}
	a := 1
	boo := false
	antsPaths := make(map[int]int)
	for i := 1; i < len(lastPath); i++ {
		diff := len(lastPath[i]) - len(lastPath[i-1])
		ar = append(ar, lastPath[i-1])
		for a <= diff {
			for k := 0; k < len(ar); k++ {
				antsPaths[k]++
				a++

			}
		}
	}
	for a <= ants {
		if !boo {
			ar = append(ar, lastPath[len(lastPath)-1])
			boo = true
		}
		for k := 0; k < len(ar); k++ {
			antsPaths[k]++
			a++
			if a == ants+1 {
				break
			}
		}
	}
	return antsPaths
}
