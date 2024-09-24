package lem_in

var occuped = make(map[string]bool)

var slice = [][]string{}

func Rec(p [][]string) [][]string {
	paths := p
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			if !occuped[paths[i][j]] {
				slice = append(slice, paths[i])
				occuped[paths[i][j]] = true
				paths = append(paths[:i], paths[i+1:]...)
				i--
				break
			}
			break
		}
	}

	return slice
}

// func Rec1(p [][]string) {
// 	paths, rest := Rec(p)
// 	for i := 0; i < len(paths); i++ {
// 		for in := range paths[i] {
// 			for t := 0; t < len(paths); t++ {
// 				if t < len(paths) {
// 					for k := 0; k < len(paths[t]); k++ {
// 						if in+1 < len(paths[t]) {
// 							if paths[i][in] == paths[t][in+1] {
// 								fmt.Println("error", paths[i], paths[t])
// 								break
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
