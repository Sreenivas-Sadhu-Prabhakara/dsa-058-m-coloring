# M-Coloring Problem

> Striver SDE Sheet 058 · Day 10 Backtracking — assign colors vertex by vertex, trying each color against the already-colored neighbors, undoing on dead ends. Go, stdlib only, tested.

## Problem

Can a graph's vertices be colored with at most m colors so that no edge connects two vertices of the same color? (Decision only — this is the backtracking face of graph coloring, which is NP-complete in general.)

## Approach

Color vertices in index order. For vertex v, try each color 1..m that no already-colored neighbor holds; recurse, and strip the color when a branch dies. Two structural facts do the pruning: later vertices never constrain earlier choices (only colored neighbors are checked), and the first failure at any vertex walks back exactly to the most recent flexible choice. Tests pin the classic landmarks — triangles reject 2, even cycles accept 2, K4 rejects 3.

**Complexity:** O(mⁿ) worst case, O(n) state — the point of the exercise is the pruning pattern, not asymptotics.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
