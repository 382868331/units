package units

import "testing"

func TestTaskUnits013StringSlotOrder(t *testing.T){
 got := ToString(1025,1024,"iB","B")
 if got != "1KiB1B" { t.Fatalf("got %v, want %v",got,"1KiB1B") }
}
