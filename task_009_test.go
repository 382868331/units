package units

import "testing"

func TestTaskUnits009CompoundMetricSegmentsAccumulate(t *testing.T) {
	got, err := ParseMetricBytes("1MB500KB")
	if err != nil {
		t.Fatal(err)
	}
	if got != MetricBytes(1500000) {
		t.Fatalf("got %d, want 1500000", got)
	}
}

func TestTaskUnits009ThreeMetricSegmentsAccumulate(t *testing.T) {
	got, err := ParseMetricBytes("2GB250MB3KB")
	if err != nil {
		t.Fatal(err)
	}
	if got != MetricBytes(2250003000) {
		t.Fatalf("got %d, want 2250003000", got)
	}
}
