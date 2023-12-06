package contextz

import (
	"context"
)

func RequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	return ""
	// span := tracing.SpanFromContext(ctx)
	// if span.Empty() {
	// 	return ""
	// }

	// return span.RequestID()
}
