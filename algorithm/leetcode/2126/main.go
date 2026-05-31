package main

import "slices"

func main() {

}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	for _, ast := range asteroids {
		if ast > mass {
			return false
		}
		mass += ast
	}

	return true
}
