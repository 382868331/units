package units

import "testing"

func TestTaskUnits004Base2FloorUnit(t *testing.T){
 got := Base2Bytes(GiB+2*MiB).Floor()
 if got != Base2Bytes(GiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(GiB)) }
}

func TestTaskUnits004Base2FloorUnitAdjacentBoundary(t *testing.T){
 got := Base2Bytes(2*GiB+MiB).Floor()
 if got != Base2Bytes(2*GiB) {
  t.Fatalf("adjacent boundary got %v, want %v",got,Base2Bytes(2*GiB))
 }
}
