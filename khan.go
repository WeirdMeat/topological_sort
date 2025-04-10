package main

import (
	"errors"
)

type EdgeNode struct {
	v    int
	u    int
	next *EdgeNode
}

type EdgeList struct {
	head *EdgeNode
}

func KahnsAlgo(adj [][]int) ([]int, error) {
	e := EdgeList{nil}
	for v, row := range adj {
		for u, con := range row {
			if con == 1 {
				node := e.head
				e.head = &EdgeNode{v, u, node}
			}
		}
	}

	var l []int
	s := make(map[int]bool)
	for vv := 0; vv < len(adj); vv++ {
		ishead := true
		for v := 0; v < len(adj); v++ {
			if adj[v][vv] == 1 {
				ishead = false
				break
			}
		}
		if ishead {
			s[vv] = true
		}
	}

	for len(s) > 0 {
		n := GetRandomFromMap(s)
		delete(s, n)
		l = append(l, n)
		var m []int
		var prev *EdgeNode
		var cur *EdgeNode
		for prev, cur = nil, e.head; cur != nil; {
			if cur.v == n {
				m = append(m, cur.u)
				if prev == nil {
					e.head = cur.next
				} else {
					prev.next = cur.next
				}
			} else {
				prev = cur
			}
			cur = cur.next
		}
		for _, m_vertex := range m {
			new_head := true
			for node := e.head; node != nil; node = node.next {
				if node.u == m_vertex {
					new_head = false
					break
				}
			}
			if new_head {
				s[m_vertex] = true
			}
		}
		clear(m)

	}
	if e.head != nil {
		return nil, errors.New("graph has at least one cycle")
	} else {
		return l, nil
	}
}
