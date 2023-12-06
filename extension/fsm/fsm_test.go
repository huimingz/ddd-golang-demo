package fsm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 使用者视角
//
// // 定义状态机
// fsm := NewFSM("current_state")
// fsm.AddEvent("state", "event", handler)
// err := fsm.Handle(ctx, result.status, result)

func TestFSM(t *testing.T) {
	fsm := NewFSM[string, string, any]("state")
	fsm.AddEvent("state", "event", func(ctx context.Context, event string, arg any) (string, error) {
		return "next_state", nil
	})

	err := fsm.Handle(context.Background(), "event", nil)
	assert.NoError(t, err)
	assert.Equal(t, "next_state", fsm.Current())

	err = fsm.Handle(context.Background(), "event", nil)
	assert.ErrorIs(t, err, ErrNoHandler)
	assert.Equal(t, "next_state", fsm.Current())
}
