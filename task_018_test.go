package units

import "testing"

func TestTaskUnits018ZeroSpecialCase(t *testing.T){
 got := func() bool { v,e:=ParseUnit("0",map[string]float64{}); return e==nil && v==0 }()
 if got != true { t.Fatalf("got %v, want %v",got,true) }
}

func TestTaskUnits018ZeroSpecialCaseAdjacentBoundary(t *testing.T){
 got := func() bool { v,e:=ParseUnit("-0",map[string]float64{}); return e==nil && v==0 }()
 if got != true {
  t.Fatalf("adjacent boundary got %v, want %v",got,true)
 }
}
