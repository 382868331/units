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
