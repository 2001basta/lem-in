package lem_in

func SortPath(ar [][]string) [][]string {
	i := 1
	for i < len(ar) {
		if len(ar[i]) < len(ar[i-1]) {
			ar[i], ar[i-1] = ar[i-1], ar[i]
			i = 0
		}
		i++
	}
	return ar
}
