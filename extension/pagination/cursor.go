package pagination

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"scan/extension/errorx"
)

type Cursor string

func NewCursor(value any) (Cursor, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return Cursor(base64.RawStdEncoding.EncodeToString(result)), nil
}

func NewCursorX(value any) Cursor {
	result, err := NewCursor(value)
	if err != nil {
		panic(err)
	}
	return result
}

func (c Cursor) Decode(value any) error {
	if c == "" {
		return nil
	}

	bytesValue, err := base64.RawStdEncoding.DecodeString(string(c))
	if err != nil {
		return errorx.ErrIllegalArgument.WithReason("specify cursor is invalid").Wrap(err)
	}

	if err := json.Unmarshal(bytesValue, value); err != nil {
		return errorx.ErrIllegalArgument.WithReason("specify cursor is invalid").Wrap(err)
	}

	return nil
}

func (c Cursor) String() string {
	return string(c)
}

func (c Cursor) IsEmpty() bool {
	return c == ""
}

type CursorPagination struct {
	// Before Cursor `json:"before" form:"before"`
	After Cursor `json:"after" form:"after"`
	Limit int    `json:"limit" form:"limit"`
}

func (p CursorPagination) SafeLimit() int {
	if p.Limit == 0 {
		return DEFAULT_PAGE_SIZE
	}
	if p.Limit > MAX_PAGE_SIZE {
		return MAX_PAGE_SIZE
	}
	return p.Limit
}

func (p CursorPagination) Validate() error {
	if p.Limit > MAX_PAGE_SIZE {
		return errorx.ErrIllegalArgument.WithReason(fmt.Sprintf("limit must be less than or equal to %d", MAX_PAGE_SIZE))
	}
	if p.Limit < 0 {
		return errorx.ErrIllegalArgument.WithReason("limit must be greater than or equal to 0")
	}

	return nil
}

type CursorPaginationResult struct {
	Self Cursor `json:"self"`
	Next Cursor `json:"next"`
	// Prev  Cursor `json:"prev"`
	Limit uint64 `json:"limit"`
}
