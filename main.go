package main

import (
	lem_in "lem_in/function"
)

func main() {
	ants, links, start, end := lem_in.ReadData("file.txt")
	graph := lem_in.ConvertToGraph(links)
	lem_in.Dfs(end, graph, []string{start})
	lem_in.Printlnlinks(lem_in.Paths, lem_in.Output2(ants, lem_in.Paths))
}
