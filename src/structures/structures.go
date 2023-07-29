package structures

import (
	"fmt"
	"functions"
	"reflect"
)


type Answers struct {
	Yes string
	No string
}


/*
	С помощью пакета "reflect" получаем значения структуры, а далее итерируемся по ним и выводим в STDOUT.
*/
func (answers Answers) ShowFields() {
	values := reflect.ValueOf(answers)

	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i).Interface()
		fmt.Printf("%v) %v\n", i + 1, value)
	}
}


type OperationNumberAndDescription struct {
	OperationNumber int
	Description string
}

func (operation *OperationNumberAndDescription) getDescription() string {
	return operation.Description
}


type Operations struct {
	Summarizing OperationNumberAndDescription
	Subtracting OperationNumberAndDescription
	Multipling OperationNumberAndDescription
	Deviding OperationNumberAndDescription
	Powerizing OperationNumberAndDescription
	Squaring OperationNumberAndDescription
}

func (operations * Operations) ShowOperations() {}


/*
	Дженерик, чтобы можно было работать со всеми числовыми типами данных. 
	Имеет два поля, которые являются типом "interfaces.Number"
*/
type Calculator struct {
	FirstNumber, SecondNumber float64
	PossibleOperations Operations
}

func (calculator *Calculator) GetFirstNumber() {
	fmt.Println("Please, enter the first number:")
	calculator.FirstNumber = calculator.scanNumber()
}

func (calculator *Calculator) GetSecondNumber() {
	fmt.Println("Please, enter the second number:")
	calculator.SecondNumber = calculator.scanNumber()
}

func (calculator *Calculator) scanNumber() float64 {
	number, err := functions.GetInputLine()
	for err != nil {
		fmt.Println("Error, you should enter a number! Please, tru again:")
		number, err = functions.GetInputLine()
	}

	return number
}

func (calculator Calculator) ShowPossibleOperations() {

}