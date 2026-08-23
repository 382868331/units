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
