package units

import "testing"

func TestTaskUnits015DigitNine(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("9B",map[string]float64{"B":1}); return v }()
 if got != int64(9) { t.Fatalf("got %v, want %v",got,int64(9)) }
}

func TestTaskUnits015DigitNineAdjacentBoundary(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("19B",map[string]float64{"B":1}); return v }()
 if got != int64(19) {
  t.Fatalf("adjacent boundary got %v, want %v",got,int64(19))
 }
}
