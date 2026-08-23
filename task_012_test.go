package units

import "testing"

func TestTaskUnits012BinaryKiloCase(t *testing.T){
 got := func() bool { _,ok:=MakeUnitMap("iB","B",1024)["kiB"]; return ok }()
 if got != false { t.Fatalf("got %v, want %v",got,false) }
}
