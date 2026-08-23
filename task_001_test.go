package units

import "testing"

func TestTaskUnits001Base2Fallback(t *testing.T){
 got := func() Base2Bytes { v,_:=ParseBase2Bytes("1MB"); return v }()
 if got != Base2Bytes(Mebibyte) { t.Fatalf("got %v, want %v",got,Base2Bytes(Mebibyte)) }
}

func TestTaskUnits001Base2FallbackAdjacentBoundary(t *testing.T){
 got := func() Base2Bytes { v,_:=ParseBase2Bytes("2GB"); return v }()
 if got != Base2Bytes(2*Gibibyte) {
  t.Fatalf("adjacent boundary got %v, want %v",got,Base2Bytes(2*Gibibyte))
 }
}
