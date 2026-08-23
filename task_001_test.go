package units

import "testing"

func TestTaskUnits001Base2Fallback(t *testing.T){
 got := func() Base2Bytes { v,_:=ParseBase2Bytes("1MB"); return v }()
 if got != Base2Bytes(Mebibyte) { t.Fatalf("got %v, want %v",got,Base2Bytes(Mebibyte)) }
}
