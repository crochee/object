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
	dataShards := 6
	enc, err := reedsolomon.New(dataShards, len(data)/dataShards)
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
	// Verify the parity set
	ok, err := enc.Verify(shards)
	require.NoError(t, err)
	require.Equal(t, true, ok)
	// Delete two shards
	shards[2], shards[3] = nil, nil

	// Reconstruct the shards
	err = enc.ReconstructData(shards)
	require.NoError(t, err)
	// Verify the data set
	ok, err = enc.Verify(shards)
	require.NoError(t, err)
	require.Equal(t, true, ok)
	t.Logf("%d", len(shards))
	for index, shard := range shards {
		t.Logf("index:%d %s", index, shard)
	}
}
