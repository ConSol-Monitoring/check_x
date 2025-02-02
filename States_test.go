package check_x

import (
	"reflect"
	"testing"
)

var stateToString = []struct {
	input  State
	result string
}{
	{OK, "OK"},
	{Warning, "WARNING"},
	{Critical, "CRITICAL"},
	{Unknown, "UNKNOWN"},
}

func TestState_String(t *testing.T) {
	for i, data := range stateToString {
		if data.input.String() != data.result {
			t.Errorf("%d - Got: %s - expected: %s", i, data.input.String(), data.result)
		}
	}
}

var intToState = []struct {
	input  int
	result State
}{
	{0, OK},
	{1, Warning},
	{2, Critical},
	{3, Unknown},
	{4, State{code: 4}},
}

func TestStateFromInt(t *testing.T) {
	for i, data := range intToState {
		if StateFromInt(data.input) != data.result {
			t.Errorf("%d - Got: %d - expected: %s", i, data.input, data.result)
		}
	}
}

var stringToState = []struct {
	input  string
	result State
}{
	{"OK", OK},
	{"Warning", Warning},
	{"Critical", Critical},
	{"Unknown", Unknown},
	{"FOO", State{name: "FOO"}},
}

func TestStateFromString(t *testing.T) {
	for i, data := range stringToState {
		if StateFromString(data.input) != data.result {
			t.Errorf("%d - Got: %s - expected: %s", i, data.input, data.result)
		}
	}
}

var worstState = []struct {
	input  States
	result *State
	err    error
}{
	{States{}, nil, ErrEmptyStates},
	{States{OK}, &OK, nil},
	{States{OK, Warning}, &Warning, nil},
	{States{Critical, OK, Warning}, &Critical, nil},
	{States{OK, Unknown, Warning, Critical}, &Unknown, nil},
}

func TestStates_GetWorst(t *testing.T) {
	for i, data := range worstState {
		if worst, err := data.input.GetWorst(); err != data.err {
			t.Errorf("%d - Got: %s - expected: %s", i, err, data.err)
		} else if (worst == nil) != (data.result == nil) {
			t.Errorf("%d - Got: %s - expected: %s", i, worst, data.result)
		} else if worst != nil && !reflect.DeepEqual(*worst, *data.result) {
			t.Errorf("%d - Got: %s - expected: %s", i, worst, data.result)
		}
	}
}

var bestState = []struct {
	input  States
	result *State
	err    error
}{
	{States{}, nil, ErrEmptyStates},
	{States{OK}, &OK, nil},
	{States{OK, Warning}, &OK, nil},
	{States{OK, Critical, Warning}, &OK, nil},
	{States{OK, Warning, Unknown, Critical}, &OK, nil},
}

func TestStates_GetBest(t *testing.T) {
	for i, data := range bestState {
		if worst, err := data.input.GetBest(); err != data.err {
			t.Errorf("%d - Got: %s - expected: %s", i, err, data.err)
		} else if (worst == nil) != (data.result == nil) {
			t.Errorf("%d - Got: %s - expected: %s", i, worst, data.result)
		} else if worst != nil && !reflect.DeepEqual(*worst, *data.result) {
			t.Errorf("%d - Got: %s - expected: %s", i, worst, data.result)
		}
	}
}
