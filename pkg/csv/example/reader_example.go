package example

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blueberry-guy/go-kit/pkg/csv"
)

// ExampleBasicReading demonstrates basic CSV file reading
func ExampleBasicReading() {
	// Sample CSV content
	csvContent := `name,age,city,salary
John Doe,30,New York,75000.00
Jane Smith,25,Los Angeles,82000.50`

	// Create a reader
	reader := csv.NewReader()

	// Read CSV string into slice of Person structs
	var people []Person
	if err := reader.ReadString(csvContent, &people); err != nil {
		log.Fatal(err)
	}

	// Process the records
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, City: %s, Salary: %.2f\n",
			person.Name, person.Age, person.City, person.Salary)
	}
}

// ExampleReadWithCustomHeaders demonstrates reading CSV with custom headers
func ExampleReadWithCustomHeaders() {
	// Sample CSV content with different headers
	csvContent := `full_name,years_old,location,annual_wage
John Doe,30,New York,75000.00
Jane Smith,25,Los Angeles,82000.50`

	// Write to temporary file
	tmpfile := createTempFile(csvContent)
	defer cleanupTempFile(tmpfile)

	// Create a reader
	reader := csv.NewReader()

	// Define custom headers that match our Person struct tags
	headers := []string{"name", "age", "city", "salary"}

	// Read with custom headers into Person structs
	var people []Person
	if err := reader.ReadWithHeaders(tmpfile, headers, &people); err != nil {
		log.Fatal(err)
	}

	// Process the records
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, City: %s, Salary: %.2f\n",
			person.Name, person.Age, person.City, person.Salary)
	}
}

// ExampleValidateCSV demonstrates CSV validation
func ExampleValidateCSV() {
	// Sample CSV content with valid headers
	validCSV := `name,age,city,salary
John Doe,30,New York,75000.00`

	// Sample CSV content with invalid headers
	invalidCSV := `full_name,years,location,wage
John Doe,30,New York,75000.00`

	reader := csv.NewReader()

	// Create temporary files
	validFile := createTempFile(validCSV)
	invalidFile := createTempFile(invalidCSV)
	defer cleanupTempFile(validFile)
	defer cleanupTempFile(invalidFile)

	// Validate valid CSV
	if err := reader.Validate(validFile, &Person{}); err != nil {
		fmt.Printf("Valid CSV validation failed: %v\n", err)
	} else {
		fmt.Println("Valid CSV structure verified")
	}

	// Validate invalid CSV
	if err := reader.Validate(invalidFile, &Person{}); err != nil {
		fmt.Printf("Invalid CSV validation failed as expected: %v\n", err)
	}
}

// ExampleReadFromReader demonstrates reading from an io.Reader
func ExampleReadFromReader() {
	// Sample CSV content
	csvContent := `name,age,city,salary
John Doe,30,New York,75000.00`

	// Create a reader from string
	strReader := strings.NewReader(csvContent)

	// Create CSV reader
	reader := csv.NewReader()

	// Read from io.Reader
	var people []Person
	if err := reader.Read(strReader, &people); err != nil {
		log.Fatal(err)
	}

	// Process the records
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, City: %s, Salary: %.2f\n",
			person.Name, person.Age, person.City, person.Salary)
	}
}

// Helper functions
func createTempFile(content string) string {
	tmpfile, err := os.CreateTemp("", "example*.csv")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.WriteString(content); err != nil {
		log.Fatal(err)
	}
	tmpfile.Close()
	return tmpfile.Name()
}

func cleanupTempFile(filename string) {
	os.Remove(filename)
}
