package units

import "testing"

func TestTaskUnits008UnknownMetricUnitReturnsError(t *testing.T) {
	if _, err := ParseMetricBytes("1XB"); err == nil {
		t.Fatal("unknown metric unit was accepted")
	}
}

func TestTaskUnits008MalformedMetricValueReturnsError(t *testing.T) {
	if _, err := ParseMetricBytes(".KB"); err == nil {
		t.Fatal("metric value without digits was accepted")
	}
}
