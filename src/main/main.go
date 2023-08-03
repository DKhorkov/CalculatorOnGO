package main

import (
	"fmt"
	"structures"
)


var answers structures.Answers = structures.Answers{
	Yes: structures.AnswerInfo{
		Number: 1, 
		Description: "Yes",
		}, 
	No: structures.AnswerInfo{
		Number: 2, 
		Description: "No",
		}, 
}

var operations structures.Operations = structures.Operations{
	Summarizing: structures.OperationInfo{
		Number: 1,
		Description: "Summarize two entered numbers",
		ResultNotification: "\nResult of summarizing %v and %v is %v.\n",
	},
	Subtracting: structures.OperationInfo{
		Number: 2,
		Description: "Subtracte two entered numbers",
		ResultNotification: "\nResult of substracting from %v number %v is %v.\n",
	},
	Multipling: structures.OperationInfo{
		Number: 3,
		Description: "Multiply two entered numbers",
		ResultNotification: "\nResult of multiplying %v on %v is %v.\n",
	},
	Deviding: structures.OperationInfo{
		Number: 4,
		Description: "Devide two entered numbers",
		ResultNotification: "\nResult of deviding %v on %v is %v.\n",
	},
	Powerizing: structures.OperationInfo{
		Number: 5,
		Description: "Powerize two entered numbers",
		ResultNotification: "\nResult of powerizing %v in %v degree is %v.\n",
	},
	Squaring: structures.OperationInfo{
		Number: 6,
		Description: "Square two entered numbers",
		ResultNotification: "\nResult of getting root from %v of %v degree is %v.\n",
	},
}



func main() {
	calculator := &structures.Calculator{
		PossibleOperations: operations, 
		NeedToContinue: true, 
		ContinueWithResult: false,
		PossibleAnswers: answers,
	}
	
	calculator.GreetUser()
	for calculator.NeedToContinue {
		if !calculator.ContinueWithResult {
			calculator.GetFirstNumber()
			baseAlgorithm(calculator)
		} else {
			calculator.UpdateFirstNumber()
			baseAlgorithm(calculator)
		}
	}

	fmt.Println("\nThanks for using! See you soon!")
}


func baseAlgorithm(calculator *structures.Calculator) {
	calculator.GetSecondNumber()
	calculator.MakeCalculation()
	calculator.CheckNeedToContinue()

	if !calculator.NeedToContinue {
		return
	}

	calculator.CheckContinueWithResult()
}