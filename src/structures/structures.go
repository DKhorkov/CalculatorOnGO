package structures

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/exp/slices"
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


type Operations struct {
	Summarizing OperationInfo
	Subtracting OperationInfo
	Multipling OperationInfo
	Deviding OperationInfo
	Powerizing OperationInfo
	Squaring OperationInfo
}


type OperationInfo struct {
	OperationNumber int
	Description string
	ResultNotification string
}


/*
	Обязательно передавать объект Operations, а не указатель на него, иначе произойдет паника.
	В данном методе мы итерируемся по объекту ValueOf пакета reflect. Для каждого поля структуры мы с помощью метода 
	Field()	достаем вложенное поле структуры (тоже структура). Но, чтобы компилятор работал корректно, ведь мы возвращаем срез типа []OperationNumberAndDescription, а не срез пустых интерфейсов []interface{}, то необходимо привести данный интерфейс к его реализации (в GO любой объект реализует пустой интерфейс):

				struct_fields.Field(i).Interface().(OperationNumberAndDescription)

	Итерация, которая подходит для данной реализации:
	https://ru.stackoverflow.com/questions/1026882/%D0%A6%D0%B8%D0%BA%D0%BB-for-%D0%BF%D0%BE-%D0%BF%D0%BE%D0%BB%D1%8F%D0%BC-%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D1%8B

	Другие варианты, котоыре не подходят для данной реализации, но стоит знать:
	https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go
	https://stackoverflow.com/questions/50098624/reflect-call-of-reflect-value-fieldbyname-on-ptr-value
*/
func (operations Operations) toSlice() []OperationInfo {
	struct_fields := reflect.ValueOf(operations)
	slice := make([]OperationInfo, struct_fields.NumField())

	for i := 0; i < struct_fields.NumField(); i++ {
		field := struct_fields.Field(i).Interface().(OperationInfo)
		slice[i] = field
	}

	return slice
}

func (operations Operations) ShowOperations() {
	fmt.Println("\nPossible operations:")

	operations_slice := operations.toSlice()
	for _, operation := range operations_slice {
		fmt.Printf("%v - %v\n", operation.OperationNumber, operation.Description)
	}
}


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
	FirstNumber, SecondNumber TypeOfNumber
	OperationNumber int
	PossibleOperations Operations
	LastOperationResult TypeOfNumber
}

func (calculator Calculator) GreetUser() {
	fmt.Println("\nWelcome to calculator on GO!")
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
	Метод возваращет один из типов данных, которые могут реализовать интерфейс "Number", а также ошибку.
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

func (calculator Calculator) showPossibleOperations() {
	calculator.PossibleOperations.ShowOperations()
}

func (calculator Calculator) MakeCalculation() {
	calculator.getOperationNumber()
	switch calculator.OperationNumber {
	case calculator.PossibleOperations.Summarizing.OperationNumber:
		calculator.summarize()
	case calculator.PossibleOperations.Subtracting.OperationNumber:
		calculator.substract()
	case calculator.PossibleOperations.Multipling.OperationNumber:
		calculator.multiply()
	case calculator.PossibleOperations.Deviding.OperationNumber:
		calculator.devide()
	case calculator.PossibleOperations.Powerizing.OperationNumber:
		calculator.powerize()
	case calculator.PossibleOperations.Squaring.OperationNumber:
		calculator.square()
	}
}

func (calculator *Calculator) getOperationNumber() {
	operations_numbers := calculator.getOperationsNumbers()

	calculator.showPossibleOperations()
	fmt.Print("\nPlease, choose operation number from list, presented above: ")

	chosen_operation_number, err := calculator.scanOperationNumber()
	for err != nil || !slices.Contains(operations_numbers, chosen_operation_number) {
		fmt.Print("\nError, you should enter an integer number from list above! Please, try again: ")
		chosen_operation_number, err = calculator.scanOperationNumber()
	}

	calculator.OperationNumber = chosen_operation_number
}

func (calculator Calculator) getOperationsNumbers() []int {
	operations_slice := calculator.PossibleOperations.toSlice()
	operations_numbers := make([]int, len(operations_slice))
	for index, operation := range operations_slice {
		operations_numbers[index] = operation.OperationNumber
	}

	return operations_numbers
}

func (calculator Calculator) scanOperationNumber() (int, error) {
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
	if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = 
			calculator.FirstNumber.intField / 
			calculator.SecondNumber.intField
	} else if calculator.FirstNumber.intField != 0 {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.intField) /
			calculator.SecondNumber.floatField
	} else if calculator.SecondNumber.intField != 0 {
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

func (calculator *Calculator) powerize() {
	if calculator.SecondNumber.intField == 0 && calculator.SecondNumber.floatField == 0 {
		calculator.LastOperationResult.intField = 1
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

func (calculator *Calculator) square() {
	if calculator.SecondNumber.intField == 0 && calculator.SecondNumber.floatField == 0 {
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

	calculator.printResult(calculator.PossibleOperations.Squaring)
	calculator.refreshNumbers()
}

func (calculator *Calculator) refreshNumbers() {
	calculator.FirstNumber, calculator.SecondNumber = TypeOfNumber{}, TypeOfNumber{}
}

func (calculator Calculator) printResult(operation_info OperationInfo) {
	if calculator.LastOperationResult.intField != 0 {
		fmt.Printf(
			operation_info.ResultNotification, 
			calculator.FirstNumber.intField,
			calculator.SecondNumber.intField,
			calculator.LastOperationResult.intField)
	} else if calculator.FirstNumber.intField != 0 && operation_info.OperationNumber == 6 {
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
}
