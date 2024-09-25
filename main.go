package main

import (
	lem_in "lem_in/function"
)

func main() {
	// fmt.Println(lem_in.ReadData("file.txt"))

	ants, links, start, end := lem_in.ReadData("file.txt")
	graph := lem_in.ConvertToGraph(links)
	// fmt.Println(graph)
	lem_in.Dfs(end, graph, []string{start})
	// fmt.Println(lem_in.Paths)
	// fmt.Println(lem_in.SortPath(lem_in.Paths))
	// lem_in.Paths = lem_in.Choose(lem_in.SortPath(lem_in.Paths))
	// ar := lem_in.Choose(lem_in.SortPath(lem_in.Paths))
	// fmt.Println("ccccccccccccc",ar)
	// // lem_in.AntsMouve(numAnt,lem_in.Paths)
	// fmt.Println(lem_in.Output2(ants, lem_in.Paths))
	lem_in.Printlnlinks(lem_in.Paths, lem_in.Output2(ants, lem_in.Paths))
}
