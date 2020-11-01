package lab2

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up go-check into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// TestPostfixToPrefixCorrect runs tests that should evaluate correctly
func (s *MySuite) TestPostfixToPrefixCorrect(c *C) {
	// 10 = 10
	res, err := PostfixToPrefix("10")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "10")

	// 10 5 + = + 10 5
	res, err = PostfixToPrefix("10 5 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ 5 10")

	// 105 7 * 12 7 / + = + / 7 12 * 7 105
	res, err = PostfixToPrefix("105 7 * 12 7 / +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ / 7 12 * 7 105")

	// 21.65 5 ^ 12 7 - 45.123 + = + 45.123 - 7 12 ^ 5 21.65
	res, err = PostfixToPrefix("21.65 5 ^ 12 7 - 45.123 +")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "+ 45.123 - 7 12 ^ 5 21.65")
}

// TestPostfixToPrefixIncorrect runs incorrect tests to make sure program catches errors correctly
func (s *MySuite) TestPostfixToPrefixIncorrect(c *C) {
	_, err := PostfixToPrefix("+")
	c.Assert(err, ErrorMatches, "incorrect expression")

	_, err = PostfixToPrefix("5 +")
	c.Assert(err, ErrorMatches, "incorrect expression")

	_, err = PostfixToPrefix("")
	c.Assert(err, ErrorMatches, "empty input")

	_, err = PostfixToPrefix("a b +")
	c.Assert(err, ErrorMatches, "incorrect symbol.*")
}
