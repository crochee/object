// Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.
// Description:
// Author: licongfu
// Create: 2021/5/31

// Package cache
package cache

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type NormTime struct {
	Mean          time.Duration
	StdDevPercent float64
}

func (n *NormTime) Time() time.Duration {
	return time.Duration(float64(n.Mean.Nanoseconds())*(100.0+n.StdDevPercent*rand.NormFloat64())/100.0) * time.Nanosecond
}
