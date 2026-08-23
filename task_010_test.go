package units

import "testing"

func TestTaskUnits010StrictFallback(t *testing.T){
 got := func() int64 { v,_:=ParseStrictBytes("1kB"); return v }()
 if got != int64(1000) { t.Fatalf("got %v, want %v",got,int64(1000)) }
}
