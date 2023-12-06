package pagination

const (
	DEFAULT_PAGE_SIZE int = 10
	MAX_PAGE_SIZE     int = 1000
)

type LimitOffsetPagination struct {
	PrimitivePage     uint64 `json:"page" form:"page"`           // 页码
	PrimitivePageSize uint64 `json:"page_size" form:"page_size"` // 每页数量
}

func (p LimitOffsetPagination) Offset() uint64 {
	if p.PrimitivePage == 0 || p.PrimitivePage == 1 {
		return 0
	}

	return (p.PrimitivePage - 1) * p.PrimitivePageSize
}

func (p LimitOffsetPagination) PageSize() uint64 {
	switch {
	case p.PrimitivePageSize == 0:
		return uint64(DEFAULT_PAGE_SIZE)
	case p.PrimitivePageSize > 100:
		return uint64(MAX_PAGE_SIZE)
	default:
		return p.PrimitivePageSize
	}
}

func (p LimitOffsetPagination) Page() uint64 {
	if p.PrimitivePage == 0 || p.PrimitivePage == 1 {
		return 1
	}
	return p.PrimitivePage
}

func (p LimitOffsetPagination) NextPage(total uint64) uint64 {
	if p.Page() >= p.TotalPage(total) {
		return p.Page()
	}
	return p.Page() + 1
}

func (p LimitOffsetPagination) TotalPage(total uint64) uint64 {
	if total == 0 {
		return 1
	}

	if total%p.PageSize() == 0 {
		return total / p.PageSize()
	}

	return total/p.PageSize() + 1
}
