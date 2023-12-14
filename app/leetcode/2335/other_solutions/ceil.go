package other_solutions

import (
	"fmt"
	"math"
	"sort"
)

func FillCupByCeil(amount []int) int {
	sort.Slice(amount, func(i, j int) bool {
		return amount[j] < amount[i]
	})

	c := math.Ceil(float64(amount[0]+amount[1]+amount[2]) / 2)
	fmt.Printf("c:%f\n",c)
	if int(c) < amount[0] {
		return amount[0]
	}

	return int(c)
}
