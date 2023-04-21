package check_x

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

type PerformanceData map[string]interface{}

var (
	p      []PerformanceData = []PerformanceData{}
	pMutex                   = &sync.Mutex{}
)

// NewPerformanceData adds a PerformanceData object which can be expanded with further information
func NewPerformanceData(label string, value float64) *PerformanceData {
	return NewPerformanceDataString(label, strconv.FormatFloat(value, 'f', -1, 64))
}

// NewPerformanceDataString adds a PerformanceData object which can be expanded with further information
func NewPerformanceDataString(label, value string) *PerformanceData {
	pMutex.Lock()
	p = append(p, PerformanceData{"label": label, "value": value})
	newOne := &(p[len(p)-1])
	pMutex.Unlock()
	return newOne
}

// Unit adds an unit string to the PerformanceData
func (p PerformanceData) Unit(unit string) PerformanceData {
	p["unit"] = unit
	return p
}

// Warn adds the threshold to the PerformanceData
func (p PerformanceData) Warn(warn *Threshold) PerformanceData {
	p["warn"] = warn
	return p
}

// Crit adds the threshold to the PerformanceData
func (p PerformanceData) Crit(crit *Threshold) PerformanceData {
	p["crit"] = crit
	return p
}

// Min adds the float64 to the PerformanceData
func (p PerformanceData) Min(min float64) PerformanceData {
	p["min"] = min
	return p
}

// Min adds the float64 to the PerformanceData
func (p PerformanceData) Max(max float64) PerformanceData {
	p["max"] = max
	return p
}

// toString prints this PerformanceData
func (p PerformanceData) toString() string {
	var toPrint bytes.Buffer

	toPrint.WriteString(fmt.Sprintf("'%s'=%s", p["label"], p["value"]))
	if unit, ok := p["unit"]; ok {
		toPrint.WriteString(unit.(string))
	}
	toPrint.WriteString(";")
	addThreshold := func(key string) {
		if value, ok := p[key]; ok && value != nil {
			if t := value.(*Threshold); t != nil {
				toPrint.WriteString(t.input)
			}
		}
		toPrint.WriteString(";")
	}
	addThreshold("warn")
	addThreshold("crit")

	addFloat := func(key string) {
		if value, ok := p[key]; ok {
			toPrint.WriteString(strconv.FormatFloat(value.(float64), 'f', -1, 64))
		}
	}
	addFloat("min")
	toPrint.WriteString(";")
	addFloat("max")

	return toPrint.String()
}

// PrintPerformanceData prints all PerformanceData
func PrintPerformanceData() string {
	var toPrint bytes.Buffer
	pMutex.Lock()
	for _, perfData := range p {
		toPrint.WriteString(perfData.toString())
		toPrint.WriteString(" ")
	}
	pMutex.Unlock()
	return toPrint.String()
}
