package main

import "fmt"

// CanColor reports whether the graph (adjacency matrix) can be colored with
// at most m colors so no edge joins two same-colored vertices.
func CanColor(adj [][]bool, m int) bool {
	n := len(adj)
	colors := make([]int, n) // 0 = uncolored, 1..m = colors
	var try func(v int) bool
	try = func(v int) bool {
		if v == n {
			return true
		}
		for c := 1; c <= m; c++ {
			ok := true
			for u := 0; u < n; u++ {
				if adj[v][u] && colors[u] == c {
					ok = false
					break
				}
			}
			if ok {
				colors[v] = c
				if try(v + 1) {
					return true
				}
				colors[v] = 0
			}
		}
		return false
	}
	return try(0)
}

func main() {
	// A triangle needs 3 colors.
	adj := [][]bool{
		{false, true, true},
		{true, false, true},
		{true, true, false},
	}
	fmt.Println(CanColor(adj, 2), CanColor(adj, 3))
}
