package units

import "testing"

func TestTaskUnits006MetricParseMap(t *testing.T){
 got := func() MetricBytes { v,_:=ParseMetricBytes("1KB"); return v }()
 if got != MetricBytes(1000) { t.Fatalf("got %v, want %v",got,MetricBytes(1000)) }
}
