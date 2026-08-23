package units

import "testing"

func TestTaskUnits020FractionScale(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("1.5KB",map[string]float64{"KB":1000}); return v }()
 if got != int64(1500) { t.Fatalf("got %v, want %v",got,int64(1500)) }
}

func TestTaskUnits020FractionScaleAdjacentBoundary(t *testing.T){
 got := func() int64 { v,_:=ParseUnit("2.25KB",map[string]float64{"KB":1000}); return v }()
 if got != int64(2250) {
  t.Fatalf("adjacent boundary got %v, want %v",got,int64(2250))
 }
}
