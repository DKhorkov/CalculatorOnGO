package calculator

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"answers"
	"operations"

	"golang.org/x/exp/slices"
)

/*
	Поскольку нельзя поставить Дженерик тип данных в возврат из функции, приходится обходитсья структурой с полями разного типа.
*/
type TypeOfNumber struct {
	intField int64
	floatField float64
}


/*
	Дженерик, чтобы можно было работать со всеми числовыми типами данных. 
	Имеет два поля, которые являются типом "interfaces.Number"
*/
type Calculator struct {
	FirstNumber, SecondNumber, LastOperationResult TypeOfNumber
	OperationNumber int
	PossibleOperations operations.Operations
	NeedToContinue bool
	ContinueWithResult bool
	PossibleAnswers answers.Answers
}

func (calculator Calculator) GreetUser() {
	calculator.createDevidingLine()
	fmt.Println("Welcome to calculator on GO!")
	calculator.createDevidingLine()
}

func (calculator Calculator) SayGoodbyeToUser() {
	fmt.Println("Thanks for using! See you soon!")
	calculator.createDevidingLine()
}

func (calculator *Calculator) GetFirstNumber() {
	fmt.Print("\nPlease, enter the first number: ")
	calculator.FirstNumber = calculator.getNumber()
}

func (calculator *Calculator) GetSecondNumber() {
	fmt.Print("\nPlease, enter the second number: ")
	calculator.SecondNumber = calculator.getNumber()
}

func (calculator *Calculator) getNumber() TypeOfNumber {
	number, err := calculator.scanNumber()
	for err != nil {
		fmt.Print("\nError, you should enter a number! Please, try again: ")
		number, err = calculator.scanNumber()
	}

	return number
}

/*
	Метод возваращет структуру TypeOfNumber, чтобы в дальнейшем можно было работать как с int64, так и с float64 типами данных, которые находятся внутри данной структуры. Также метод возвращает ошибку, если такая имеется.
*/
func (calculator Calculator) scanNumber() (TypeOfNumber, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()	
	input := scanner.Text()

	int_number, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		return TypeOfNumber{intField: int_number}, err 
	}

	float_number, err := strconv.ParseFloat(input, 64)
	return TypeOfNumber{floatField: float_number}, err 
}

func (calculator *Calculator) MakeCalculation() {
	calculator.getOperationNumber()
	switch calculator.OperationNumber {
	case calculator.PossibleOperations.Summarizing.Number:
		calculator.summarize()
	case calculator.PossibleOperations.Subtracting.Number:
		calculator.substract()
	case calculator.PossibleOperations.Multipling.Number:
		calculator.multiply()
	case calculator.PossibleOperations.Deviding.Number:
		calculator.devide()
	case calculator.PossibleOperations.Powerizing.Number:
		calculator.powerize()
	case calculator.PossibleOperations.GettingRoot.Number:
		calculator.getRoot()
	}
}

/*
	Метод получает один из доступных номеров для совершения операции. Если указан некорректный номер или другая невалидная информация, будет запущен бесконечный цикл до тех пор, пока ответ не будет валидным.
*/
func (calculator *Calculator) getOperationNumber() {
	operations_numbers := calculator.getOperationsNumbers()

	calculator.PossibleOperations.ShowOperations()
	fmt.Print("\nPlease, choose operation number from list, presented above: ")

	chosen_operation_number, err := calculator.scanOperationOrAnswerNumber()
	for err != nil || !slices.Contains(operations_numbers, chosen_operation_number) {
		fmt.Print("\nError, you should enter an integer number from list above! Please, try again: ")
		chosen_operation_number, err = calculator.scanOperationOrAnswerNumber()
	}

	calculator.OperationNumber = chosen_operation_number
}

/*
	Метод позволяет получить срез из доступных для использования номеров операций для дальнейшей реализации логики расчетов калькулятора.
*/
func (calculator Calculator) getOperationsNumbers() []int {
	operations_slice := calculator.PossibleOperations.ToSlice()
	operations_numbers := make([]int, len(operations_slice))
	for index, operation := range operations_slice {
		operations_numbers[index] = operation.Number
	}

	return operations_numbers
}

