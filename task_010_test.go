package units

import "testing"

func TestTaskUnits010StrictFallback(t *testing.T){
 got := func() int64 { v,_:=ParseStrictBytes("1kB"); return v }()
 if got != int64(1000) { t.Fatalf("got %v, want %v",got,int64(1000)) }
}

func TestTaskUnits010StrictFallbackAdjacentBoundary(t *testing.T){
 got := func() int64 { v,_:=ParseStrictBytes("2MB"); return v }()
 if got != int64(2000000) {
  t.Fatalf("adjacent boundary got %v, want %v",got,int64(2000000))
 }
}
