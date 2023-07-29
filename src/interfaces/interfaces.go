package interfaces

/*
	Используем интерфейс типов данных, чтобы иметь возможность совершать арифметические операции с любым числовым типом данных.
*/
type Number interface {
	int | float32 | float64 | byte
}


/*
	Используем дженерик-интерфейс, чтобы методы объекта, реализующего данный интерфейс возвращали нам объект типа "Number".
*/
type Calculator[number Number] interface {
	Summarize() number
	Subtract() number
	Multiply() number
	Devide() number
	Powerize() number
	Square() number
}