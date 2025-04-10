package main

import "errors"

func CheckSorted(adj [][]int, sorted []int) (bool, error) {
	vs := len(sorted)
	if vs != len(adj) {
		return false, errors.New("mismatch of number of vertexes")
	}
	visited := make([]bool, len(sorted))
	for _, vertex := range sorted {
		if vertex < 0 || vertex >= vs {
			return false, errors.New("incorect vertex in sorted list")
		}
		for v := 0; v < vs; v++ {
			if adj[v][vertex] == 1 && visited[v] == false {
				return false, nil
			}
		}
		visited[vertex] = true
	}
	return true, nil
}

func GetRandomFromMap(m map[int]bool) int {
	for k := range m {
		return k
	}
	return -1
}
