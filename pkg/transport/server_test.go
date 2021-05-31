package transport

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type appPanic struct {
}

func (a appPanic) Start() error {
	panic("implement me")
}

func (a appPanic) Stop() error {
	panic("implement me")
}

type appError struct {
}

func (a appError) Start() error {
	return errors.New("start error")
}

func (a appError) Stop() error {
	return errors.New("stop error")
}

type appOk struct {
}

func (a appOk) Start() error {
	return nil
}

func (a appOk) Stop() error {
	return nil
}

func TestNewApp(t *testing.T) {
	testList := []struct {
		name string
		context.Context
		input    []AppServer
		expected bool
	}{
		{
			name: "panic",
			input: []AppServer{
				appPanic{},
			},
			expected: true,
		},
		{
			name: "error",
			input: []AppServer{
				appError{},
			},
			expected: true,
		},
		{
			name:    "ok",
			Context: context.Background(),
			input: []AppServer{
				appOk{},
			},
			expected: false,
		},
	}
	for _, data := range testList {
		t.Run(data.name, func(t *testing.T) {
			var (
				ctx        context.Context
				cancelFunc context.CancelFunc
			)
			if data.Context != nil {
				ctx, cancelFunc = context.WithCancel(data.Context)
			} else {
				ctx, cancelFunc = context.WithCancel(context.Background())
			}
			s := shutdown(func(ctx context.Context) error {
				time.Sleep(5 * time.Second)
				cancelFunc()
				return nil
			})
			a := NewApp(Context(ctx), Signal(), s, Servers(data.input...))
			if data.expected {
				assert.Error(t, a.Run())
			} else {
				assert.NoError(t, a.Run())
			}
		})
	}
}
