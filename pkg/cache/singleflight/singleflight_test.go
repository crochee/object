// Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.
// Description:
// Author: licongfu
// Create: 2021/5/25

// Package singleflight
package singleflight

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/require"
)

func TestGroup_Do(t *testing.T) {
	g := New()

	v, err := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})
	require.NoError(t, err)
	require.Equal(t, "bar", v)
}

func TestGroup_DoMany(t *testing.T) {
	g := New()
	c := make(chan string)
	var calls int32
	fn := func() (interface{}, error) {
		atomic.AddInt32(&calls, 1)
		return <-c, nil
	}

	const n = 100
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			v, err := g.Do("key", fn)
			require.NoError(t, err)
			require.Equal(t, "bar", v)
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond) // let goroutines above block
	c <- "bar"
	wg.Wait()
	require.Equal(t, int32(1), atomic.LoadInt32(&calls))
}

func TestGroup_Do2(t *testing.T) {

	for i := 0; i < 10000; i++ {
		t.Log(rand.NormFloat64()*float64(2) + 10)
	}

}

func BenchmarkGroup_Do(b *testing.B) {
	rand.Seed(time.Now().Unix())
	g := New()
	fn := func() (interface{}, error) {
		fmt.Println("test")
		return "test", nil
	}
	c := cache.New(50*time.Second, 75*time.Second)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, ok := c.Get("key"); !ok {
			v, err := g.Do("key", fn)
			if err != nil {
				b.FailNow()
			}
			c.Set("key", v, time.Duration(rand.NormFloat64()*float64(2)+10)*time.Second)
		}
	}
}
