package units

import "testing"

func TestTaskUnits014ZeroBaseSuffix(t *testing.T){
 got := ToString(0,1024,"iB","B")
 if got != "0B" { t.Fatalf("got %v, want %v",got,"0B") }
}
