package calculator

import (
	"answers"
	"math"
	"operations"
	"reflect"
	"testing"
)


type testingDataValue struct {
	FirstNumber, SecondNumber, ExpectedResult TypeOfNumber
	OperationNumber int
}


var possible_answers answers.Answers = answers.Answers{
	Yes: answers.AnswerInfo{
		Number: 1, 
		Description: "Yes",
		}, 
	No: answers.AnswerInfo{
		Number: 2, 
		Description: "No",
		}, 
}

var possible_operations operations.Operations = operations.Operations{
	Summarizing: operations.OperationInfo{
		Number: 1,
		Description: "Summarize first number with second",
		ResultNotification: "Result of summarizing %v and %v is %v.\n",
	},
	Subtracting: operations.OperationInfo{
		Number: 2,
		Description: "Subtracte second number from first",
		ResultNotification: "Result of substracting from %v number %v is %v.\n",
	},
	Multipling: operations.OperationInfo{
		Number: 3,
		Description: "Multiply first number on second",
		ResultNotification: "Result of multiplying %v on %v is %v.\n",
	},
	Deviding: operations.OperationInfo{
		Number: 4,
		Description: "Devide first number on second",
		ResultNotification: "Result of deviding %v on %v is %v.\n",
	},
	Powerizing: operations.OperationInfo{
		Number: 5,
		Description: "Raise the first number to the power of the second number",
		ResultNotification: "Result of raising %v to the power %v is %v.\n",
	},
	GettingRoot: operations.OperationInfo{
		Number: 6,
		Description: "Taking the root whose power is equal to the second number from the first number",
		ResultNotification: "Result of getting root from %v of %v degree is %v.\n",
	},
}


var calculator = &Calculator{
	PossibleOperations: possible_operations, 
	NeedToContinue: true, 
	ContinueWithResult: false,
	PossibleAnswers: possible_answers,
}


func TestGetOperationsNumbers(test *testing.T) {
	operations_numbers := calculator.getOperationsNumbers()
	expected_len := 6
	fact_len := len(operations_numbers)
	if fact_len != expected_len {
		test.Fatalf("Operations_slice length is incorrect: %v expected, but %v received", expected_len, fact_len)
	}

	if reflect.TypeOf(operations_numbers) != reflect.TypeOf([]int{}) {
		test.Fatalf("%v is not a slice of type OperationInfo", operations_numbers)
	}
}

func TestGetAnswersNumbers(test *testing.T) {
	answers_numbers := calculator.getAnswersNumbers()
	expected_len := 2
	fact_len := len(answers_numbers)
	if fact_len != expected_len {
		test.Fatalf("answers_numbers length is incorrect: %v expected, but %v received", expected_len, fact_len)
	}

	if reflect.TypeOf(answers_numbers) != reflect.TypeOf([]int{}) {
		test.Fatalf("%v is not a slice of type OperationInfo", answers_numbers)
	}
}

func calculatingAlgorithmForTest(testing_data []testingDataValue, operation_name string, union string, test *testing.T) {
	for _, data := range(testing_data) {
		calculator.FirstNumber = data.FirstNumber
		calculator.SecondNumber = data.SecondNumber
		calculator.OperationNumber = data.OperationNumber
		calculator.MakeCalculation()

		if calculator.LastOperationResult != data.ExpectedResult {
			test.Fatalf(
				"Incorrected calculation of %v %v %v %v: expected %v, but received %v",
				operation_name,
				data.FirstNumber, 
				union,
				data.SecondNumber, 
				data.ExpectedResult, 
				calculator.LastOperationResult)
		}

		calculator.LastOperationResult = TypeOfNumber{}
	}
}

func TestSummarizing(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 12,
			},
			ExpectedResult: TypeOfNumber{
				intField: 22,
			},
			OperationNumber: 1,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 22.5,
			},
			OperationNumber: 1,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.0,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 2.0,
			},
			OperationNumber: 1,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 5,
			},
			ExpectedResult: TypeOfNumber{
				intField: -5,
			},
			OperationNumber: 1,
		},
	}

	calculatingAlgorithmForTest(testing_data, "summarizing", "and", test)
}

func TestSubstructing(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 12,
			},
			ExpectedResult: TypeOfNumber{
				intField: -2,
			},
			OperationNumber: 2,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -2.5,
			},
			OperationNumber: 2,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.0,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -22.0,
			},
			OperationNumber: 2,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 5,
			},
			ExpectedResult: TypeOfNumber{
				intField: -15,
			},
			OperationNumber: 2,
		},
	}

	calculatingAlgorithmForTest(testing_data, "substructing", "and", test)
}

