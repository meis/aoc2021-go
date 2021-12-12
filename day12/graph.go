package main

import "strings"

type graph map[string]node
type node map[string]bool

func newGraph(links []string) graph {
	g := make(graph)

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

func (g graph) getPathsFrom(cave string, maxSmallCaveVisits int, previousPath path) []path {
	previousPath = append(previousPath, cave)
	var paths []path

	if cave == "end" {
		paths = append(paths, previousPath)
		return paths
	} else if cave == "start" && len(previousPath) > 1 {
		return paths
	} else if cave[0] > 97 && previousPath.countDoubleSmallCaves() > maxSmallCaveVisits {
		return paths
	} else {
		for k := range g[cave] {
			paths = append(paths, g.getPathsFrom(k, maxSmallCaveVisits, previousPath)...)
		}

		return paths
	}
}