/*
	Метод сканирует вводимое юзером потенциальное число. Если введено не число или число, которое нельзя привести к целочисленному, будет возварщен 0 и ошибка.
*/
func (calculator Calculator) scanOperationOrAnswerNumber() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()	
	input := scanner.Text()
	number, err := strconv.ParseInt(input, 10, 8)
	return int(number), err
}

func (calculator *Calculator) summarize() {
	 if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = 
			calculator.FirstNumber.intField + 
			calculator.SecondNumber.intField
	} else {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) + 
			float64(calculator.SecondNumber.intField) + 
			calculator.FirstNumber.floatField + 
			calculator.SecondNumber.floatField
	}

	calculator.printResult(calculator.PossibleOperations.Summarizing)
	calculator.refreshNumbers()
}

func (calculator *Calculator) substract() {
	if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = 
			calculator.FirstNumber.intField - 
			calculator.SecondNumber.intField
	} else {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) -
			float64(calculator.SecondNumber.intField) + 
			calculator.FirstNumber.floatField - 
			calculator.SecondNumber.floatField
	}

	calculator.printResult(calculator.PossibleOperations.Subtracting)
	calculator.refreshNumbers()
}

func (calculator *Calculator) multiply() {
	if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = 
			calculator.FirstNumber.intField * 
			calculator.SecondNumber.intField
	} else if calculator.FirstNumber.intField != 0 {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) * 
			calculator.SecondNumber.floatField
	} else if calculator.SecondNumber.intField != 0 {
		calculator.LastOperationResult.floatField = 
			calculator.FirstNumber.floatField * 
			float64(calculator.SecondNumber.intField)
	} else {
		calculator.LastOperationResult.floatField = 
			calculator.FirstNumber.floatField * 
			calculator.SecondNumber.floatField
	}

	calculator.printResult(calculator.PossibleOperations.Multipling)
	calculator.refreshNumbers()
}

func (calculator *Calculator) devide() {
	if calculator.FirstNumber.intField == 0 && calculator.FirstNumber.floatField == 0 {
		calculator.LastOperationResult.intField = 0
	} else if calculator.SecondNumber.intField == 0 && calculator.SecondNumber.floatField == 0 {
		calculator.devideByZero()
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) / 
			float64(calculator.SecondNumber.intField)
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.floatField !=0 {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) /
			calculator.SecondNumber.floatField
	} else if calculator.FirstNumber.floatField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.floatField = 
			calculator.FirstNumber.floatField /
			float64(calculator.SecondNumber.intField)
	} else {
		calculator.LastOperationResult.floatField = 
			calculator.FirstNumber.floatField / 
			calculator.SecondNumber.floatField
	}

	calculator.printResult(calculator.PossibleOperations.Deviding)
	calculator.refreshNumbers()
}

func (calculator *Calculator) devideByZero() {
	if calculator.FirstNumber.intField > 0 || calculator.FirstNumber.floatField > 0 {
		calculator.LastOperationResult.floatField = math.Inf(1)
	} else 	if calculator.FirstNumber.intField < 0 || calculator.FirstNumber.floatField < 0 {
		calculator.LastOperationResult.floatField = math.Inf(-1)
	}
}

