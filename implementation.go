package lab2

import (
	"errors"
	"strconv"
	"strings"
)

// Simple stack structure
type Stack []string

// IsEmpty: Checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push: Add new value to end of stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop: Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("incorrect expression")
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]

		return elem, nil
	}
}

// isOperator checks if rune is one of acceptable operands
func isOperator(el string) bool {
	var ops = "+-*/^"
	return strings.Contains(ops, el)
}

// PostfixToPrefix converts postfix expression into prefix expression
//goland:noinspection ALL
func PostfixToPrefix(input string) (string, error) {
	input = strings.ReplaceAll(input, `"`, "")
	expr := strings.Fields(input)
	var stack Stack

	if len(expr) == 0 {
		return "", errors.New("empty input")
	}

	for _, el := range expr {
		if _, err := strconv.ParseFloat(el, 32); err == nil {
			stack.Push(el)
		} else if isOperator(el){
			op1, err1 := stack.Pop()
			op2, err2 := stack.Pop()

			if err1 != nil {
				return "", err1
			} else if err2 != nil {
				return "", err2
			} else {
				stack.Push(op2)
				stack.Push(op1)
				stack.Push(el)
			}
		} else {
			return "", errors.New("incorrect symbol: " + el)
		}
	}

	var ans []string

	for {
		temp, err := stack.Pop()
		if err != nil {
			break
		} else {
			ans = append(ans, temp)
		}
	}

	return strings.Join(ans, " "), nil
}
