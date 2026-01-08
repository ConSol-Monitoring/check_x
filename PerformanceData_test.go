package check_x

import (
	"testing"
)

var perfdataToString = []struct {
	f        func() PerformanceDataCollection
	expected string
}{
	{func() PerformanceDataCollection {
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "1")
		return col
	}, "'a'=1;;;;"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "2")
		col.Warn("a", warn)
		return col
	}, "'a'=2;10:;;;"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "3")
		col.Warn("a", warn)
		col.Crit("a", crit)
		return col
	}, "'a'=3;10:;@10:20;;"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "3")
		col.Warn("a", warn)
		col.Crit("a", crit)
		col.Min("a", 0)
		return col
	}, "'a'=3;10:;@10:20;0;"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "4")
		col.Warn("a", warn)
		col.Crit("a", crit)
		col.Min("a", 0)
		col.Max("a", 100)
		return col
	}, "'a'=4;10:;@10:20;0;100"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceData("a", "5")
		col.Warn("a", warn)
		col.Crit("a", crit)
		col.Min("a", 0)
		col.Max("a", 100)
		col.Unit("a", "C")
		return col
	}, "'a'=5C;10:;@10:20;0;100"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("10:")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceDataFloat64("a", 6)
		col.Warn("a", warn)
		col.Crit("a", crit)
		col.Min("a", 0)
		col.Max("a", 100)
		col.Unit("a", "C")
		return col
	}, "'a'=6C;10:;@10:20;0;100"},
	{func() PerformanceDataCollection {
		warn, _ := NewThreshold("")
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceDataFloat64("a", 6)
		col.Warn("a", warn)
		col.Crit("a", crit)
		col.Min("a", 0)
		col.Max("a", 100)
		col.Unit("a", "C")
		return col
	}, "'a'=6C;;@10:20;0;100"},
	{func() PerformanceDataCollection {
		crit, _ := NewThreshold("@10:20")
		col := NewPerformanceDataCollection()
		col.AddPerformanceDataFloat64("a", 6)
		col.Warn("a", nil)
		col.Crit("a", crit)
		col.Min("a", 0)
		col.Max("a", 100)
		col.Unit("a", "C")
		return col
	}, "'a'=6C;;@10:20;0;100"},
}

func TestPerformanceData_toString(t *testing.T) {
	for i, data := range perfdataToString {
		collection := data.f()
		collectionString, err := collection.PrintPerformanceData("a")
		if err != nil {
			t.Errorf("Error when finding the performance data: %s", err.Error())
		}
		if collectionString != data.expected {
			t.Errorf("%d - Expected: %s, got: %s", i, data.expected, collectionString)
		}
	}
}

func TestPrintPerformanceData(t *testing.T) {
	col := NewPerformanceDataCollection()
	col.AddPerformanceData("a", "1")
	col.AddPerformanceData("b", "2")
	expected := "'a'=1;;;;" + " " + "'b'=2;;;; "
	if expected != col.PrintAllPerformanceData() {
		t.Errorf("Expected: %s, got: %s", expected, col.PrintAllPerformanceData())
	}
}
