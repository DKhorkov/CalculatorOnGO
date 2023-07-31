package structures

import (
	"bufio"
	"fmt"
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


type OperationNumberAndDescription struct {
	OperationNumber int
	Description string
	OperationName string
}


type Operations struct {
	Summarizing OperationNumberAndDescription
	Subtracting OperationNumberAndDescription
	Multipling OperationNumberAndDescription
	Deviding OperationNumberAndDescription
	Powerizing OperationNumberAndDescription
	Squaring OperationNumberAndDescription
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
func (operations Operations) toSlice() []OperationNumberAndDescription {
	struct_fields := reflect.ValueOf(operations)
	slice := make([]OperationNumberAndDescription, struct_fields.NumField())

	for i := 0; i < struct_fields.NumField(); i++ {
		field := struct_fields.Field(i).Interface().(OperationNumberAndDescription)
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
	uintField uint64
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
	fmt.Println("\nPlease, enter the first number:")
	calculator.FirstNumber = calculator.getNumber()
}

func (calculator *Calculator) GetSecondNumber() {
	fmt.Println("\nPlease, enter the second number:")
	calculator.SecondNumber = calculator.getNumber()
}

func (calculator *Calculator) getNumber() TypeOfNumber {
	number, err := calculator.scanNumber()
	for err != nil {
		fmt.Println("\nError, you should enter a number! Please, try again:")
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

	uint_number, err := strconv.ParseUint(input, 10, 8)

	if err == nil {
		return TypeOfNumber{uintField: uint_number}, err 
	}

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

func (calculator *Calculator) GetOperationNumber() {
	operations_numbers := calculator.getOperationsNumbers()

	calculator.showPossibleOperations()
	fmt.Println("Please, choose operation number from list, presented above:")

	chosen_operation_number, err := calculator.scanOperationNumber()
	for err != nil || !slices.Contains(operations_numbers, chosen_operation_number) {
		fmt.Println("\nError, you should enter an integer number from list above! Please, try again:")
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

func (calculator *Calculator) Summarize() {
	if calculator.FirstNumber.uintField != 0 && calculator.SecondNumber.uintField != 0 {
		calculator.LastOperationResult.uintField = calculator.FirstNumber.uintField + calculator.SecondNumber.uintField
	} else if calculator.FirstNumber.intField != 0 && calculator.SecondNumber.intField !=0 {
		calculator.LastOperationResult.intField = calculator.FirstNumber.intField + calculator.SecondNumber.intField
	} else {
		calculator.LastOperationResult.floatField = 
			float64(calculator.FirstNumber.uintField) + 
			float64(calculator.SecondNumber.uintField) + 
			float64(calculator.FirstNumber.intField) + 
			float64(calculator.SecondNumber.intField) + 
			calculator.FirstNumber.floatField + 
			calculator.SecondNumber.floatField
	}

	calculator.refreshNumbers()

	calculator.printResult(calculator.PossibleOperations.Summarizing.OperationName)
}

func (calculator *Calculator) refreshNumbers() {
	calculator.FirstNumber, calculator.SecondNumber = TypeOfNumber{}, TypeOfNumber{}
}

func (calculator Calculator) printResult(operation_name string) {
	if calculator.LastOperationResult.uintField != 0 {
		fmt.Printf("\nResult of %v is %v.\n", operation_name, calculator.LastOperationResult.uintField)
	} else if calculator.LastOperationResult.intField != 0 {
		fmt.Printf("\nResult of %v is %v.\n", operation_name, calculator.LastOperationResult.intField)
	} else {
		fmt.Printf("\nResult of %v is %v.\n", operation_name, calculator.LastOperationResult.floatField)
	}
}