func TestMultiplying(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 12,
			},
			ExpectedResult: TypeOfNumber{
				intField: 120,
			},
			OperationNumber: 3,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 125.0,
			},
			OperationNumber: 3,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 12.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -125,
			},
			OperationNumber: 3,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 5,
			},
			ExpectedResult: TypeOfNumber{
				intField: -50,
			},
			OperationNumber: 3,
		},
	}

	calculatingAlgorithmForTest(testing_data, "multiplying", "on", test)
}

func TestDeviding(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 2,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 5,
			},
			OperationNumber: 4,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 2.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 4.0,
			},
			OperationNumber: 4,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: 2.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -4.0,
			},
			OperationNumber: 4,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: -5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 2.0,
			},
			OperationNumber: 4,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 0,
			},
			ExpectedResult: TypeOfNumber{
				floatField: math.Inf(-1),
			},
			OperationNumber: 4,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 0,
			},
			ExpectedResult: TypeOfNumber{
				floatField: math.Inf(1),
			},
			OperationNumber: 4,
		},
	}

	calculatingAlgorithmForTest(testing_data, "deviding", "on", test)
}

func TestPowerizing(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				intField: 2,
			},
			ExpectedResult: TypeOfNumber{
				intField: 100,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 2,
			},
			ExpectedResult: TypeOfNumber{
				intField: 100,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: -2,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -0.01,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: 3,
			},
			ExpectedResult: TypeOfNumber{
				intField: -1000,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				intField: -3,
			},
			ExpectedResult: TypeOfNumber{
				floatField: -0.001,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 4,
			},
			SecondNumber: TypeOfNumber{
				floatField: 0.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 2,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 4,
			},
			SecondNumber: TypeOfNumber{
				floatField: -0.5,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 0.5,
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -10,
			},
			SecondNumber: TypeOfNumber{
				floatField: math.Inf(-1),
			},
			ExpectedResult: TypeOfNumber{},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: math.Inf(1),
			},
			ExpectedResult: TypeOfNumber{
				floatField: math.Inf(1),
			},
			OperationNumber: 5,
		},
	}

	calculatingAlgorithmForTest(testing_data, "powerizing", "to the power", test)
}

func TestGetRoot(test *testing.T) {
	testing_data := []testingDataValue{
		{
			FirstNumber: TypeOfNumber{
				intField: 4,
			},
			SecondNumber: TypeOfNumber{
				intField: 2,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 2,
			},
			OperationNumber: 6,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: -4,
			},
			SecondNumber: TypeOfNumber{
				intField: 2,
			},
			ExpectedResult: TypeOfNumber{},
			OperationNumber: 6,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 8,
			},
			SecondNumber: TypeOfNumber{
				floatField: 3,
			},
			ExpectedResult: TypeOfNumber{
				floatField: 1.9999999999999998,
			},
			OperationNumber: 6,
		},
		{
			FirstNumber: TypeOfNumber{
				intField: 10,
			},
			SecondNumber: TypeOfNumber{
				floatField: math.Inf(1),
			},
			ExpectedResult: TypeOfNumber{
				floatField: math.Inf(1),
			},
			OperationNumber: 5,
		},
		{
			FirstNumber: TypeOfNumber{
				floatField: math.Inf(1),
			},
			SecondNumber: TypeOfNumber{
				floatField: 2,
			},
			ExpectedResult: TypeOfNumber{
				floatField: math.Inf(1),
			},
			OperationNumber: 5,
		},
	}

	calculatingAlgorithmForTest(testing_data, "getting from", "root of power", test)
}

func TestRefreshNumbers(test *testing.T) {
	calculator.FirstNumber.intField = 10
	calculator.SecondNumber.intField = 20
	calculator.refreshNumbers()

	if calculator.FirstNumber.intField != 0 || calculator.SecondNumber.intField != 0 {
		test.Fatalf(
			"Expected to refresh numbers, but got %v and %v", 
			calculator.FirstNumber, 
			calculator.FirstNumber)
	}
}

func TestUpdateFirstNumber(test *testing.T) {
	var  expected_number int64 = 20
	calculator.LastOperationResult.intField = expected_number
	calculator.UpdateFirstNumber()

	if calculator.FirstNumber.intField != expected_number {
		test.Fatalf(
			"Expected to update first number to zero, but got %v", 
			expected_number)
	}
}