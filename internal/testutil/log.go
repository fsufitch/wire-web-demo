package testutil

import (
	"log"

	customlog "github.com/fsufitch/wire-web-demo/log"
)

type Foo interface{}

// HeadlessLogOutput is an io.Writer that writes to an array of strings
type HeadlessLogOutput struct {
	linesPtr *[]string
}

func (h HeadlessLogOutput) Write(bytes []byte) (int, error) {
	*h.linesPtr = append(*h.linesPtr, string(bytes))
	return len(bytes), nil
}

// Lines returns a copy of the current array of strings in the output
func (h HeadlessLogOutput) Lines() []string { return append([]string{}, *h.linesPtr...) }

func newHeadlessLogOutput() HeadlessLogOutput {
	return HeadlessLogOutput{&[]string{}}
}

// HeadlessMultiLogger returns a MultiLogger that writes to two HeadlessLogOutput instances, for test purposes
func HeadlessMultiLogger() (logger *customlog.MultiLogger, stdout HeadlessLogOutput, stderr HeadlessLogOutput) {
	stdout = newHeadlessLogOutput()
	stderr = newHeadlessLogOutput()
	logger = &customlog.MultiLogger{
		PrintLevel:  customlog.Debug,
		ErrorLevel:  customlog.Error,
		PrintLogger: log.New(stdout, "", 0),
		ErrorLogger: log.New(stderr, "", 0),
	}
	return
}
