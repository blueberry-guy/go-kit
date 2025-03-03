# CSV Examples

This package provides examples of how to use the CSV reader and writer functionality from the `github.com/blueberry-guy/go-kit/csv` package.

## Examples Overview

### Types
- `Person`: A basic struct with CSV tags for name, age, city, and salary
- `Employee`: An alternative struct with different field names to demonstrate custom header mapping

### Reader Examples
1. `ExampleBasicReading`: Basic CSV file reading
2. `ExampleReadWithCustomHeaders`: Reading CSV with custom headers
3. `ExampleValidateCSV`: CSV structure validation
4. `ExampleReadFromReader`: Reading from an io.Reader

### Writer Examples
1. `ExampleBasicWriting`: Basic CSV file writing
2. `ExampleWriteWithCustomHeaders`: Writing CSV with custom headers
3. `ExampleWriteToString`: Writing CSV to a string
4. `ExampleWriteToBuffer`: Writing CSV to an io.Writer
5. `ExampleAppendToFile`: Appending records to an existing CSV file

## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/blueberry-guy/go-kit/csv"
)

func main() {
    // Create a reader
    reader := csv.NewReader()
    
    // Read from a CSV string
    csvContent := `name,age,city,salary
John Doe,30,New York,75000.00`
    
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
```

## Features Demonstrated

### Reading Features
- Reading from files
- Reading from strings
- Reading from io.Reader
- Custom header mapping
- CSV validation
- Error handling

### Writing Features
- Writing to files
- Writing to strings
- Writing to io.Writer
- Custom headers
- Appending to existing files
- Error handling

## Best Practices
1. Always close files after operations
2. Use defer for cleanup operations
3. Handle errors appropriately
4. Use custom headers when working with different CSV formats
5. Validate CSV structure before processing
6. Use appropriate types in struct fields
7. Use meaningful CSV tags

## Notes
- All examples use temporary files that are automatically cleaned up
- Error handling is demonstrated in all examples
- Custom headers can be used to map different CSV formats to the same struct
- The examples show both basic and advanced usage patterns 