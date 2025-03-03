package csv

import (
	"io"
)

type CsvRecords []byte

// Writer defines methods for writing CSV data.
type Writer interface {
	// WriteToFile writes CSV data to a file
	WriteToFile(v interface{}, fileName string) error

	// Write writes CSV data to an io.Writer
	Write(writer io.Writer, v interface{}) error

	// WriteString writes CSV data and returns it as a string
	WriteString(v interface{}) (string, error)

	// WriteWithHeaders writes CSV data with custom headers
	WriteWithHeaders(v interface{}, headers []string, fileName string) error

	// AppendToFile appends CSV data to an existing file
	AppendToFile(v interface{}, fileName string) error
}
