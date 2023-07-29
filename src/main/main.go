package main

import (
	"structures"
)


var answers structures.Answers = structures.Answers{Yes: "yes", No: "no"}

var operations structures.Operations = structures.Operations{
	Summarizing: structures.OperationNumberAndDescription{
		OperationNumber: 1,
		Description: "Summatize two entered numbers",
	},
}



func main() {
	calculator := &structures.Calculator{PossibleOperations: operations}
	
	answers.ShowFields()
	calculator.GetFirstNumber()
}