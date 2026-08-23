package units

import "testing"

func TestTaskUnits004Base2FloorUnit(t *testing.T){
 got := Base2Bytes(GiB+2*MiB).Floor()
 if got != Base2Bytes(GiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(GiB)) }
}
