package errorx

import (
	"fmt"
	"testing"

	stderr "github.com/pkg/errors"
)

type poser struct {
	msg string
}

func (p poser) Error() string {
	return p.msg
}

func TestErrorAs(t *testing.T) {
	testcases := []struct {
		err    error
		target any
		want   bool
	}{
		{
			err:    NewError(404, "not found"),
			target: &Error{},
			want:   true,
		},
		{
			err:    NewError(404, "not found").Wrap(stderr.New("original")),
			target: &Error{},
			want:   true,
		},
		{
			err:    poser{msg: "poser"},
			target: &Error{},
			want:   false,
		},
		{
			err:    &Error{},
			target: &poser{},
			want:   false,
		},
		{
			err:    nil,
			target: &Error{},
			want:   false,
		},
		{
			err:    stderr.New("original"),
			target: &Error{},
			want:   false,
		},
	}

	for i, tt := range testcases {
		name := fmt.Sprintf("%d:As(Errorf(..., %v), %v)", i, tt.err, tt.target)
		t.Run(name, func(t *testing.T) {
			result := stderr.As(tt.err, tt.target)
			if result != tt.want {
				t.Errorf("stderr.As(%v, %T(%v)) = %v, want %v", tt.err, tt.target, tt.target, result, tt.want)
			}
		})
	}
}

func TestErrorAsValueAssign(t *testing.T) {
	var target Error
	err := &Error{
		Code:     404,
		Reason:   "not found",
		Message:  "user not fount",
		GRPCCode: 404,
		err:      nil,
	}
	wrappedErr := err.Wrap(stderr.New("original"))
	if !stderr.As(wrappedErr, &target) {
		t.Errorf("stderr.As(%v, %T(%v)) = %v, want true", err, target, target, false)
	}

	if target.Code != err.Code {
		t.Errorf("target.Code = %v, want %d", target.Code, err.Code)
	}
	if target.Reason != err.Reason {
		t.Errorf("target.Message = %v, want '%v'", target.Message, err.Reason)
	}
	if target.Message != err.Message {
		t.Errorf("target.Message = %v, want '%v'", target.Message, err.Message)
	}
	if target.GRPCCode != err.GRPCCode {
		t.Errorf("target.GRPCCode = %v, want %d", target.GRPCCode, err.GRPCCode)
	}
}

func TestErrorLevel(t *testing.T) {
	testcases := []struct {
		name string
		err  error
		want Level
	}{
		{
			name: "should get error level",
			err:  NewErrorWithLevel(404, "not found", LevelError),
			want: LevelError,
		},
		{
			name: "should get warning level",
			err:  NewErrorWithLevel(404, "not found", LevelWarning),
			want: LevelWarning,
		},
		{
			name: "should get info level",
			err:  NewErrorWithLevel(404, "not found", LevelInfo),
			want: LevelInfo,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if e, ok := tc.err.(LevelContainer); ok {
				if e.Level() != tc.want {
					t.Errorf("LevelOf(%v) = %v, want %v", tc.err, e.Level(), tc.want)
				}
			} else {
				t.Errorf("LevelOf(%v) = %v, want %v", tc.err, nil, tc.want)
			}
		})
	}
}
