package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

// TestHandlerCorrect tests ComputeHandler.Compute for correct evaluation
func (s *MySuite) TestHandlerCorrect(c *C) {
	expr := "105 7 * 12 7 / +"
	input := strings.NewReader(expr)

	buf := new(bytes.Buffer)

	handler := ComputeHandler{
		Input: input,
		Output: buf,
	}

	handler.Compute()

	c.Assert(buf.String(), Equals, "+ / 7 12 * 7 105")
}

// TestHandlerIncorrect confirms that ComputeHandler.Compute
// will return error if there is syntax error in input
func (s *MySuite) TestHandlerIncorrect(c *C) {
	expr := "- 10 5"
	input := strings.NewReader(expr)

	buf := new(bytes.Buffer)

	handler := ComputeHandler{
		Input: input,
		Output: buf,
	}

	err := handler.Compute()

	c.Assert(err, NotNil)
}
