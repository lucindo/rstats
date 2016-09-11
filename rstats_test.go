package rstats_test

import (
	"math"
	"testing"

	"github.com/lucindo/rstats"
)

func equalf(a, b float64) bool {
	return !((a < b) || (a > b))
}

func inequalf(a, b float64) bool {
	return !equalf(a, b)
}

func TestEmpty(t *testing.T) {
	stats := rstats.New()
	zero := float64(0.0)

	if stats.Count() != 0 {
		t.Error("Count of empty stats should be zero")
	}

	if inequalf(stats.Min(), math.Inf(0)) {
		t.Error("Min of empty should be equal of math.Inf(0)")
	}

	if inequalf(stats.Max(), math.Inf(-1)) {
		t.Error("Max of empty should be equal of math.Inf(-1)")
	}

	if inequalf(stats.Mean(), zero) {
		t.Error("Mean of empty should be equal to zero")
	}

	if inequalf(stats.Variance(), zero) {
		t.Error("Variance of empty should be equal to zero")
	}

	if inequalf(stats.StandardDeviation(), zero) {
		t.Error("StandardDeviation of empty should be equal to zero")
	}

	if inequalf(stats.Skewness(), zero) {
		t.Error("Skewness of empty should be equal to zero")
	}

	if inequalf(stats.Kurtosis(), zero) {
		t.Error("Kurtosis of empty should be equal to zero")
	}

	//fmt.Println(stats)
}
