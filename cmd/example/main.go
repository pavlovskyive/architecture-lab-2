package main

import (
	"flag"
	lab2 "github.com/pavlovskyive/architecture-lab-2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Expression to compute")
	outputFile = flag.String("o", "", "Expression to compute")
)

func main() {
	flag.Parse()

	var source io.Reader
	var destination io.Writer

	if *inputExpression != "" {
		source = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		source = strings.NewReader(string(data))
	} else {
		_, _ = os.Stderr.WriteString("error: no expression provided")
		return
	}

	if *outputFile != "" {
		if file, err := os.Create(*outputFile); err == nil {
			destination = file
		} else {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
	} else {
		destination = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  source,
		Output: destination,
	}

	if err := handler.Compute(); err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}
}
