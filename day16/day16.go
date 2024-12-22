package day16

import (
	"fmt"

	"github.com/alexchro93/aoc-2024/utils"
)

// many thanks to inspiration from https://github.com/arjunpathak072/aoc-2024/blob/main/day-16/main.go

const max = int(^uint(0) >> 1)

const wall rune = '#'
const start rune = 'S'
const target rune = 'E'

type direction int

const (
	north direction = iota
	south
	east
	west
)

var directions = []direction{north, south, east, west}

type tile struct {
	point utils.Point
	dir   direction
}

type path struct {
	head   tile
	points []utils.Point
	score  int
}

func parseInput() ([][]rune, error) {
	lines, err := utils.ReadAllLines("day16/day16.txt")
	if err != nil {
		return nil, err
	}
	grid := make([][]rune, len(lines))
	for j, line := range lines {
		grid[j] = []rune(line)
	}
	return grid, nil
}

func neighbors(t tile) []tile {
	p := t.point
	neighbors := make([]tile, len(directions))
	for i, dir := range directions {
		if dir == north {
			neighbors[i] = tile{point: utils.Point{X: p.X, Y: p.Y - 1}, dir: dir}
		} else if dir == south {
			neighbors[i] = tile{point: utils.Point{X: p.X, Y: p.Y + 1}, dir: dir}
		} else if dir == east {
			neighbors[i] = tile{point: utils.Point{X: p.X + 1, Y: p.Y}, dir: dir}
		} else if dir == west {
			neighbors[i] = tile{point: utils.Point{X: p.X - 1, Y: p.Y}, dir: dir}
		}
	}
	return neighbors
}

func Run() {
	grid, err := parseInput()
	if err != nil {
		return
	}

	begin := utils.Point{}
	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] == start {
				begin = utils.Point{X: i, Y: j}
			}
		}
	}

	minScore := max
	minScoreVisited := make(map[int][]utils.Point)
	visited := make(map[tile]int)
	q := []path{{head: tile{point: begin, dir: east}, points: []utils.Point{begin}, score: 0}}

	// Part One and Two
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.score > minScore {
			continue
		}

		if grid[curr.head.point.Y][curr.head.point.X] == target {
			if curr.score <= minScore {
				minScore = curr.score
				minScoreVisited[curr.score] = append(minScoreVisited[curr.score], curr.points...)
			}
			continue
		}

		for _, neighbor := range neighbors(curr.head) {
			if grid[neighbor.point.Y][neighbor.point.X] == wall {
				continue
			}

			score := curr.score + 1
			if neighbor.dir != curr.head.dir {
				score += 1000
			}

			if prev, has := visited[neighbor]; has && prev < score {
				continue
			}

			visited[neighbor] = score

			nPath := make([]utils.Point, 0)
			nPath = append(nPath, curr.points...)
			nPath = append(nPath, neighbor.point)

			q = append(q, path{head: neighbor, points: nPath, score: score})
		}
	}

	fmt.Println("Day 16 Part 1:", minScore)

	unique := make(map[utils.Point]bool)
	for _, p := range minScoreVisited[minScore] {
		unique[p] = true
	}

	fmt.Println("Day 16 Part 2:", len(unique))
}
