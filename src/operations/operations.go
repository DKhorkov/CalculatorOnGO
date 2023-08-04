package operations

import (
	"fmt"
	"reflect"
)


type Operations struct {
	Summarizing OperationInfo
	Subtracting OperationInfo
	Multipling OperationInfo
	Deviding OperationInfo
	Powerizing OperationInfo
	GettingRoot OperationInfo
}


type OperationInfo struct {
	Number int
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
func (operations Operations) ToSlice() []OperationInfo {
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

	operations_slice := operations.ToSlice()
	for _, operation := range operations_slice {
		fmt.Printf("%v - %v\n", operation.Number, operation.Description)
	}
}