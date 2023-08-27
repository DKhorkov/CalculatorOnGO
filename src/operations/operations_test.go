package operations

import (
	"reflect"
	"testing"
)


var operations = Operations{
	Summarizing: OperationInfo{
		Number: 1,
		Description: "Summarize first number with second",
		ResultNotification: "Result of summarizing %v and %v is %v.\n",
	},
	Subtracting: OperationInfo{
		Number: 2,
		Description: "Subtracte second number from first",
		ResultNotification: "Result of substracting from %v number %v is %v.\n",
	},
}

func TestOperationsToSlice(test *testing.T) {
	operations_slice := operations.ToSlice()
	operations_slice_len := len(operations_slice)
	expected_len := 6

	if expected_len != operations_slice_len {
		test.Fatalf("Operations_slice length is incorrect: %v expected, but %v received", expected_len, operations_slice_len)
	}

	if reflect.TypeOf(operations_slice) != reflect.TypeOf([]OperationInfo{}) {
		test.Fatalf("%v is not a slice of type OperationInfo", operations_slice)
	}
}
