package tea

import (
	"io"
	"os"
)

func LogToFile(path string, prefix string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LogOptionsSetter interface {
	SetOutput(io.Writer)
	SetPrefix(string)
}

func LogToFileWith(path string, prefix string, log LogOptionsSetter) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:mnd
