package units

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"unicode/utf8"
)

var (
	_ = context.Background
	_ = errors.Is
	_ = reflect.DeepEqual
	_ = utf8.ValidString
)

func TestUnitCompositeUnitOrder(t *testing.T) {
	if got := UnitCompositeUnitOrder(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitCompositeUnitOrderRegression(t *testing.T) {
	TestUnitCompositeUnitOrder(t)
	TestUnitCompositeUnitOrder(t)
}

func TestUnitBinaryMetricSuffix(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := UnitBinaryMetricSuffix(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestUnitBinaryMetricSuffixRegression(t *testing.T) {
	TestUnitBinaryMetricSuffix(t)
	TestUnitBinaryMetricSuffix(t)
}

func TestUnitFractionalMagnitude(t *testing.T) {
	got := UnitFractionalMagnitude("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}
