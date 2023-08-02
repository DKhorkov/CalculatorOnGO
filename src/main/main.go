package main

import (
	"structures"
)


var answers structures.Answers = structures.Answers{Yes: "yes", No: "no"}

var operations structures.Operations = structures.Operations{
	Summarizing: structures.OperationInfo{
		OperationNumber: 1,
		Description: "Summarize two entered numbers",
		ResultNotification: "\nResult of summarizing %v and %v is %v.\n",
	},
	Subtracting: structures.OperationInfo{
		OperationNumber: 2,
		Description: "Subtracte two entered numbers",
		ResultNotification: "\nResult of substracting from %v number %v is %v.\n",
	},
	Multipling: structures.OperationInfo{
		OperationNumber: 3,
		Description: "Multiply two entered numbers",
		ResultNotification: "\nResult of multiplying %v on %v is %v.\n",
	},
	Deviding: structures.OperationInfo{
		OperationNumber: 4,
		Description: "Devide two entered numbers",
		ResultNotification: "\nResult of deviding %v on %v is %v.\n",
	},
	Powerizing: structures.OperationInfo{
		OperationNumber: 5,
		Description: "Powerize two entered numbers",
		ResultNotification: "\nResult of powerizing %v in %v degree is %v.\n",
	},
	Squaring: structures.OperationInfo{
		OperationNumber: 6,
		Description: "Square two entered numbers",
		ResultNotification: "\nResult of getting root from %v of %v degree is %v.\n",
	},
}



func main() {
	calculator := &structures.Calculator{PossibleOperations: operations}
	
	// answers.ShowFields()
	calculator.GreetUser()
	calculator.GetFirstNumber()
	calculator.GetSecondNumber()
	calculator.MakeCalculation()
}