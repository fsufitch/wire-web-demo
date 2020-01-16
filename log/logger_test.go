package log

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPrefix(t *testing.T) {
	// Setup
	buf1 := bytes.NewBufferString("")
	buf2 := bytes.NewBufferString("")
	logger := MultiLogger{
		PrintLogger: log.New(buf1, "", 0),
		ErrorLogger: log.New(buf2, "", 0),
	}

	// Tested code
	logger.SetPrefix("abc")

	// Asserts
	assert.Equal(t, "abc", logger.PrintLogger.Prefix())
	assert.Equal(t, "abc", logger.ErrorLogger.Prefix())
}

func TestSetFlags(t *testing.T) {
	// Setup
	buf1 := bytes.NewBufferString("")
	buf2 := bytes.NewBufferString("")
	logger := MultiLogger{
		PrintLogger: log.New(buf1, "", 0),
		ErrorLogger: log.New(buf2, "", 0),
	}
	flag := log.Ldate | log.Ltime | log.LUTC

	// Tested code
	logger.SetFlags(flag)

	// Asserts
	assert.Equal(t, flag, logger.PrintLogger.Flags())
	assert.Equal(t, flag, logger.ErrorLogger.Flags())
}

func TestSetOutput(t *testing.T) {
	// Setup
	buf1 := bytes.NewBufferString("")
	buf2 := bytes.NewBufferString("")
	logger := MultiLogger{
		PrintLogger: log.New(buf1, "", 0),
		ErrorLogger: log.New(buf2, "", 0),
	}
	buf3 := bytes.NewBufferString("")
	buf4 := bytes.NewBufferString("")

	// Tested code
	logger.SetOutput(buf3, buf4)

	// Asserts
	assert.Equal(t, buf3, logger.PrintLogger.Writer())
	assert.Equal(t, buf4, logger.ErrorLogger.Writer())
}

func TestLevels(t *testing.T) {
	// Setup
	buf1 := bytes.NewBufferString("")
	buf2 := bytes.NewBufferString("")
	logger := MultiLogger{
		PrintLogger: log.New(buf1, "", 0),
		ErrorLogger: log.New(buf2, "", 0),
		PrintLevel:  Info,
		ErrorLevel:  Error,
	}
	testStr := "test%d"

	// Tested code
	logger.Debugf(testStr, 1)
	logger.Infof(testStr, 2)
	logger.Warningf(testStr, 3)
	logger.Errorf(testStr, 4)
	logger.Criticalf(testStr, 5)

	// Asserts
	assert.NotContains(t, buf1.String(), "test1")
	assert.Contains(t, buf1.String(), "test2")
	assert.Contains(t, buf1.String(), "test3")
	assert.NotContains(t, buf1.String(), "test4")
	assert.NotContains(t, buf1.String(), "test5")

	assert.NotContains(t, buf2.String(), "test1")
	assert.NotContains(t, buf2.String(), "test2")
	assert.NotContains(t, buf2.String(), "test3")
	assert.Contains(t, buf2.String(), "test4")
	assert.Contains(t, buf2.String(), "test5")
}

func TestProvideStdOutErrMultiLogger(t *testing.T) {
	// Tested code
	debugLogger, _ := ProvideStdOutErrMultiLogger(true)
	prodLogger, _ := ProvideStdOutErrMultiLogger(false)

	// Asserts
	assert.Equal(t, os.Stdout, debugLogger.PrintLogger.Writer())
	assert.Equal(t, os.Stderr, debugLogger.ErrorLogger.Writer())
	assert.Equal(t, Debug, debugLogger.PrintLevel)
	assert.Equal(t, Error, debugLogger.ErrorLevel)
	assert.Equal(t, log.Ldate|log.Ltime|log.LUTC|log.Llongfile|log.Lmicroseconds, debugLogger.PrintLogger.Flags())
	assert.Equal(t, log.Ldate|log.Ltime|log.LUTC|log.Llongfile|log.Lmicroseconds, debugLogger.ErrorLogger.Flags())

	assert.Equal(t, os.Stdout, prodLogger.PrintLogger.Writer())
	assert.Equal(t, os.Stderr, prodLogger.ErrorLogger.Writer())
	assert.Equal(t, Info, prodLogger.PrintLevel)
	assert.Equal(t, Error, prodLogger.ErrorLevel)
	assert.Equal(t, log.Ldate|log.Ltime|log.LUTC, prodLogger.PrintLogger.Flags())
	assert.Equal(t, log.Ldate|log.Ltime|log.LUTC, prodLogger.ErrorLogger.Flags())

}
