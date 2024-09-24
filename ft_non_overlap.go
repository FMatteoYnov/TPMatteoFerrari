package tp

import (
	"sort"
)

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	nonOverlapCount := 1
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= lastEnd {
			nonOverlapCount++
			lastEnd = intervals[i][1]
		}
	}

	return len(intervals) - nonOverlapCount
}
