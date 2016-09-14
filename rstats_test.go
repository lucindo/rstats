package rstats_test

import (
	"math"
	"math/rand"
	"sync"
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
	stString := "count 0 min +Inf max -Inf mean 0.00 (std dev 0.000 variance 0.00)"

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

	if stats.String() != stString {
		t.Errorf("String of empty should be equal to:\n'%s'\nbut got:\n'%s'", stString, stats.String())
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

func TestStatsStruct(t *testing.T) {
	stats := rstats.New()
	random := rand.Float64()
	count := uint64(10000)

	for i := uint64(0); i < count; i++ {
		stats.Add(random)
	}

	statsStruct := new(rstats.StatsStruct)
	rstats.GetStatsStruct(statsStruct, stats)

	if statsStruct.Count != stats.Count() {
		t.Errorf("Count should be %d, got %d", statsStruct.Count, stats.Count())
	}

	if inequalf(statsStruct.Min, stats.Min()) {
		t.Errorf("Min should be %f, got %f", statsStruct.Min, stats.Min())
	}

	if inequalf(statsStruct.Max, stats.Max()) {
		t.Errorf("Max should be %f, got %f", statsStruct.Max, stats.Max())
	}

	if inequalf(statsStruct.Mean, stats.Mean()) {
		t.Errorf("Mean should be %f, got %f", statsStruct.Mean, stats.Mean())
	}

	if inequalf(statsStruct.Variance, stats.Variance()) {
		t.Errorf("Variance should be %f, got %f", statsStruct.Variance, stats.Variance())
	}

	if inequalf(statsStruct.StandardDeviation, stats.StandardDeviation()) {
		t.Errorf("StandardDeviation should be %f, got %f", statsStruct.StandardDeviation, stats.StandardDeviation())
	}

	if inequalf(statsStruct.Skewness, stats.Skewness()) {
		t.Errorf("Skewness should be %f, got %f", statsStruct.Skewness, stats.Skewness())
	}

	if inequalf(statsStruct.Kurtosis, stats.Kurtosis()) {
		t.Errorf("Kurtosis should be %f, got %f", statsStruct.Kurtosis, stats.Kurtosis())
	}
}

func TestConcurrency(t *testing.T) {
	stats := rstats.New()
	zero := float64(0.0)
	random := rand.Float64()
	count := uint64(1000)
	gorotines := 100
	var group sync.WaitGroup

	for g := 0; g < gorotines; g++ {
		group.Add(1)
		go func(stats *rstats.Stats, group *sync.WaitGroup) {
			for i := uint64(0); i < count; i++ {
				stats.Add(random)
			}
			group.Done()
		}(stats, &group)
	}
	group.Wait()

	if stats.Count() != count*uint64(gorotines) {
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
