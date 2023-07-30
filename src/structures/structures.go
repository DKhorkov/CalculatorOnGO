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
	fmt.Println(operations.toSlice())
}


/*
	Дженерик, чтобы можно было работать со всеми числовыми типами данных. 
	Имеет два поля, которые являются типом "interfaces.Number"
*/
type Calculator struct {
	FirstNumber, SecondNumber functions.TypeOfNumber
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

func (calculator *Calculator) scanNumber() functions.TypeOfNumber {
	number, err := functions.GetInputLine()
	for err != nil {
		fmt.Println("Error, you should enter a number! Please, tru again:")
		number, err = functions.GetInputLine()
	}

	return number
}

func (calculator Calculator) ShowPossibleOperations() {
	calculator.PossibleOperations.ShowOperations()
}