package leetcode

import "math"

func countPaths(n int, roads [][]int) int {
	edgesMap := make(map[int][][]int)
	visitedMap := make(map[int]bool)
	for index := range roads {
		edgesMap[roads[index][0]] = append(edgesMap[roads[index][0]], roads[index])
		edgesMap[roads[index][1]] = append(edgesMap[roads[index][1]], roads[index])
	}
	visitedMap[0] = true
	resultMap := make(map[int]int)
	getTheShortestPath(0, n-1, edgesMap, visitedMap, resultMap, 0)
	var minTime = math.MaxInt
	var minValue = 0
	for key := range resultMap {
		if key < minTime {
			minValue = resultMap[key]
			minTime = key
		}
	}
	return minValue
}

func getTheShortestPath(start int, end int, edgesMap map[int][][]int, visitedMap map[int]bool, resultMap map[int]int, time int) {
	for index := range edgesMap[start] {
		road := edgesMap[start][index]
		if road[1] == end || road[0] == end {
			resultMap[time+road[2]] = resultMap[time+road[2]] + 1
			continue
		}
		if !visitedMap[road[1]] {
			visitedMap[road[1]] = true
			getTheShortestPath(road[1], end, edgesMap, visitedMap, resultMap, time+road[2])
			visitedMap[road[1]] = false
		}
		if !visitedMap[road[0]] {
			visitedMap[road[0]] = true
			getTheShortestPath(road[0], end, edgesMap, visitedMap, resultMap, time+road[2])
			visitedMap[road[0]] = false
		}
	}
	return
}

// git config user.name "lomtom"
// git config user.email "lomtom@qq.com"
