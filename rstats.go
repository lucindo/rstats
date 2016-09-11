// Package rstats provides a simple and fast way to compute statistics
// online
// Code from: http://www.johndcook.com/blog/standard_deviation/
// and http://www.johndcook.com/blog/skewness_kurtosis/
package rstats

import (
	"fmt"
	"math"
	"sync"
)

// Stats holds the minimal set of variables to calculate
// online statiscs
type Stats struct {
	// only private fields
	lock           sync.RWMutex
	count          uint64
	min, max       float64
	m1, m2, m3, m4 float64
}

// New returs a new instance of Stats
func New() *Stats {
	stats := &Stats{}
	stats.Reset()
	return stats
}

// Add a data point to calculate statistics
func (stats *Stats) Add(value float64) {
	stats.lock.Lock()
	defer stats.lock.Unlock()

	stats.min = math.Min(stats.min, value)
	stats.max = math.Max(stats.max, value)

	lastCount := stats.count
	stats.count++
	delta := value - stats.m1
	deltaCount := delta / float64(stats.count)
	deltaCountSquare := deltaCount * deltaCount
	termLast := delta * deltaCount * float64(lastCount)

	stats.m1 += deltaCount
	stats.m4 += termLast*deltaCountSquare*(float64(stats.count*stats.count)-3*float64(stats.count)+3) + 6*deltaCountSquare*stats.m2 - 4*deltaCount*stats.m3
	stats.m3 += termLast*deltaCount*float64(stats.count-2) - 3*deltaCount*stats.m2
	stats.m2 += termLast
}

// Reset all values
func (stats *Stats) Reset() {
	stats.lock.Lock()
	defer stats.lock.Unlock()

	stats.count = 0
	stats.min = math.Inf(0)
	stats.max = math.Inf(-1)
	stats.m1, stats.m2, stats.m3, stats.m4 = 0.0, 0.0, 0.0, 0.0
}

// Count return the number of elements computed so far
func (stats *Stats) Count() uint64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	return stats.count
}

// Min returns the min value added so far
func (stats *Stats) Min() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	return stats.min
}

// Max returns the man value added so far
func (stats *Stats) Max() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	return stats.max
}

// Mean returns the mean of values added so far
func (stats *Stats) Mean() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	return stats.m1
}

// Variance returns the variance of values added so far
func (stats *Stats) Variance() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	if stats.count > 1 {
		return stats.m2 / float64(stats.count-1.0)
	}
	return 0.0
}

// StandardDeviation returns the standard deviation of values added so far
func (stats *Stats) StandardDeviation() float64 {
	// no locks
	return math.Sqrt(stats.Variance())
}

// Skewness returns the skewness of values added so far
func (stats *Stats) Skewness() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	if stats.count > 0 {
		return math.Sqrt(float64(stats.count)) * stats.m3 / math.Pow(stats.m2, 1.5)
	}
	return 0.0
}

// Kurtosis returns the kurtosis of values added so far
func (stats *Stats) Kurtosis() float64 {
	stats.lock.RLock()
	defer stats.lock.RUnlock()

	if stats.count > 0 {
		return float64(stats.count)*stats.m4/(stats.m2*stats.m2) - 3.0
	}
	return 0.0
}

// String returns a printable string summary of stats
func (stats *Stats) String() string {
	return fmt.Sprintf("count %d min %.2f max %.2f mean %.2f (std dev %.3f variance %.2f) [skewness %.2f kurtosis %.2f]",
		stats.Count(), stats.Min(), stats.Max(), stats.Mean(), stats.StandardDeviation(), stats.Variance(), stats.Skewness(), stats.Kurtosis())
}

// TODO: Implement Linear Regression: http://www.johndcook.com/blog/running_regression/
