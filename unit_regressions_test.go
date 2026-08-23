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

func TestUnitNormalizeBounds(t *testing.T) {
	if got := UnitNormalizeBounds(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitNormalizeBoundsRegression(t *testing.T) {
	TestUnitNormalizeBounds(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitNormalizeBounds(t)
}

func TestUnitSaturatingAdd(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := UnitSaturatingAdd(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestUnitSaturatingAddRegression(t *testing.T) {
	TestUnitSaturatingAdd(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitSaturatingAdd(t)
}

func TestUnitSplitEscapedTokens(t *testing.T) {
	got := UnitSplitEscapedTokens("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestUnitSplitEscapedTokensRegression(t *testing.T) {
	TestUnitSplitEscapedTokens(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitSplitEscapedTokens(t)
}

func TestUnitStableUnique(t *testing.T) {
	got := UnitStableUnique([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestUnitStableUniqueRegression(t *testing.T) {
	TestUnitStableUnique(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitStableUnique(t)
}

func TestUnitPartitionValues(t *testing.T) {
	if got := UnitPartitionValues([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestUnitPartitionValuesRegression(t *testing.T) {
	TestUnitPartitionValues(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitPartitionValues(t)
}

func TestUnitTruncateLabel(t *testing.T) {
	got := UnitTruncateLabel("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestUnitTruncateLabelRegression(t *testing.T) {
	TestUnitTruncateLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitTruncateLabel(t)
}

func TestUnitParseBooleanOption(t *testing.T) {
	got, err := UnitParseBooleanOption(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestUnitParseBooleanOptionRegression(t *testing.T) {
	TestUnitParseBooleanOption(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitParseBooleanOption(t)
}

func TestUnitBoundedBackoff(t *testing.T) {
	if got := UnitBoundedBackoff(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitBoundedBackoffRegression(t *testing.T) {
	TestUnitBoundedBackoff(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitBoundedBackoff(t)
}

func TestUnitSelectUpperQuantile(t *testing.T) {
	if got := UnitSelectUpperQuantile([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitSelectUpperQuantileRegression(t *testing.T) {
	TestUnitSelectUpperQuantile(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitSelectUpperQuantile(t)
}

func TestUnitCloneNestedState(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := UnitCloneNestedState(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestUnitCloneNestedStateRegression(t *testing.T) {
	TestUnitCloneNestedState(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitCloneNestedState(t)
}

func TestUnitReverseUnicodeLabel(t *testing.T) {
	if got := UnitReverseUnicodeLabel("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestUnitReverseUnicodeLabelRegression(t *testing.T) {
	TestUnitReverseUnicodeLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitReverseUnicodeLabel(t)
}

func TestUnitSlidingWindows(t *testing.T) {
	got := UnitSlidingWindows([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestUnitSlidingWindowsRegression(t *testing.T) {
	TestUnitSlidingWindows(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitSlidingWindows(t)
}

func TestUnitJoinOptionalParts(t *testing.T) {
	if got := UnitJoinOptionalParts(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestUnitJoinOptionalPartsRegression(t *testing.T) {
	TestUnitJoinOptionalParts(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitJoinOptionalParts(t)
}

func TestUnitCountConcurrentUpdates(t *testing.T) {
	if got := UnitCountConcurrentUpdates(64); got != 64 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitCountConcurrentUpdatesRegression(t *testing.T) {
	TestUnitCountConcurrentUpdates(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitCountConcurrentUpdates(t)
}

func TestUnitProcessUntilCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := UnitProcessUntilCanceled(ctx, 20); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestUnitProcessUntilCanceledRegression(t *testing.T) {
	TestUnitProcessUntilCanceled(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitProcessUntilCanceled(t)
}

func TestUnitWrapCause(t *testing.T) {
	base := errors.New("root")
	if got := UnitWrapCause(base); !errors.Is(got, base) {
		t.Fatalf("chain lost: %v", got)
	}
}

func TestUnitWrapCauseRegression(t *testing.T) {
	TestUnitWrapCause(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestUnitWrapCause(t)
}
