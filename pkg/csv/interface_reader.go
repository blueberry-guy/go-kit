package csv

import (
	"io"
)

// Reader defines methods for reading CSV data.
type Reader interface {
	// ReadFromFile reads CSV data from a file into a struct slice
	ReadFromFile(fileName string, v interface{}) error

	// Read reads CSV data from an io.Reader into a struct slice
	Read(reader io.Reader, v interface{}) error

	// ReadString reads CSV data from a string into a struct slice
	ReadString(data string, v interface{}) error

	// ReadWithHeaders reads CSV data with custom headers
	ReadWithHeaders(fileName string, headers []string, v interface{}) error

	// Validate validates if the CSV file matches the expected struct fields
	Validate(fileName string, v interface{}) error
}
