package main

import (
	"answers"
	"calculator"
	"operations"
)


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



func main() {
	calculator := &calculator.Calculator{
		PossibleOperations: possible_operations, 
		NeedToContinue: true, 
		ContinueWithResult: false,
		PossibleAnswers: possible_answers,
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

	calculator.SayGoodbyeToUser()
}


/*
	Базовая часть алгоритма работы калькулятора, как для использования результата предыдущей операции, так и для новой расчета с учетом первого и второго числа.
*/
func baseAlgorithm(calculator *calculator.Calculator) {
	calculator.GetSecondNumber()
	calculator.GetOperationNumber()
	calculator.MakeCalculation()
	calculator.CheckNeedToContinue()

	if !calculator.NeedToContinue {
		return
	}

	calculator.CheckContinueWithResult()
}