package check_x

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

// PerformanceData is a map with string keys and any values, essentially a dictionary
// Since its an dictionary, always add add a label
// "label" = "health_rate"
// "value" = 80
// "critical" = 160
type PerformanceData map[string]any

// This is a struct that users of the library should use.
// They have to manage this object through functions in this file
// Library clients should get an instance of this and keep calling its methods
type PerformanceDataCollection struct {
	data      []PerformanceData
	dataMutex *sync.Mutex
}

// Returns an empty PerformanceDataCollection
func NewPerformanceDataCollection() PerformanceDataCollection {
	return PerformanceDataCollection{
		data:      make([]PerformanceData, 0),
		dataMutex: &sync.Mutex{},
	}
}

// Adds a new PerformanceData element to the data array
// Use the label to get the reference to the PerformanceData later on
func (collection *PerformanceDataCollection) AddPerformanceData(label string, value string) {
	collection.dataMutex.Lock()
	collection.data = append(collection.data, PerformanceData{"label": label, "value": value})
	collection.dataMutex.Unlock()
}

// Calls collection.AddPerformanceData after converting float to string
func (collection *PerformanceDataCollection) AddPerformanceDataFloat64(label string, value float64) {
	collection.AddPerformanceData(label, strconv.FormatFloat(value, 'f', -1, 64))
}

// Internal function to find a PerformanceData with specified label
func (collection *PerformanceDataCollection) findPerformanceData(label string) (*PerformanceData, error) {
	for index, pd := range collection.data {
		if pd_label, ok := pd["label"]; ok && pd_label == label {
			return &collection.data[index], nil
		}
	}
	return nil, fmt.Errorf("no performance data with the label '%s' found", label)
}

// Adds a field called "unit" to the PerformanceData with the label
func (collection *PerformanceDataCollection) Unit(label string, unit string) error {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return err
	}
	(*pd)["unit"] = unit
	return nil
}

// Adds a field called "unit" to the PerformanceData with the label
func (collection *PerformanceDataCollection) Warn(label string, warn *Threshold) error {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return err
	}
	(*pd)["warn"] = warn
	return nil
}

// Adds a field called "unit" to the PerformanceData with the label
func (collection *PerformanceDataCollection) Crit(label string, crit *Threshold) error {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return err
	}
	(*pd)["crit"] = crit
	return nil
}

// Adds a field called "unit" to the PerformanceData with the label
func (collection *PerformanceDataCollection) Min(label string, min float64) error {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return err
	}
	(*pd)["min"] = min
	return nil
}

// Adds a field called "unit" to the PerformanceData with the label
func (collection *PerformanceDataCollection) Max(label string, max float64) error {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return err
	}
	(*pd)["max"] = max
	return nil
}

// internal function to print a PerformanceData
func (pd PerformanceData) toString() string {
	var toPrint bytes.Buffer

	toPrint.WriteString(fmt.Sprintf("'%s'=%s", pd["label"], pd["value"]))
	if unit, ok := pd["unit"]; ok {
		toPrint.WriteString(unit.(string))
	}
	toPrint.WriteString(";")
	addThreshold := func(key string) {
		if value, ok := pd[key]; ok && value != nil {
			if t := value.(*Threshold); t != nil {
				toPrint.WriteString(t.input)
			}
		}
		toPrint.WriteString(";")
	}
	addThreshold("warn")
	addThreshold("crit")

	addFloat := func(key string) {
		if value, ok := pd[key]; ok {
			toPrint.WriteString(strconv.FormatFloat(value.(float64), 'f', -1, 64))
		}
	}
	addFloat("min")
	toPrint.WriteString(";")
	addFloat("max")

	return toPrint.String()
}

// Finds and prints the PerformanceData found in this collection
func (collection *PerformanceDataCollection) PrintPerformanceData(label string) (string, error) {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	pd, err := collection.findPerformanceData(label)
	if err != nil {
		return "", err
	}

	return (*pd).toString(), nil
}

// Prints all PerformanceData to a string in this collection
func (collection *PerformanceDataCollection) PrintAllPerformanceData() string {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	var toPrint bytes.Buffer
	for _, perfData := range collection.data {
		toPrint.WriteString(perfData.toString())
		toPrint.WriteString(" ")
	}
	return toPrint.String()
}

// Clears all PerformanceData stored in this collection
func (collection *PerformanceDataCollection) ClearPerformanceCollection() {
	collection.dataMutex.Lock()
	defer collection.dataMutex.Unlock()
	collection.data = make([]PerformanceData, 0)
}
