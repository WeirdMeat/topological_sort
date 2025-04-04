package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	graph := CreateGraph(20)
	//PrintGraph(graph)

	start := time.Now()
	sorted, err := KahnsAlgo(graph)
	elapsed := time.Since(start)
	if err != nil {
		panic(err)
	}
	check, err := CheckSorted(graph, sorted)
	if err != nil {
		panic(err)
	}
	if check == false {
		panic(errors.New("Khan's algorithm failed to sort"))
	}
	fmt.Printf("Khan's Algorithm get result %v in %v\n", sorted, elapsed)

	start = time.Now()
	sorted, err = DepthFirstSearch(graph)
	elapsed = time.Since(start)
	if err != nil {
		panic(err)
	}
	check, err = CheckSorted(graph, sorted)
	if err != nil {
		panic(err)
	}
	if check == false {
		panic(errors.New("Depth-first search failed to sort"))
	}
	fmt.Printf("Depth-first search get result %v in %v\n", sorted, elapsed)
}
