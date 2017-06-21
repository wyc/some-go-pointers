package main

import (
	"math/rand"

	"github.com/pkg/profile"
)

// Do NOT go near 1,000,000. This crashed my Linux box.
const CYCLES = 100

func dumbNewValue() *int {
	s := make([]int, 65536, 65536)
	s[0] = rand.Int()
	return &s[0]
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	rvs := []*int{}
	for i := 0; i < CYCLES; i++ {
		rvs = append(rvs, dumbNewValue())
	}
}
