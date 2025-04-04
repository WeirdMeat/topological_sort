package main

import (
	"errors"
)

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func DepthFirstSearch(adj [][]int) ([]int, error) {
	var err error
	vs := len(adj)
	visited := make([]int, vs)
	var l []int

visiting:
	for n := 0; n < vs; n++ {
		if visited[n] != 2 {

			l, err = visit(n, visited, adj, l)

			if err != nil {
				panic(err)
			}
			n = 0
			goto visiting
		}
	}

	reverse(l)

	return l, nil
}

func visit(n int, visited []int, adj [][]int, l []int) ([]int, error) {
	var err error

	if visited[n] == 2 {
		return l, nil
	}
	if visited[n] == 1 {
		return nil, errors.New("graph has a cycle")
	}
	visited[n] = 1
	for m, con := range adj[n] {
		if con == 1 {
			l, err = visit(m, visited, adj, l)
			if err != nil {
				panic(err)
			}
		}
	}

	visited[n] = 2
	l = append(l, n)

	return l, nil
}
