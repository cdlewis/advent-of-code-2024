package one

import (
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func One() int {
	pairs := strings.Split(util.GetInput(1, false, ""), "\n")

	var leftList []int
	rightListFreq := map[int]int{}
	for _, i := range pairs {
		tokens := strings.Split(i, "   ")
		leftList = append(leftList, util.ToInt(tokens[0]))
		rightListFreq[util.ToInt(tokens[1])]++
	}

	totalDist := 0
	for _, i := range leftList {
		totalDist += util.Abs(i * rightListFreq[i])
	}

	return totalDist
}
