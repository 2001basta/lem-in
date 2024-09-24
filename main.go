package main

import (
	lem_in "lem_in/function"
)

var out = make(chan string)

func main() {
	// fmt.Println(lem_in.ReadData("file.txt"))

	ants, links, start, end := lem_in.ReadData("file.txt")
	graph := lem_in.ConvertToGraph(links)
	lem_in.Dfs(end, graph, []string{start})
	// fmt.Println(lem_in.SortPath(lem_in.Paths))
	// lem_in.Paths = lem_in.ChosePath(start, lem_in.SortPath(lem_in.Paths))
	// fmt.Println(lem_in.Paths)
	// // lem_in.AntsMouve(numAnt,lem_in.Paths)
	lem_in.Output(ants, lem_in.Paths)
}
