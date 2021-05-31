// Package cache
package cache

import (
	"math/rand"
	"time"
)

type normTime struct {
	*rand.Rand
	mean          time.Duration
	stdDevPercent float64
}

func NormTime(mean time.Duration, stdDevPercent float64) *normTime {
	return &normTime{
		Rand:          rand.New(rand.NewSource(time.Now().Unix())),
		mean:          mean,
		stdDevPercent: stdDevPercent,
	}
}

func (n *normTime) Time() time.Duration {
	return time.Duration(float64(n.mean.Nanoseconds())*(100.0+n.stdDevPercent*n.Rand.NormFloat64())/100.0) * time.Nanosecond
}
