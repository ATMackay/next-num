package randomgen

import (
	"math"
	"math/rand"
)

var (
	rndList = []int{-1, 0, 1, 2, 3}
	prbList = []float64{0.01, 0.3, 0.58, 0.1, 0.01}
)

// randomGen
type randomGen struct {
	src            *rand.Rand
	randomNums     []int
	probabilities  []float64
	cumulativeProb []float64
}

// New initializes a randomGen random number generator instance
func New(src rand.Source) *randomGen {

	r := &randomGen{
		randomNums:    rndList,
		probabilities: prbList,
	}

	if src != nil {
		r.src = rand.New(src)
	}

	cum := 0.0
	for _, v := range r.probabilities {
		cum += v
		r.cumulativeProb = append(r.cumulativeProb, math.Round(cum*100)/100)
	}

	return r
}

func (r *randomGen) NextNum() (n int) {

	// Use rand package to create random value between 0 and 1
	var randFloat float64
	if r.src == nil {
		randFloat = rand.Float64()
	} else {
		randFloat = r.src.Float64()
	}

	// Find location in the cumulative probability distribution
	// and return value at corresponding index
	for i, v := range r.cumulativeProb {
		if randFloat < v {
			n = r.randomNums[i]
			break
		}
	}

	return
}
