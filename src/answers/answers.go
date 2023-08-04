package answers

import (
	"fmt"
	"reflect"
)


type Answers struct {
	Yes AnswerInfo
	No AnswerInfo
}


/*
	Метод приводит поля структуры к срезу типа AnswerInfo, где первым элементом будет номер, а вторым описание. 
*/
func (answers Answers) ToSlice() []AnswerInfo {
	answers_fields := reflect.ValueOf(answers)
	slice := make([]AnswerInfo, answers_fields.NumField())

	for i := 0; i < answers_fields.NumField(); i++ {
		field := answers_fields.Field(i).Interface().(AnswerInfo)
		slice[i] = field
	}

	return slice
}

/*
	С помощью пакета "reflect" получаем значения структуры, а далее итерируемся по ним и выводим в STDOUT.
*/
func (answers Answers) ShowAnswers() {
	fmt.Println("\nPossible answers:")

	answers_slice := answers.ToSlice()
	for _, answer := range answers_slice {
		fmt.Printf("%v - %v\n", answer.Number, answer.Description)
	}
}


type AnswerInfo struct {
	Number int
	Description string
}
