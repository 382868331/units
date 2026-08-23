package units

import "testing"

func TestTaskUnits011MegaMultiplier(t *testing.T){
 got := int64(MakeUnitMap("B","B",1000)["MB"])
 if got != int64(1000000) { t.Fatalf("got %v, want %v",got,int64(1000000)) }
}
