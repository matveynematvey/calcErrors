package calc

import (
	"errors"
	"fmt"
)

var ErrResNotInt = errors.New("result is not integer")
var ErrZero = errors.New("divided by zero")

func Add(a int, b int) (string, error) {
	return fmt.Sprintf("%d + %d = %d", a, b, a+b), nil
}

func Div(a int, b int) (string, error) {
	if b == 0 {
		return "", fmt.Errorf("can't div %d by %d : %v", a, b, ErrZero)
	}
	if (a % b) != 0 {
		return "", fmt.Errorf("can't div %d by %d : %v", a, b, ErrResNotInt)
	}
	return fmt.Sprintf("%d / %d = %d", a, b, a/b), nil
}
