package fsm

import (
	"context"
	"errors"
	"sync"

	"scan/extension/logz"
)

var ErrNoHandler = errors.New("no handler")

type State string
type Event string

type Handler[S, E comparable, T any] func(ctx context.Context, event E, arg T) (S, error)
type CallBack[S, E comparable, T any] func(ctx context.Context, event E, nextState S, arg T) error

type FSM[S, E comparable, T any] struct {
	currentState        S
	mutex               sync.Mutex
	handlers            map[S]map[E]Handler[S, E, T]
	globalAfterCallback CallBack[S, E, T]
}

func NewFSM[S, E comparable, T any](initState S) *FSM[S, E, T] {
	return &FSM[S, E, T]{
		currentState: initState,
		handlers:     make(map[S]map[E]Handler[S, E, T]),
	}
}

func (f *FSM[S, E, T]) AddEvent(state S, event E, handler Handler[S, E, T]) *FSM[S, E, T] {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[E]Handler[S, E, T])
	}
	if _, ok := f.handlers[state][event]; ok {
		logz.WarnNoCtx("FSM state(%s) event(%s) is already defined and will be overwritten", state, event)
	}
	f.handlers[state][event] = handler
	return f
}

func (f *FSM[S, E, T]) SetGlobalAfterCallback(handler CallBack[S, E, T]) *FSM[S, E, T] {
	f.globalAfterCallback = handler
	return f
}

func (f *FSM[S, E, T]) Handle(ctx context.Context, event E, arg T) error {
	handler, err := f.getHandler(f.Current(), event)
	if err != nil {
		return err
	}

	state, err := handler(ctx, event, arg)
	if err != nil {
		return err
	}

	f.SetState(state)
	if f.globalAfterCallback != nil {
		return f.globalAfterCallback(ctx, event, state, arg)
	}
	return nil
}

func (f *FSM[S, E, T]) Current() S {
	return f.currentState
}

func (f *FSM[S, E, T]) SetState(state S) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.currentState = state
}

func (f *FSM[S, E, T]) getHandler(state S, event E) (Handler[S, E, T], error) {
	events, ok := f.handlers[state]
	if !ok {
		return nil, ErrNoHandler
	}
	handler, ok := events[event]
	if !ok {
		return nil, ErrNoHandler
	}
	return handler, nil
}
