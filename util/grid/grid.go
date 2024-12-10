package grid

import (
	"fmt"
	"math"
	"strings"

	"github.com/cdlewis/advent-of-code/util/cast"
)

type Point [2]int

func (p Point) Add(another Point) Point {
	return AddPoints(p, another)
}

func (p Point) Subtract(another Point) Point {
	return SubtractPoints(p, another)
}

type Grid[T any] [][]T

func (g Grid[T]) ValidPoint(point Point) bool {
	return ValidPointCoordinate(point, g)
}

func (g Grid[T]) Get(point Point) T {
	return g[point[0]][point[1]]
}

func (g Grid[T]) GetAdjacent(point Point) []Point {
	var result []Point
	for _, i := range Directions {
		newPoint := point.Add(i)
		if g.ValidPoint(newPoint) {
			result = append(result, newPoint)
		}
	}
	return result
}

func ValidCoordinate[U any](i int, j int, grid [][]U) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func ValidPointCoordinate[U any](point [2]int, grid [][]U) bool {
	return ValidCoordinate(point[0], point[1], grid)
}

var Directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

var DirectionsDiagonal = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

var DirectionsDiagonalGrouped = [][][2]int{
	{{-1, -1}, {-1, 0}, {-1, 1}},
	{{1, -1}, {1, 0}, {1, 1}},
	{{1, -1}, {0, -1}, {-1, -1}},
	{{-1, 1}, {0, 1}, {1, 1}},
}

var Directions3D = [][]int{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func ShortestUnweightedPath[U any](graph [][]U, start [2]int, isEnd func(x [2]int) bool, validatePath func(x [2]int, y [2]int) bool) (int, bool) {
	steps := 0
	stack := [][2]int{start}
	visited := map[[2]int]bool{}

	for len(stack) > 0 {
		newStack := [][2]int{}

		for len(stack) > 0 {
			curr := stack[0]
			stack = stack[1:]

			if isEnd(curr) {
				return steps, true
			}

			if visited[curr] {
				continue
			}

			visited[curr] = true

			for _, d := range Directions {
				nextCoord := [2]int{curr[0] + d[0], curr[1] + d[1]}
				if ValidCoordinate(curr[0]+d[0], curr[1]+d[1], graph) && validatePath(curr, nextCoord) {
					newStack = append(newStack, nextCoord)
				}
			}
		}

		stack = newStack
		steps++
	}

	return -1, false
}

func ToGrid(s string) [][]int {
	lines := strings.Split(s, "\n")
	result := make([][]int, 0, len(lines))

	for _, l := range lines {
		line := make([]int, 0, len(l))
		for _, j := range l {
			line = append(line, cast.ToInt(j))
		}
		result = append(result, line)
	}

	return result
}

func ToByteGrid(s string) [][]byte {
	lines := strings.Split(s, "\n")
	result := make([][]byte, 0, len(lines))

	for _, l := range lines {
		result = append(result, []byte(l))
	}

	return result
}

func ValidCoordinate3D[U any](i, j, k int, space [][][]U) bool {
	if i < 0 || i >= len(space) {
		return false
	}

	if j < 0 || j >= len(space[i]) {
		return false
	}

	if k < 0 || k >= len(space[i][j]) {
		return false
	}

	return true
}

func AddPoints(x, y [2]int) [2]int {
	return [2]int{x[0] + y[0], x[1] + y[1]}
}

func SubtractPoints(x, y [2]int) [2]int {
	return [2]int{x[0] - y[0], x[1] - y[1]}
}

func BoundingBox[U any](graph map[[2]int]U) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for pos := range graph {
		minY = min(minY, pos[0])
		maxY = max(maxY, pos[0])
		minX = min(minX, pos[1])
		maxX = max(maxX, pos[1])
	}
	return minX, minY, maxX, maxY
}

func Print[U any](graph map[[2]int]U) {
	minX, minY, maxX, maxY := BoundingBox(graph)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if _, ok := graph[[2]int{i, j}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
