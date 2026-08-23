package units

import "testing"

func TestTaskUnits019InitialDigitGuard(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("3B",map[string]float64{"B":1}); return v }()
 if got != int64(3) { t.Fatalf("got %v, want %v",got,int64(3)) }
}
