package test

import (
	"errors"
	"studyGO/src/tools"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func findPrime(end int) ([]int, error) {
	if end < 2 {
		return nil, errors.New("end must be bigger than 2")
	}

	primes := []int{2}

out:
	for i := 3; i <= end; i++ {
		for _, prime := range primes {
			if i%prime == 0 {
				continue out
			}
		}
		primes = append(primes, i)
	}
	return primes, nil
}

func TestFindPrimes(t *testing.T) {
	convey.Convey("find primes in 10", t, func() {
		primes, _ := findPrime(100)
		expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

		convey.So(primes, tools.SliceEqual, expected)
	})
}
