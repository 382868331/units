package units

import "testing"

func TestTaskUnits003UnmarshalAssignment(t *testing.T){
 got := func() Base2Bytes { var v Base2Bytes; _ = v.UnmarshalText([]byte("2KiB")); return v }()
 if got != Base2Bytes(2*KiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(2*KiB)) }
}

func TestTaskUnits003UnmarshalAssignmentAdjacentBoundary(t *testing.T){
 got := func() Base2Bytes { var v Base2Bytes; _ = v.UnmarshalText([]byte("3MiB")); return v }()
 if got != Base2Bytes(3*MiB) {
  t.Fatalf("adjacent boundary got %v, want %v",got,Base2Bytes(3*MiB))
 }
}
