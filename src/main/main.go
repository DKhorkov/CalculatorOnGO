package main

import (
	"structures"
)


var answers structures.Answers = structures.Answers{Yes: "yes", No: "no"}

var operations structures.Operations = structures.Operations{
	Summarizing: structures.OperationNumberAndDescription{
		OperationNumber: 1,
		Description: "Summarize two entered numbers",
		OperationName : "summarizing",
	},
	Subtracting: structures.OperationNumberAndDescription{
		OperationNumber: 2,
		Description: "Subtracte two entered numbers",
		OperationName : "subtracting",
	},
	Multipling: structures.OperationNumberAndDescription{
		OperationNumber: 3,
		Description: "Multiply two entered numbers",
		OperationName : "multipling",
	},
	Deviding: structures.OperationNumberAndDescription{
		OperationNumber: 4,
		Description: "Devide two entered numbers",
		OperationName : "deviding",
	},
	Powerizing: structures.OperationNumberAndDescription{
		OperationNumber: 5,
		Description: "Powerize two entered numbers",
		OperationName : "powerizing",
	},
	Squaring: structures.OperationNumberAndDescription{
		OperationNumber: 6,
		Description: "Square two entered numbers",
		OperationName : "squaring",
	},
}



func main() {
	calculator := &structures.Calculator{PossibleOperations: operations}
	
	// answers.ShowFields()
	calculator.GreetUser()
	calculator.GetFirstNumber()
	calculator.GetSecondNumber()
	// calculator.GetOperationNumber()
	calculator.Summarize()
}