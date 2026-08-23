package units

import (
	"encoding/json"
	"testing"
)

func TestTaskUnits007JSONUsesBinaryUnits(t *testing.T) {
	got, err := json.Marshal(Base2Bytes(KiB))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `"1KiB"` {
		t.Fatalf("got %s, want %s", got, `"1KiB"`)
	}
}

func TestTaskUnits007JSONKeepsLargerBinaryUnit(t *testing.T) {
	got, err := json.Marshal(Base2Bytes(2 * MiB))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `"2MiB"` {
		t.Fatalf("got %s, want %s", got, `"2MiB"`)
	}
}
