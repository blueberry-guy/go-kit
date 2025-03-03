package example

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/blueberry-guy/go-kit/csv"
)

// ExampleBasicWriting demonstrates basic CSV file writing
func ExampleBasicWriting() {
	// Create sample data
	people := []Person{
		{Name: "John Doe", Age: 30, City: "New York", Salary: 75000.00},
		{Name: "Jane Smith", Age: 25, City: "Los Angeles", Salary: 82000.50},
	}

	// Create a writer
	writer := csv.NewWriter()

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example*.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	// Write to file
	if err := writer.WriteToFile(people, tmpfile.Name()); err != nil {
		log.Fatal(err)
	}

	// Read and display the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Written CSV content:\n%s\n", string(content))
}

// ExampleWriteWithCustomHeaders demonstrates writing CSV with custom headers
func ExampleWriteWithCustomHeaders() {
	// Create sample data
	people := []Person{
		{Name: "John Doe", Age: 30, City: "New York", Salary: 75000.00},
		{Name: "Jane Smith", Age: 25, City: "Los Angeles", Salary: 82000.50},
	}

	// Create custom headers
	headers := []string{"Full Name", "Years Old", "Location", "Annual Wage"}

	// Create a writer
	writer := csv.NewWriter()

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example*.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	// Write with custom headers
	if err := writer.WriteWithHeaders(people, headers, tmpfile.Name()); err != nil {
		log.Fatal(err)
	}

	// Read and display the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Written CSV content with custom headers:\n%s\n", string(content))
}

// ExampleWriteToString demonstrates writing CSV to a string
func ExampleWriteToString() {
	// Create sample data
	people := []Person{
		{Name: "John Doe", Age: 30, City: "New York", Salary: 75000.00},
	}

	// Create a writer
	writer := csv.NewWriter()

	// Write to string
	content, err := writer.WriteString(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CSV as string:\n%s\n", content)
}

// ExampleWriteToBuffer demonstrates writing CSV to an io.Writer
func ExampleWriteToBuffer() {
	// Create sample data
	people := []Person{
		{Name: "John Doe", Age: 30, City: "New York", Salary: 75000.00},
	}

	// Create a buffer
	var buf bytes.Buffer

	// Create a writer
	writer := csv.NewWriter()

	// Write to buffer
	if err := writer.Write(&buf, people); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CSV from buffer:\n%s\n", buf.String())
}

// ExampleAppendToFile demonstrates appending records to an existing CSV file
func ExampleAppendToFile() {
	// Create initial data
	initialPeople := []Person{
		{Name: "John Doe", Age: 30, City: "New York", Salary: 75000.00},
	}

	// Create new data to append
	newPeople := []Person{
		{Name: "Jane Smith", Age: 25, City: "Los Angeles", Salary: 82000.50},
	}

	// Create a writer
	writer := csv.NewWriter()

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example*.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	// Write initial data
	if err := writer.WriteToFile(initialPeople, tmpfile.Name()); err != nil {
		log.Fatal(err)
	}

	// Append new data
	if err := writer.AppendToFile(newPeople, tmpfile.Name()); err != nil {
		log.Fatal(err)
	}

	// Read and display the content
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final CSV content after append:\n%s\n", string(content))
}
