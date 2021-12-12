package main

import "strings"

type Graph map[string]node
type node map[string]bool
type visitedCache []string

func NewGraph(links []string) Graph {
	g := make(Graph)

	for _, l := range links {
		parts := strings.Split(l, "-")
		from := parts[0]
		to := parts[1]

		_, foundFrom := g[from]
		if !foundFrom {
			g[from] = make(node)
		}

		g[from][to] = true

		_, foundTo := g[to]
		if !foundTo {
			g[to] = make(node)
		}

		g[to][from] = true
	}

	return g
}

func (g Graph) GetPathsFromStart(maxSmallCaveVisits int) int {
	return g.getPathsFromRec("start", maxSmallCaveVisits, visitedCache{})
}

func (g Graph) getPathsFromRec(cave string, maxSmallCaveVisits int, smallCavesVisited visitedCache) int {
	if cave[0] > 97 {
		smallCavesVisited = append(smallCavesVisited, cave)
	}

	if cave == "end" {
		return 1
	} else if cave == "start" && len(smallCavesVisited) > 1 {
		return 0
	} else if cave[0] > 97 && smallCavesVisited.countDoubleSmallCaves() > maxSmallCaveVisits {
		return 0
	} else {
		nextPaths := 0
		for next := range g[cave] {
			nextPaths += g.getPathsFromRec(next, maxSmallCaveVisits, smallCavesVisited)
		}

		return nextPaths
	}
}

func (cache visitedCache) countDoubleSmallCaves() int {
	small := make(map[string]int)
	count := 0

	for _, c := range cache {
		v := small[c]

		if v > 0 {
			count += v
		}

		small[c]++
	}

	return count
}
