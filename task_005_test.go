package units

import "testing"

func TestTaskUnits005Base2RoundIndex(t *testing.T){
 got := Base2Bytes(GiB+MiB).Round(1)
 if got != Base2Bytes(GiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(GiB)) }
}

func TestTaskUnits005Base2RoundIndexAdjacentBoundary(t *testing.T){
 got := Base2Bytes(2*GiB+MiB).Round(1)
 if got != Base2Bytes(2*GiB) {
  t.Fatalf("adjacent boundary got %v, want %v",got,Base2Bytes(2*GiB))
 }
}
