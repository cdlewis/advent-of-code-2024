package six

import (
	"runtime"
	"sync"

	"github.com/cdlewis/advent-of-code/util"
)

var playerStates = map[byte][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

var nextState = map[[2]int][2]int{
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
	{0, -1}: {-1, 0},
}

type GuardState struct {
	Position  [2]int
	Direction [2]int
}

func (g GuardState) Serialize() [4]int {
	return [4]int{g.Direction[0], g.Direction[1], g.Position[0], g.Position[1]}
}

func Six() int {
	grid := util.ToByteGrid(util.GetInput(6, false, ""))

	startingPosition := getInitialState(grid)
	chunks := runtime.NumCPU()
	chunkSize := len(grid) / chunks
	results := make([]int, chunks)
	var wg sync.WaitGroup
	for i := range chunks {
		wg.Add(1)
		go func() {
			results[i] = search(
				startingPosition,
				i*chunkSize,
				i*chunkSize+chunkSize,
				grid,
			)
			wg.Done()
		}()
	}

	wg.Wait()

	locations := 0
	for _, i := range results {
		locations += i
	}

	return locations
}

func search(
	startingPosition GuardState,
	fromI int,
	toI int,
	grid [][]byte,
) int {
	locations := 0

	for idx := fromI; idx < toI; idx++ {
		for jdx, j := range grid[idx] {
			if idx == startingPosition.Position[0] && jdx == startingPosition.Position[1] {
				continue
			}

			if j == '#' {
				continue
			}

			currentPosition := startingPosition
			if hasCycle(currentPosition, grid, idx, jdx) {
				locations++
			}
		}
	}

	return locations
}

func getInitialState(grid [][]byte) GuardState {
	for idx, i := range grid {
		for jdx, j := range i {
			if state, ok := playerStates[j]; ok {
				return GuardState{
					Direction: state,
					Position:  [2]int{idx, jdx},
				}
			}
		}
	}

	panic("no guard present")
}

func hasCycle(
	state GuardState,
	grid [][]byte,
	newBarrierI int,
	newBarrierJ int,
) bool {
	seen := map[[4]int]struct{}{}

	for {
		serializedState := state.Serialize()
		if _, ok := seen[serializedState]; ok {
			return true
		}

		seen[serializedState] = struct{}{}

		nextI := state.Position[0] + state.Direction[0]
		nextJ := state.Position[1] + state.Direction[1]
		if !util.ValidCoordinate(nextI, nextJ, grid) {
			return false
		}

		if (nextI == newBarrierI && nextJ == newBarrierJ) || grid[nextI][nextJ] == '#' {
			state.Direction = nextState[state.Direction]
		} else {
			state.Position[0] = nextI
			state.Position[1] = nextJ
		}
	}
}
