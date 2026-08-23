package units

import "testing"

func TestTaskUnits017NegativeSign(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("-2B",map[string]float64{"B":1}); return v }()
 if got != int64(-2) { t.Fatalf("got %v, want %v",got,int64(-2)) }
}
