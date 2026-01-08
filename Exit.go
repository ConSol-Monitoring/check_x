package check_x

import (
	"fmt"
	"os"
)

// ExitOnError quits with unknown and the error message if an error was passed
func ExitOnError(err error) {
	if err != nil {
		Exit(Unknown, err.Error())
	}
}

// ErrorExit quits with unknown and the error message
func ErrorExit(err error) {
	Exit(Unknown, err.Error())
}

// Exit returns with the given returncode and message and optional PerformanceData
func Exit(state State, msg string) {
	LongExit(state, msg, "", nil)
}

// LongExit returns with the given returncode and message and optional PerformanceData and long message
func LongExit(state State, msg, longMsg string, collection *PerformanceDataCollection) {
	perfString := ""
	if collection != nil {
		perfString = collection.PrintAllPerformanceData()
	}

	if perfString == "" {
		fmt.Printf("%s - %s\n%s", state.Name, msg, longMsg)
	} else {
		fmt.Printf("%s - %s|%s\n%s", state.Name, msg, perfString, longMsg)
	}
	os.Exit(state.Code)
}
