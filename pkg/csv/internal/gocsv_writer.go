package internal

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

// GoCsvWriterImpl implements Writer.
type GoCsvWriterImpl struct{}

// WriteToFile writes CSV data to a file.
func (c *GoCsvWriterImpl) WriteToFile(v interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return gocsv.MarshalFile(v, file)
}

// Write writes CSV data to an io.Writer
func (c *GoCsvWriterImpl) Write(writer io.Writer, v interface{}) error {
	return gocsv.Marshal(v, writer)
}

// WriteString writes CSV data and returns it as a string
func (c *GoCsvWriterImpl) WriteString(v interface{}) (string, error) {
	var buf bytes.Buffer
	if err := c.Write(&buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// WriteWithHeaders writes CSV data with custom headers
func (c *GoCsvWriterImpl) WriteWithHeaders(v interface{}, headers []string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write custom headers first
	if _, err := file.WriteString(strings.Join(headers, ",") + "\n"); err != nil {
		return err
	}

	// Marshal the data without headers
	var buf bytes.Buffer
	if err := gocsv.Marshal(v, &buf); err != nil {
		return err
	}

	// Skip the header line from the marshaled content
	content := buf.String()
	if idx := bytes.Index(buf.Bytes(), []byte("\n")); idx >= 0 {
		content = content[idx+1:]
	}

	// Write the content
	_, err = file.WriteString(content)
	return err
}

// AppendToFile appends CSV data to an existing file
func (c *GoCsvWriterImpl) AppendToFile(v interface{}, fileName string) error {
	// Open file in append mode
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get CSV content without headers
	var buf bytes.Buffer
	if err := gocsv.Marshal(v, &buf); err != nil {
		return err
	}

	// Skip the header line from the marshaled content
	content := buf.String()
	if idx := bytes.Index(buf.Bytes(), []byte("\n")); idx >= 0 {
		content = content[idx+1:]
	}

	// Write the content
	_, err = file.WriteString(content)
	return err
}
