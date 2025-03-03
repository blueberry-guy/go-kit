package internal

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/gocarina/gocsv"
)

// GoCsvReaderImpl implements CSVReader.
type GoCsvReaderImpl struct{}

// ReadFromFile reads CSV data from a file into a struct slice.
func (c *GoCsvReaderImpl) ReadFromFile(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.UnmarshalFile(file, v)
}

// Read reads CSV data from an io.Reader into a struct slice
func (c *GoCsvReaderImpl) Read(reader io.Reader, v interface{}) error {
	return gocsv.Unmarshal(reader, v)
}

// ReadString reads CSV data from a string into a struct slice
func (c *GoCsvReaderImpl) ReadString(data string, v interface{}) error {
	return c.Read(strings.NewReader(data), v)
}

// ReadWithHeaders reads CSV data with custom headers
func (c *GoCsvReaderImpl) ReadWithHeaders(fileName string, headers []string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new reader with custom headers
	csvContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Skip the first line (original headers) and prepend new headers
	lines := strings.Split(string(csvContent), "\n")
	if len(lines) > 1 {
		newContent := strings.Join(headers, ",") + "\n" + strings.Join(lines[1:], "\n")
		return c.ReadString(newContent, v)
	}

	return fmt.Errorf("empty or invalid CSV file")
}

// Validate validates if the CSV file matches the expected struct fields
func (c *GoCsvReaderImpl) Validate(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get CSV headers
	csvReader := gocsv.DefaultCSVReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	// Get struct fields
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Create a map of CSV tags from struct fields
	structTags := make(map[string]bool)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("csv")
		if tag != "" {
			structTags[tag] = true
		}
	}

	// Validate that all headers exist in struct tags
	for _, header := range headers {
		if !structTags[header] {
			return fmt.Errorf("CSV header '%s' does not match any struct field", header)
		}
	}

	return nil
}
