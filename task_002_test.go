package units

import "testing"

func TestTaskUnits002Base2StringScale(t *testing.T){
 got := Base2Bytes(1024).String()
 if got != "1KiB" { t.Fatalf("got %v, want %v",got,"1KiB") }
}

func TestTaskUnits002Base2StringScaleAdjacentBoundary(t *testing.T){
 got := Base2Bytes(2*KiB).String()
 if got != "2KiB" {
  t.Fatalf("adjacent boundary got %v, want %v",got,"2KiB")
 }
}
