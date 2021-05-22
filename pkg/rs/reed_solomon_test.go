// Copyright 2021, The Go Authors. All rights reserved.
// Author: crochee
// Date: 2021/5/22

package rs

import (
	"os"
	"testing"

	"github.com/klauspost/reedsolomon"
	"github.com/stretchr/testify/require"
)

func TestRS(t *testing.T) {
	data, err := os.ReadFile("../../test/test_rs.txt")
	require.NoError(t, err)
	t.Logf("%s", data)
	enc, err := reedsolomon.New(2, 3)
	require.NoError(t, err)
	// Split the data into shards
	shards, err := enc.Split(data)
	require.NoError(t, err)
	t.Logf("%d", len(shards))
	for index, shard := range shards {
		t.Logf("index:%d %s", index, shard)
	}
	// Encode the parity set
	err = enc.Encode(shards)
	require.NoError(t, err)
	t.Logf("%d", len(shards))
	for index, shard := range shards {
		t.Logf("index:%d %s", index, shard)
	}
	// Verify the parity set
	ok, err := enc.Verify(shards)
	require.NoError(t, err)
	require.Equal(t, true, ok)
	// Delete two shards
	//shards[len(shards)-1] = nil
	//shards[len(shards)-2] = nil
	shards[len(shards)-3] = nil
	shards[len(shards)-4] = nil
	shards[len(shards)-5] = nil

	enc1, err := reedsolomon.New(2, 3)
	require.NoError(t, err)
	// Reconstruct the shards
	err = enc1.Reconstruct(shards)
	require.NoError(t, err)
	// Verify the data set
	ok, err = enc1.Verify(shards)
	require.NoError(t, err)
	require.Equal(t, true, ok)

	t.Logf("%d", len(shards))
	for index, shard := range shards {
		t.Logf("index:%d %s", index, shard)
	}
}