func (calculator *Calculator) powerize() {
	if calculator.FirstNumber.floatField == math.Inf(1)  {
		calculator.LastOperationResult.floatField = math.Inf(1)
	} else if calculator.FirstNumber.floatField == math.Inf(-1) {
		calculator.LastOperationResult.floatField = math.Inf(-1)
	} else if calculator.SecondNumber.intField == 0 && calculator.SecondNumber.floatField == 0 {
		calculator.LastOperationResult.intField = 1
	} else if calculator.SecondNumber.intField < 0 || calculator.SecondNumber.floatField < 0 {
		calculator.raiseToNegativePower()
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = 
		int64(
			math.Pow(
				float64(calculator.FirstNumber.intField), 
				float64(calculator.SecondNumber.intField)))
	} else if calculator.FirstNumber.intField != 0 {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				float64(calculator.FirstNumber.intField), 
				calculator.SecondNumber.floatField)
	} else if calculator.SecondNumber.intField != 0 {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				calculator.FirstNumber.floatField, 
				float64(calculator.SecondNumber.intField))
	} else {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				calculator.FirstNumber.floatField, 
				calculator.SecondNumber.floatField)
	}

	calculator.printResult(calculator.PossibleOperations.Powerizing)
	calculator.refreshNumbers()
}

func (calculator *Calculator) raiseToNegativePower() {
	if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField < 0 {
		calculator.LastOperationResult.floatField = math.Pow(
				float64(calculator.FirstNumber.intField), 
				float64(calculator.SecondNumber.intField))
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.floatField < 0 {
		calculator.LastOperationResult.floatField = math.Pow(
				float64(calculator.FirstNumber.intField), 
				calculator.SecondNumber.floatField)
	} else if calculator.FirstNumber.floatField != 0 && calculator.SecondNumber.intField < 0 {
		calculator.LastOperationResult.floatField = math.Pow(
				calculator.FirstNumber.floatField, 
				float64(calculator.SecondNumber.intField))
	} else {
		calculator.LastOperationResult.floatField = math.Pow(
				calculator.FirstNumber.floatField, 
				calculator.SecondNumber.floatField)
	}

	if (calculator.FirstNumber.intField < 0 || calculator.FirstNumber.floatField < 0) && 
		calculator.LastOperationResult.floatField > 0 {
			calculator.LastOperationResult.floatField *= -1
	}
}

func (calculator *Calculator) getRoot() {
	if calculator.FirstNumber.floatField == math.Inf(1)  {
		calculator.LastOperationResult.floatField = math.Inf(1)
	} else if calculator.FirstNumber.floatField == math.Inf(-1) {
		calculator.LastOperationResult.floatField = math.Inf(-1)
	} else if calculator.SecondNumber.intField == 0 && calculator.SecondNumber.floatField == 0 {
		calculator.LastOperationResult.floatField = math.Inf(1)
	} else if calculator.FirstNumber.intField == 0 && calculator.FirstNumber.floatField == 0 {
		calculator.LastOperationResult.intField = 0
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.floatField = 
		math.Pow(
			float64(calculator.FirstNumber.intField), 
			1.0 / float64(calculator.SecondNumber.intField))
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.floatField !=0 {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				float64(calculator.FirstNumber.intField), 
				1.0 / calculator.SecondNumber.floatField)
	} else if calculator.FirstNumber.floatField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				calculator.FirstNumber.floatField, 
				1.0 / float64(calculator.SecondNumber.intField))
	} else {
		calculator.LastOperationResult.floatField = 
			math.Pow(
				calculator.FirstNumber.floatField, 
				1.0 / calculator.SecondNumber.floatField)
	}

	calculator.printResult(calculator.PossibleOperations.GettingRoot)
	calculator.refreshNumbers()
}

func (calculator *Calculator) refreshNumbers() {
	calculator.FirstNumber, calculator.SecondNumber = TypeOfNumber{}, TypeOfNumber{}
}

func (calculator Calculator) printResult(operation_info operations.OperationInfo) {
	calculator.createDevidingLine()

	if calculator.LastOperationResult.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.intField)
	} else if calculator.LastOperationResult.floatField != 0 && operation_info.Number == 5 {
		calculator.printRaisingInNegativePower(operation_info)
	} else if calculator.FirstNumber.intField != 0 && (operation_info.Number == 6 || operation_info.Number == 4) {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.floatField)
	} else if calculator.FirstNumber.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.floatField,
			calculator.LastOperationResult.floatField)
	} else if calculator.SecondNumber.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.floatField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.floatField)
	} else {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.floatField,
			calculator.SecondNumber.floatField,
			calculator.LastOperationResult.floatField)
	}

	calculator.createDevidingLine()
}

