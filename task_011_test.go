package units

import "testing"

func TestTaskUnits011MegaMultiplier(t *testing.T){
 got := int64(MakeUnitMap("B","B",1000)["MB"])
 if got != int64(1000000) { t.Fatalf("got %v, want %v",got,int64(1000000)) }
}

func TestTaskUnits011MegaMultiplierAdjacentBoundary(t *testing.T){
 got := int64(MakeUnitMap("B","B",1000)["GB"])
 if got != int64(1000000000) {
  t.Fatalf("adjacent boundary got %v, want %v",got,int64(1000000000))
 }
}
