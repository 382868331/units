package units

import "testing"

func TestTaskUnits003UnmarshalAssignment(t *testing.T){
 got := func() Base2Bytes { var v Base2Bytes; _ = v.UnmarshalText([]byte("2KiB")); return v }()
 if got != Base2Bytes(2*KiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(2*KiB)) }
}
