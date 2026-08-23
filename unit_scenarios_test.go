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

func TestUnitFractionalMagnitudeRegression(t *testing.T) {
	TestUnitFractionalMagnitude(t)
	TestUnitFractionalMagnitude(t)
}

func TestUnitStableAliasOrder(t *testing.T) {
	got := UnitStableAliasOrder([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestUnitStableAliasOrderRegression(t *testing.T) {
	TestUnitStableAliasOrder(t)
	TestUnitStableAliasOrder(t)
}

func TestUnitEmptySegmentProgress(t *testing.T) {
	if got := UnitEmptySegmentProgress([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestUnitEmptySegmentProgressRegression(t *testing.T) {
	TestUnitEmptySegmentProgress(t)
	TestUnitEmptySegmentProgress(t)
}

func TestUnitUnicodeMicroPrefix(t *testing.T) {
	got := UnitUnicodeMicroPrefix("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestUnitUnicodeMicroPrefixRegression(t *testing.T) {
	TestUnitUnicodeMicroPrefix(t)
	TestUnitUnicodeMicroPrefix(t)
}

func TestUnitSignedZeroParsing(t *testing.T) {
	got, err := UnitSignedZeroParsing(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestUnitSignedZeroParsingRegression(t *testing.T) {
	TestUnitSignedZeroParsing(t)
	TestUnitSignedZeroParsing(t)
}

func TestUnitPrefixExponentOverflow(t *testing.T) {
	if got := UnitPrefixExponentOverflow(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitPrefixExponentOverflowRegression(t *testing.T) {
	TestUnitPrefixExponentOverflow(t)
	TestUnitPrefixExponentOverflow(t)
}

func TestUnitExactFloorBoundary(t *testing.T) {
	if got := UnitExactFloorBoundary([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitExactFloorBoundaryRegression(t *testing.T) {
	TestUnitExactFloorBoundary(t)
	TestUnitExactFloorBoundary(t)
}

func TestUnitRegistryCopyIsolation(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := UnitRegistryCopyIsolation(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestUnitRegistryCopyIsolationRegression(t *testing.T) {
	TestUnitRegistryCopyIsolation(t)
	TestUnitRegistryCopyIsolation(t)
}

func TestUnitNegativeFormatting(t *testing.T) {
	if got := UnitNegativeFormatting("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestUnitNegativeFormattingRegression(t *testing.T) {
	TestUnitNegativeFormatting(t)
	TestUnitNegativeFormatting(t)
}

func TestUnitRoundLastComponent(t *testing.T) {
	got := UnitRoundLastComponent([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestUnitRoundLastComponentRegression(t *testing.T) {
	TestUnitRoundLastComponent(t)
	TestUnitRoundLastComponent(t)
}

func TestUnitZeroValueSuffix(t *testing.T) {
	if got := UnitZeroValueSuffix(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestUnitZeroValueSuffixRegression(t *testing.T) {
	TestUnitZeroValueSuffix(t)
	TestUnitZeroValueSuffix(t)
}

func TestUnitConcurrentRegistryRead(t *testing.T) {
	if got := UnitConcurrentRegistryRead(64); got != 64 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitConcurrentRegistryReadRegression(t *testing.T) {
	TestUnitConcurrentRegistryRead(t)
	TestUnitConcurrentRegistryRead(t)
}

func TestUnitParseCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := UnitParseCancellation(ctx, 20); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitParseCancellationRegression(t *testing.T) {
	TestUnitParseCancellation(t)
	TestUnitParseCancellation(t)
}

func TestUnitWrappedParseCause(t *testing.T) {
	base := errors.New("root")
	if got := UnitWrappedParseCause(base); !errors.Is(got, base) {
		t.Fatalf("chain lost: %v", got)
	}
}

func TestUnitWrappedParseCauseRegression(t *testing.T) {
	TestUnitWrappedParseCause(t)
	TestUnitWrappedParseCause(t)
}

func TestUnitParserScratchCleanup(t *testing.T) {
	active = 0
	UnitParserScratchCleanup(true)
	if active != 0 {
		t.Fatalf("active=%d", active)
	}
}
