package randomgen

import (
	"math/rand"
	"testing"
)

func Test_RandomGenInit(t *testing.T) {

	r := New(nil)

	if g, w := len(r.cumulativeProb), len(r.probabilities); g != w {
		t.Errorf("unexpected cumulative probability domain, wanted %v, got %v", w, g)
	}
	if g, w := len(r.cumulativeProb), len(r.randomNums); g != w {
		t.Errorf("cumulative probability domain length different to random number set length, wanted %v, got %v", w, g)
	}
}

func Test_CumulativeProbability(t *testing.T) {

	r := New(nil)

	expected := []float64{0.01, 0.31, 0.89, 0.99, 1.0}
	for i, v := range r.cumulativeProb {
		if g, w := v, expected[i]; g != w {
			t.Errorf("%v: got %v, want %v", i, g, w)
		}
	}

}

func Test_RandomGenNextNum(t *testing.T) {

	// source will create an rnd nuber generator with expected sequence
	src := rand.NewSource(0)

	// create rnd gen with source - ensuring fixed generated values
	r := New(src)

	expectedSequence := []int{2, 0, 1}

	for i, v := range expectedSequence {

		num := r.NextNum()

		if g, w := num, v; g != w {
			t.Errorf("%v: got %v, want %v", i, g, w)
		}
	}

}
