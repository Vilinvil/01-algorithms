package main

import "fmt"

// DifferenceSets return result mathematical operation (Set a) - (Set b)
func DifferenceSets(a map[string]struct{}, b map[string]struct{}) map[string]struct{} {
	res := a
	for i := range b {
		_, ok := a[i]
		if ok {
			delete(res, i)
		}
	}
	return res
}

// IntersectSets return result mathematical operation (Set a) & (Set b)
func IntersectSets(a map[string]struct{}, b map[string]struct{}) map[string]struct{} {
	res := make(map[string]struct{})
	for i := range a {
		_, ok := b[i]
		if ok {
			res[i] = struct{}{}
		}
	}
	return res
}

// This program implements a greedy algorithm for solving the set covering problem.
// Sets are maps of empty structures (to save memory)
func main() {
	// Init system of states. map[string]struct{} is mathematical set(множество)
	statesNeeded := make(map[string]struct{})
	statesNeeded["mt"] = struct{}{}
	statesNeeded["wa"] = struct{}{}
	statesNeeded["or"] = struct{}{}
	statesNeeded["id"] = struct{}{}
	statesNeeded["nv"] = struct{}{}
	statesNeeded["ut"] = struct{}{}
	statesNeeded["ca"] = struct{}{}
	statesNeeded["az"] = struct{}{}

	// RadioStations and what states he covered
	stations := make(map[string]map[string]struct{})
	stations["kone"] = map[string]struct{}{
		"id": {},
		"nv": {},
		"ut": {},
	}
	stations["ktwo"] = map[string]struct{}{
		"wa": {},
		"id": {},
		"mt": {},
	}
	stations["kthree"] = map[string]struct{}{
		"or": {},
		"nv": {},
		"ca": {},
	}
	stations["kfour"] = map[string]struct{}{
		"nv": {},
		"ut": {},
	}
	stations["kfive"] = map[string]struct{}{
		"ca": {},
		"az": {},
	}

	//Set of stations which we will choose
	finalStations := make(map[string]struct{})

	for len(statesNeeded) != 0 {
		// Choose best stations(who cowered bigger). And what it covers
		bestStation := ""
		statesCovered := map[string]struct{}{}
		for station, states := range stations {
			covered := IntersectSets(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		statesNeeded = DifferenceSets(statesNeeded, statesCovered)
		finalStations[bestStation] = struct{}{}
	}

	fmt.Println(finalStations)
}
