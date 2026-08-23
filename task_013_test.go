package units

import "testing"

func TestTaskUnits013StringSlotOrder(t *testing.T){
 got := ToString(1025,1024,"iB","B")
 if got != "1KiB1B" { t.Fatalf("got %v, want %v",got,"1KiB1B") }
}

func TestTaskUnits013StringSlotOrderAdjacentBoundary(t *testing.T){
 got := ToString(int64(MiB+KiB),1024,"iB","B")
 if got != "1MiB1KiB" {
  t.Fatalf("adjacent boundary got %v, want %v",got,"1MiB1KiB")
 }
}
