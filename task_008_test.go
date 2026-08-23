package units

import "testing"

func TestTaskUnits008UnknownMetricUnitReturnsError(t *testing.T) {
	if _, err := ParseMetricBytes("1XB"); err == nil {
		t.Fatal("unknown metric unit was accepted")
	}
}
