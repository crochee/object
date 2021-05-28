// Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.
// Description:
// Author: licongfu
// Create: 2021/5/26

// Package bloom
package bloom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"runtime"
	"testing"
	"time"
)

type fsy struct {
	ff *os.File
}

func (f *fsy) Close() error {
	fmt.Println("close")
	// Clear the finalizer.
	runtime.SetFinalizer(f, nil)
	fmt.Println("close1")
	return f.ff.Close()
}
func open(path string) (*fsy, error) {
	fs, err := os.OpenFile("../../../test/test_rs.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	f := &fsy{ff: fs}
	fmt.Println("open")
	runtime.SetFinalizer(f, (*fsy).Close)
	fmt.Println("open1")
	return f, nil
}

func TestF(t *testing.T) {
	fs, err := open("../../../test/test_rs.txt")
	require.NoError(t, err)
	defer fs.Close()
	go func() {
		buf := make([]byte, 4096)
		for {
			fs.ff.SetReadDeadline(time.Now().Add(6 * time.Second))
			n, err := fs.ff.Read(buf)
			if assert.Error(t, err) {
				assert.Equal(t, err, io.EOF)
			}
			require.Equal()
		}

	}()
}
