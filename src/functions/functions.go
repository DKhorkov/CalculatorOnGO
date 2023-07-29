package functions

import (
	"bufio"
	"os"
	"strconv"
)



func GetInputLine() (float64, error) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()	
	input := scanner.Text()
	number, err := strconv.ParseFloat(input, 64)
	
	return number, err
}
