package datetime

import "testing"

func TestNow(t *testing.T) {
	t.Parallel()
	now := Now()

	if now.Location().String() != "Asia/Shanghai" {
		t.Errorf("Now() = %v, want %v", now.Location().String(), "Asia/Shanghai")
	}
}
