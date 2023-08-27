package answers

import (
	"reflect"
	"testing"
)


var answers Answers = Answers{
	Yes: AnswerInfo{
		Number: 1, 
		Description: "Yes",
		}, 
	No: AnswerInfo{
		Number: 2, 
		Description: "No",
		}, 
}


func TestAnswersToSlice(test *testing.T) {
	answers_slice := answers.ToSlice()
	answers_slice_len := len(answers_slice)
	wanted_len := 2

	if wanted_len != answers_slice_len {
		test.Fatalf("Operations_slice length is incorrect: %v wanted, but %v received", wanted_len, answers_slice_len)
	}

	if reflect.TypeOf(answers_slice) != reflect.TypeOf([]AnswerInfo{}) {
		test.Fatalf("%v is not a slice of type OperationInfo", answers_slice)
	}
}
