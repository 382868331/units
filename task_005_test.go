package units

import "testing"

func TestTaskUnits005Base2RoundIndex(t *testing.T){
 got := Base2Bytes(GiB+MiB).Round(1)
 if got != Base2Bytes(GiB) { t.Fatalf("got %v, want %v",got,Base2Bytes(GiB)) }
}
