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
	Subtracting: structures.OperationNumberAndDescription{
		OperationNumber: 2,
		Description: "Subtracting two entered numbers",
	},
	Multipling: structures.OperationNumberAndDescription{
		OperationNumber: 3,
		Description: "Multipling two entered numbers",
	},
	Deviding: structures.OperationNumberAndDescription{
		OperationNumber: 4,
		Description: "Deviding two entered numbers",
	},
	Powerizing: structures.OperationNumberAndDescription{
		OperationNumber: 5,
		Description: "Powerizing two entered numbers",
	},
	Squaring: structures.OperationNumberAndDescription{
		OperationNumber: 6,
		Description: "Squaring two entered numbers",
	},
}



func main() {
	calculator := &structures.Calculator{PossibleOperations: operations}
	
	// answers.ShowFields()
	calculator.GetFirstNumber()
	calculator.ShowPossibleOperations()

}