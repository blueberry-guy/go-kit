package internal

import (
	"os"
	"strings"
	"testing"
)

type TestRecord struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
	City string `csv:"city"`
}

func TestGoCsvReaderImpl_ReadFromFile(t *testing.T) {
	// Create a temporary CSV file
	content := "name,age,city\nJohn Doe,30,New York\nJane Smith,25,Los Angeles"
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Test reading from file
	reader := &GoCsvReaderImpl{}
	var records []TestRecord
	err = reader.ReadFromFile(tmpfile.Name(), &records)
	if err != nil {
		t.Errorf("ReadFromFile failed: %v", err)
	}

	// Verify the records
	if len(records) != 2 {
		t.Errorf("Expected 2 records, got %d", len(records))
	}
	if records[0].Name != "John Doe" || records[0].Age != 30 || records[0].City != "New York" {
		t.Errorf("First record doesn't match expected values")
	}
}

func TestGoCsvReaderImpl_ReadString(t *testing.T) {
	reader := &GoCsvReaderImpl{}
	csvContent := "name,age,city\nJohn Doe,30,New York"

	var records []TestRecord
	err := reader.ReadString(csvContent, &records)
	if err != nil {
		t.Errorf("ReadString failed: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Expected 1 record, got %d", len(records))
	}
	if records[0].Name != "John Doe" || records[0].Age != 30 || records[0].City != "New York" {
		t.Errorf("Record doesn't match expected values")
	}
}

func TestGoCsvReaderImpl_ReadWithHeaders(t *testing.T) {
	// Create a temporary CSV file with different headers
	content := "full_name,years,location\nJohn Doe,30,New York"
	tmpfile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Test reading with custom headers
	reader := &GoCsvReaderImpl{}
	var records []TestRecord
	headers := []string{"name", "age", "city"}
	err = reader.ReadWithHeaders(tmpfile.Name(), headers, &records)
	if err != nil {
		t.Errorf("ReadWithHeaders failed: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Expected 1 record, got %d", len(records))
	}
	if records[0].Name != "John Doe" || records[0].Age != 30 || records[0].City != "New York" {
		t.Errorf("Record doesn't match expected values")
	}
}

func TestGoCsvReaderImpl_Validate(t *testing.T) {
	tests := []struct {
		name    string
		headers string
		wantErr bool
	}{
		{
			name:    "Valid headers",
			headers: "name,age,city",
			wantErr: false,
		},
		{
			name:    "Invalid header",
			headers: "name,age,country",
			wantErr: true,
		},
		{
			name:    "Missing header",
			headers: "name,age",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary CSV file
			tmpfile, err := os.CreateTemp("", "test*.csv")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			content := tt.headers + "\nJohn Doe,30,New York"
			if _, err := tmpfile.WriteString(content); err != nil {
				t.Fatal(err)
			}
			tmpfile.Close()

			reader := &GoCsvReaderImpl{}
			err = reader.Validate(tmpfile.Name(), &TestRecord{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGoCsvReaderImpl_Read(t *testing.T) {
	reader := &GoCsvReaderImpl{}
	csvContent := "name,age,city\nJohn Doe,30,New York"

	var records []TestRecord
	err := reader.Read(strings.NewReader(csvContent), &records)
	if err != nil {
		t.Errorf("Read failed: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("Expected 1 record, got %d", len(records))
	}
	if records[0].Name != "John Doe" || records[0].Age != 30 || records[0].City != "New York" {
		t.Errorf("Record doesn't match expected values")
	}
}
