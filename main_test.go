package main

import "testing"

func matrix(n int, edges [][2]int) [][]bool {
	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = true
		adj[e[1]][e[0]] = true
	}
	return adj
}

func TestCanColor(t *testing.T) {
	triangle := matrix(3, [][2]int{{0, 1}, {1, 2}, {0, 2}})
	if CanColor(triangle, 2) {
		t.Error("triangle is not 2-colorable")
	}
	if !CanColor(triangle, 3) {
		t.Error("triangle is 3-colorable")
	}

	// Square (4-cycle) is bipartite: 2 colors suffice.
	square := matrix(4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}})
	if !CanColor(square, 2) {
		t.Error("even cycle is 2-colorable")
	}

	// K4 needs 4.
	k4 := matrix(4, [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}})
	if CanColor(k4, 3) {
		t.Error("K4 is not 3-colorable")
	}
	if !CanColor(k4, 4) {
		t.Error("K4 is 4-colorable")
	}

	// Edgeless graph: 1 color is enough; empty graph: trivially true.
	if !CanColor(matrix(3, nil), 1) {
		t.Error("edgeless graph is 1-colorable")
	}
	if !CanColor(matrix(0, nil), 0) {
		t.Error("empty graph is colorable with zero colors")
	}
}
