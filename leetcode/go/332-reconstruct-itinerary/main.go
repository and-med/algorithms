// idea - backtracking
// 1. set currentPos to JFK
// 2. for every possible edge thats not used, set currentPos to the edge, add the edge to used set
// 3. if all edges are used at current position, exit and try a different route
package main

import (
	"fmt"
	"slices"
)

var adjMap map[string][]string
var usedMap map[string][]int
var path []string
var ticketCnt int
var usedCnt int
var ans []string

// func findItinerary(tickets [][]string) []string {
// 	adjMap = make(map[string][]string, 0)
// 	usedMap = make(map[string][]bool, 0)
// 	path = make([]string, len(tickets)+1)
// 	ticketCnt = len(tickets)

// 	for _, ticket := range tickets {
// 		from, to := ticket[0], ticket[1]
// 		adjMap[from] = append(adjMap[from], to)
// 	}

// 	for from, destinations := range adjMap {
// 		slices.Sort(destinations)

// 		adjMap[from] = destinations
// 		usedMap[from] = make([]bool, len(destinations))
// 	}

// 	path[0] = "JFK"
// 	sol, found := findCycle("JFK", "JFK", 1)
// 	if !found {
// 		panic("not found cycle from JFK")
// 	}

// 	ans = make([]string, len(sol))
// 	copy(ans, sol)

// 	it := 0
// 	for len(ans) != ticketCnt+1 {
// 		for _, city := range ans {
// 			path[0] = city
// 			sol, found = findCycle(city, city, 1)
// 			if found {
// 				for i := len(ans) - 1; i >= 0; i-- {
// 					if ans[i] == city {
// 						sol = append(sol, ans[i+1:]...)
// 						ans = append(ans[0:i], sol...)
// 						break
// 					}
// 				}

// 				break
// 			}
// 		}
// 		it++
// 		if it == 10 {
// 			break
// 		}
// 	}

// 	return ans
// }

// func findCycle(start string, curr string, i int) ([]string, bool) {
// 	if curr == start && i != 1 {
// 		fmt.Println("found cycle", path[:i])
// 		return path[:i], true
// 	}

// 	if usedCnt == ticketCnt {
// 		fmt.Println("found path", path[:i])
// 		return path[:i], true
// 	}

// 	for idx, to := range adjMap[curr] {
// 		used := usedMap[curr]
// 		if !used[idx] {
// 			used[idx] = true
// 			path[i] = to
// 			usedCnt++
// 			if sol, found := findCycle(start, to, i+1); found {
// 				return sol, true
// 			}
// 			path[i] = ""
// 			usedCnt--
// 			used[idx] = false
// 		}
// 	}

// 	return nil, false
// }

func findItinerary(tickets [][]string) []string {
	adjMap = make(map[string][]string, 0)
	usedMap = make(map[string][]int, 0)
	path = make([]string, len(tickets)+1)

	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		adjMap[from] = append(adjMap[from], to)
	}

	for from, destinations := range adjMap {
		slices.Sort(destinations)
		counts := map[string]int{}

		for _, destination := range destinations {
			counts[destination]++
		}

		uniqueDestinations := []string{}

		for key := range counts {
			uniqueDestinations = append(uniqueDestinations, key)
		}

		slices.Sort(uniqueDestinations)
		sortedCounts := make([]int, len(uniqueDestinations))

		for idx, destination := range uniqueDestinations {
			sortedCounts[idx] = counts[destination]
		}

		adjMap[from] = uniqueDestinations
		usedMap[from] = sortedCounts
	}

	path[0] = "JFK"
	if !findItineraryRec("JFK", 1) {
		panic("itinerary not found")
	}

	return path
}

func findItineraryRec(from string, i int) bool {
	if i == len(path) {
		return true
	}

	for idx, to := range adjMap[from] {
		used := usedMap[from]
		if used[idx] > 0 {
			used[idx]--
			path[i] = to
			found := findItineraryRec(to, i+1)
			if found {
				return true
			}
			path[i] = ""
			used[idx]++
		}
	}

	return false
}

func main() {
	fmt.Println(findItinerary([][]string{[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"}, []string{"ATL", "JFK"}, []string{"ATL", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}}))
}