func (calculator Calculator) printRaisingInNegativePower(operation_info operations.OperationInfo) {
	if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.floatField)
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.floatField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.floatField,
			calculator.LastOperationResult.floatField)
	} else if calculator.FirstNumber.floatField != 0 && calculator.SecondNumber.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.floatField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.floatField)
	} else {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.floatField,
			calculator.SecondNumber.floatField,
			calculator.LastOperationResult.floatField)
	}
}

func (calculator Calculator) createDevidingLine() {
	fmt.Println("--------------------------------------------------------------------------------------")
}

/*
	Метод проверяет, желает ли юзер совершить еще какой-либо расчет. В зависимости от этого меняется флаг, используемый для основного алгоритма работы калькулятора.
*/
func (calculator *Calculator) CheckNeedToContinue() {
	fmt.Println("Do you want to make another calculation?")
	answer_number := calculator.getAnswerNumber()
	switch answer_number {
	case calculator.PossibleAnswers.Yes.Number:
		calculator.NeedToContinue = true
	case calculator.PossibleAnswers.No.Number:
		calculator.NeedToContinue = false
	}

	calculator.createDevidingLine()
}

/*
	Метод получает число, относящееся к одному из возможных вариантов ответа на вопрос. Если юзер указал невалдиное число, то начнется работа бесконечного цикла, пока не будет получено валидное число.
*/
func (calculator Calculator) getAnswerNumber() int {
	answers_numbers := calculator.getAnswersNumbers()

	calculator.PossibleAnswers.ShowAnswers()
	fmt.Print("\nPlease, choose answer number from list, presented above: ")

	chosen_answer_number, err := calculator.scanOperationOrAnswerNumber()
	for err != nil || !slices.Contains(answers_numbers, chosen_answer_number) {
		fmt.Print("\nError, you should enter an integer number from list above! Please, try again: ")
		chosen_answer_number, err = calculator.scanOperationOrAnswerNumber()
	}

	return chosen_answer_number
}

/*
	Метод позволяет получить числа, которые относятся к структуре с ответами, чтобы в дальнейшем оперировать данным номерами и проверять, корректное ли число ввел юзер.
*/
func (calculator Calculator) getAnswersNumbers() []int {
	answers_slice := calculator.PossibleAnswers.ToSlice()
	answers_numbers := make([]int, len(answers_slice))
	for index, operation := range answers_slice {
		answers_numbers[index] = operation.Number
	}

	return answers_numbers
}

/*
	Метод проверяет, желает ли юзер продолжить работу с полученным ранее результатом. В зависимости от этого меняется флаг, используемый для основного алгоритма работы калькулятора.
*/
func (calculator *Calculator) CheckContinueWithResult() {
	question_text := "\nDo you want make next operation with the last operation result: %v?\n"
	if calculator.LastOperationResult.intField != 0 {
		fmt.Printf(
			question_text, 
			calculator.LastOperationResult.intField)
	} else {
		fmt.Printf(
			question_text, 
			calculator.LastOperationResult.floatField)
	}

	answer_number := calculator.getAnswerNumber()
	switch answer_number {
	case calculator.PossibleAnswers.Yes.Number:
		calculator.ContinueWithResult = true
	case calculator.PossibleAnswers.No.Number:
		calculator.ContinueWithResult = false
	}

	calculator.createDevidingLine()
}

/*
	Метод перезаписывает результат последней операции в первое число, чтобы можно было совершать дальнейшие операции с ним.
*/
func (calculator *Calculator) UpdateFirstNumber() {
	if calculator.LastOperationResult.intField != 0 {
		calculator.FirstNumber.intField = calculator.LastOperationResult.intField
	} else {
		calculator.FirstNumber.floatField = calculator.LastOperationResult.floatField
	}

	calculator.LastOperationResult = TypeOfNumber{}
}