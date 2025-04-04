package main

import (
	"fmt"
	"math/rand"
)

func checkGraphAcyclic(adj [][]int, check []bool, v int) bool {
	if check[v] == true {
		return false
	}
	check[v] = true
	for u, con := range adj[v] {
		if con == 1 {
			if checkGraphAcyclic(adj, check, u) == false {
				return false
			}
		}
	}
	check[v] = false
	return true
}

func checkGraphCompletion(adj [][]int) bool {
	var GoFromVertex func([][]int, []bool, int)
	GoFromVertex = func(adj [][]int, check []bool, v int) {
		if check[v] == true {
			return
		}
		check[v] = true
		for u, con := range adj[v] {
			if con == 1 {
				GoFromVertex(adj, check, u)
			}
		}
		for vv := 0; vv < len(adj); vv++ {
			if adj[vv][v] == 1 {
				GoFromVertex(adj, check, vv)
			}
		}
	}
	check := make([]bool, len(adj))
	GoFromVertex(adj, check, 0)
	for _, vert := range check {
		if vert == false {
			return false
		}
	}
	return true
}

func CreateGraph(vs int) [][]int {
	adj := make([][]int, vs)
	for i := range vs {
		adj[i] = make([]int, vs)
	}
	for {
		v := rand.Intn(vs)
		u := rand.Intn(vs - 1)
		if v <= u {
			u += 1
		}
		adj[v][u] = 1
		check := make([]bool, vs)
		if checkGraphAcyclic(adj, check, v) == false {
			adj[v][u] = 0
			continue
		}
		if checkGraphCompletion(adj) {
			break
		}
	}
	return adj
}

func PrintGraph(adj [][]int) {
	for i, row := range adj {
		fmt.Printf("%v -> {", i)
		for j := range row {
			if row[j] == 1 {
				fmt.Printf("%v, ", j)
			}
		}
		fmt.Println("}")
	}
}
