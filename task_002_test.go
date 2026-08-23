package units

import "testing"

func TestTaskUnits002Base2StringScale(t *testing.T){
 got := Base2Bytes(1024).String()
 if got != "1KiB" { t.Fatalf("got %v, want %v",got,"1KiB") }
}
