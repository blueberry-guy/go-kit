package internal

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGoCsvWriterImpl_WriteToFile(t *testing.T) {
	// Create test data
	records := []TestRecord{
		{Name: "John Doe", Age: 30, City: "New York"},
		{Name: "Jane Smith", Age: 25, City: "Los Angeles"},
	}

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	// Write to file
	writer := &GoCsvWriterImpl{}
	err = writer.WriteToFile(records, tmpfile.Name())
	if err != nil {
		t.Errorf("WriteToFile failed: %v", err)
	}

	// Read and verify the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "John Doe,30,New York") {
		t.Errorf("Written content doesn't match expected values")
	}
}

func TestGoCsvWriterImpl_Write(t *testing.T) {
	records := []TestRecord{
		{Name: "John Doe", Age: 30, City: "New York"},
	}

	var buf bytes.Buffer
	writer := &GoCsvWriterImpl{}
	err := writer.Write(&buf, records)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	if !strings.Contains(buf.String(), "John Doe,30,New York") {
		t.Errorf("Written content doesn't match expected values")
	}
}

func TestGoCsvWriterImpl_WriteString(t *testing.T) {
	records := []TestRecord{
		{Name: "John Doe", Age: 30, City: "New York"},
	}

	writer := &GoCsvWriterImpl{}
	content, err := writer.WriteString(records)
	if err != nil {
		t.Errorf("WriteString failed: %v", err)
	}

	if !strings.Contains(content, "John Doe,30,New York") {
		t.Errorf("Written content doesn't match expected values")
	}
}

func TestGoCsvWriterImpl_WriteWithHeaders(t *testing.T) {
	records := []TestRecord{
		{Name: "John Doe", Age: 30, City: "New York"},
	}

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	// Write with custom headers
	writer := &GoCsvWriterImpl{}
	headers := []string{"Full Name", "Years", "Location"}
	err = writer.WriteWithHeaders(records, headers, tmpfile.Name())
	if err != nil {
		t.Errorf("WriteWithHeaders failed: %v", err)
	}

	// Read and verify the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedHeader := "Full Name,Years,Location"
	if !strings.Contains(string(content), expectedHeader) {
		t.Errorf("Custom headers not found in content")
	}
	if !strings.Contains(string(content), "John Doe,30,New York") {
		t.Errorf("Written content doesn't match expected values")
	}
}

func TestGoCsvWriterImpl_AppendToFile(t *testing.T) {
	// Create initial content
	initialRecords := []TestRecord{
		{Name: "John Doe", Age: 30, City: "New York"},
	}

	// Create a temporary file with initial content
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write initial content
	writer := &GoCsvWriterImpl{}
	err = writer.WriteToFile(initialRecords, tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Append new record
	newRecords := []TestRecord{
		{Name: "Jane Smith", Age: 25, City: "Los Angeles"},
	}

	err = writer.AppendToFile(newRecords, tmpfile.Name())
	if err != nil {
		t.Errorf("AppendToFile failed: %v", err)
	}

	// Read and verify the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "John Doe,30,New York") {
		t.Errorf("Original content not found")
	}
	if !strings.Contains(string(content), "Jane Smith,25,Los Angeles") {
		t.Errorf("Appended content not found")
	}

	// Count number of header rows (should be only one)
	headerCount := strings.Count(string(content), "name,age,city")
	if headerCount != 1 {
		t.Errorf("Expected 1 header row, got %d", headerCount)
	}
}
