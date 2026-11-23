package main

import "fmt"

func bestLap(runnerLaps map[string]int, name string) (int, bool) {
	val, ok := runnerLaps[name]

	return val, ok
}

func main() {
	laps := map[string]int{
		"Maya": 72,
		"Leo":  0,
		"Ada":  68,
	}
	fmt.Println(bestLap(laps, "Sam"))
}
