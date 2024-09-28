package main

import (
	"fmt"
	lem_in "lem_in/function"
)


func main() {
	ants, links, start, end := lem_in.ReadFile()
	graph := lem_in.ConvertToGraph(links)
	lem_in.Defs(end, graph, []string{start})
	fmt.Println(len(lem_in.Paths))

	errorMsg := lem_in.CheckErrors(ants, start, end, len(lem_in.Paths))

	if errorMsg == "error" {
		return
	}

	lem_in.PrintAnts(ants, lem_in.Paths)
}
