package eight

import (
	"github.com/cdlewis/advent-of-code/util"
)

type Tile struct {
	start       [2]int
	position    [2]int
	direction   [2]int
	antennaType byte
}

func Eight() int {
	uniquePoints := map[[2]int]struct{}{}
	grid := util.ToByteGrid(util.GetInput(8, false, ""))
	for idx, i := range grid {
		for jdx, j := range i {
			if j == '.' {
				continue
			}

			explore([2]int{idx, jdx}, grid, uniquePoints)
		}
	}

	return len(uniquePoints)
}

func explore(start [2]int, grid [][]byte, points map[[2]int]struct{}) {
	var q []Tile
	for _, d := range util.DirectionsDiagonal {
		newPosition := util.AddPoints(start, d)
		if !util.ValidPointCoordinate(newPosition, grid) {
			continue
		}

		q = append(q, Tile{
			start:       start,
			position:    newPosition,
			direction:   d,
			antennaType: grid[start[0]][start[1]],
		})
	}

	seenSameAntenna := false
	seen := map[[2]int]struct{}{}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := seen[curr.position]; ok {
			continue
		}
		seen[curr.position] = struct{}{}

		if grid[curr.position[0]][curr.position[1]] == curr.antennaType {
			seenSameAntenna = true

			moved := util.SubtractPoints(curr.position, curr.start)

			firstPoint := util.AddPoints(curr.position, moved)
			for util.ValidPointCoordinate(firstPoint, grid) {
				points[firstPoint] = struct{}{}
				firstPoint = util.AddPoints(firstPoint, moved)
			}

			secondPoint := util.SubtractPoints(curr.start, moved)
			for util.ValidPointCoordinate(secondPoint, grid) {
				points[secondPoint] = struct{}{}

				secondPoint = util.SubtractPoints(secondPoint, moved)
			}
		}

		for _, d := range util.DirectionsDiagonal {
			if curr.direction[0] != 0 && curr.direction[0] != d[0] {
				continue
			}

			if curr.direction[1] != 0 && curr.direction[1] != d[1] {
				continue
			}

			newPosition := util.AddPoints(curr.position, d)
			if !util.ValidPointCoordinate(newPosition, grid) {
				continue
			}

			q = append(q, Tile{
				start:       curr.start,
				position:    newPosition,
				direction:   curr.direction,
				antennaType: curr.antennaType,
			})
		}
	}

	if seenSameAntenna {
		points[start] = struct{}{}
	}
}
