package functions

import (
	"bufio"
	"os"
	"strconv"
)

/*
	Поскольку нельзя поставить Дженерик тип данных в возврат из функции, приходится обходитсья структурой с полями разного типа.
*/
type TypeOfNumber struct {
	uintField uint64
	intField int64
	floatField float64
}


/*
	Метод возваращет один из типов данных, которые могут реализовать интерфейс "Number", а также ошибку.
*/
func GetInputLine() (TypeOfNumber, error) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()	
	input := scanner.Text()

	uint_number, err := strconv.ParseUint(input, 10, 8)

	if err == nil {
		return TypeOfNumber{uintField: uint_number}, err 
	}

	int_number, err := strconv.ParseInt(input, 2, 2)
	if err == nil {
		return TypeOfNumber{intField: int_number}, err 
	}

	float_number, err := strconv.ParseFloat(input, 64)
	return TypeOfNumber{floatField: float_number}, err 
}
