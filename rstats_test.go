package rstats_test

import (
	"math"
	"math/rand"
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
}

func TestSameValue(t *testing.T) {
	stats := rstats.New()
	zero := float64(0.0)
	random := rand.Float64()
	count := uint64(10000)

	for i := uint64(0); i < count; i++ {
		stats.Add(random)
	}

	if stats.Count() != count {
		t.Errorf("Count should be %d, got %d", count, stats.Count())
	}

	if inequalf(stats.Min(), random) {
		t.Error("Min should be equal to random")
	}

	if inequalf(stats.Max(), random) {
		t.Error("Max should be equal to random")
	}

	if inequalf(stats.Mean(), random) {
		t.Error("Mean should be equal to random")
	}

	if inequalf(stats.Variance(), zero) {
		t.Error("Variance should be equal to zero")
	}

	if inequalf(stats.StandardDeviation(), zero) {
		t.Error("StandardDeviation should be equal to zero")
	}

	if inequalf(stats.Skewness(), zero) {
		t.Error("Skewness should be equal to zero")
	}

	if inequalf(stats.Kurtosis(), zero) {
		t.Error("Kurtosis should be equal to zero")
	}
}
