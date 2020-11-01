package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ComputeHandler handles input io.Reader and output io.Writer for computation
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute method reads the expression from input and write the computed result to the output.
func (ch *ComputeHandler) Compute() error {
	expr, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	if res, err := PostfixToPrefix(string(expr)); err != nil {
		return err
	} else {
		buf := []byte(fmt.Sprintf("%s", res))
		if _, e := ch.Output.Write(buf); e != nil {
			return e
		}
	}

	return nil
}
