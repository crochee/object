// Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.
// Description:
// Author: licongfu
// Create: 2021/5/31

// Package cache
package cache

import (
	"testing"
	"time"
)

func TestNormTime_Time(t *testing.T) {
	tt := &NormTime{
		Mean:          5 * time.Second,
		StdDevPercent: 20,
	}
	for i := 0; i < 1000; i++ {
		t.Log(tt.Time())
	}
}

func BenchmarkNormTime_Time(b *testing.B) {
	tt := &NormTime{
		Mean:          5 * time.Second,
		StdDevPercent: 20,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt.Time()
	}
}
